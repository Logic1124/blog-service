// Description: 类型转换
package convert

import "strconv"

type StrTo string

// String 将 StrTo 类型转换为 string 类型
func (s StrTo) String() string {
	return string(s)
}

// Int 将 StrTo 类型转换为 int 类型
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt 将 StrTo 类型转换为 int 类型，如果转换失败则返回 0
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// UInt32 将 StrTo 类型转换为 uint32 类型
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// MustUInt32 将 StrTo 类型转换为 uint32 类型，如果转换失败则返回 0
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
