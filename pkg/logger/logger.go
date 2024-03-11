package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

// Level 预定义了应用日志的 Level 和 Fields 的具体类型，并且分为了 Debug、Info、Warn、Error、Fatal、Panic 六个日志等级
type Level int8

// Fields 定义了日志的字段类型
type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// String 将日志等级转换为字符串
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"

	}
	return ""
}

// Logger 定义了日志的结构体，包含了一个 log.Logger 类型的 newLogger 字段，一个 context.Context 类型的 ctx 字段，一个 Fields 类型的 fields 字段，一个 []string 类型的 callers 字段
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

// NewLogger 实例化一个 Logger 结构体
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// WithFields 设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	// 将 f 的字段复制到 ll.fields 中
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

// WithContext 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	// 将 ctx 赋值给 ll.ctx
	ll.ctx = ctx
	return ll
}

// WithCaller 设置当前某一层调用栈的信息（程序计数器、文件信息、行号）skip参数决定向上追溯多少层堆栈
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	// runtime.Caller 获取当前的调用栈信息
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		// runtime.FuncForPC 获取调用栈信息对应的函数信息
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

// WithCallersFrames 设置当前的整个调用栈信息 日志内容将包含从最小调用深度到调用栈末尾的所有调用者信息（文件名、行号和函数名）。
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	// pcs 为程序计数器
	pcs := make([]uintptr, maxCallerDepth)
	// depth 为调用栈的深度
	depth := runtime.Callers(minCallerDepth, pcs)
	// frams 为调用栈的帧信息
	frames := runtime.CallersFrames(pcs[:depth])
	// 遍历调用栈的帧信息
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// JSONFormat 日志内容的格式化和日志输出动作的相关方法
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	// len(l.fields)+4 为了预留 level、time、message、callers 四个字段
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		// 遍历 l.fields，将 l.fields 的字段复制到 data 中
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

// Output 根据日志等级，输出日志内容
func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Panic(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// Info 根据先前定义的日志分级，编写对应的日志输出的外部方法
func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

// Infof
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

// Fatal
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

// Fatalf
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

// Errorf
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}
