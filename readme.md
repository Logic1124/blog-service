## 目录结构
configs：配置文件。       
docs：文档集合。  
global：全局变量。    
internal：内部模块。  
dao：数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL、ElasticSearch 等。   
middleware：HTTP 中间件。    
model：模型层，用于存放 model 对象。    
routers：路由相关逻辑处理。   
service：项目核心业务逻辑。   
pkg：项目相关的模块包。   
storage：项目生成的临时文件。  
scripts：各类构建，安装，分析等操作的脚本。   
third_party：第三方的资源工具，例如 Swagger UI。     
## 接口文档
### 标签管理接口
- 新增标签	POST	/tags   
- 删除指定标签	DELETE	/tags/:id   
- 更新指定标签	PUT	/tags/:id   
- 获取标签列表	GET	/tags   
### 文章管理接口
- 新增文章	POST	/articles
- 删除指定文章	DELETE	/articles/:id
- 更新指定文章	PUT	/articles/:id
- 获取指定文章	GET	/articles/:id
- 获取文章列表	GET	/articles
## 数据库
- main.go中初始化数据库连接要引入	_ "github.com/jinzhu/gorm/dialects/mysql"
## 响应处理
- 编写统一处理接口返回的响应处理方法，它也正正与错误码标准化是相对应的。