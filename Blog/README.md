一、基础代码说明：
1.设计状态码和提示信息:
types/code.go

2.统一APT返回格式:
utils/result.go

3.md5加密:
utils/encrypt.go

4.连接访问数据库:
dao/core.go

5.设计路由:
routers/router.go

6.中间件：
middleware/AuthMiddleware.go

二、功能代码说明：
1.创建管理员、 账号列表、删除账号、 更新账号、账号详情、修改账号信息、禁用、启用
models/user.go
services/user.go
controllers/user.go

2. 创建标签、标签列表、修改标签名字、删除标签
models/tags.go
controllers/tags.go
services/tags.go

3.创建分类、分类列表、修改分类名字、删除分类
models/cate.go
controllers/cate.go
services/cate.go

4. 写文章、文章列表、文章修改、删除文章
models/posts.go
controllers/posts.go
services/posts.go

5. 评论、评论列表、删除评论
models/comment.go
controllers/comment.go
services/comment.go

6. 添加友联、修改友联、删除友联
models/link.go
controllers/link.go
services/link.go

7.授权认证
controllers/auth.go

三、调用关系说明
main.go —>(调用) routers/router.go；
routers/router.go —>(调用) controllers/auth.go[ AuthLogin() ] 进行身份验证登录；
routers/router.go —>(调用) middleware/AuthMiddleware.go 进行中间件认证，middleware/AuthMiddleware.go —>(调用) controllers/auth.go[ ParseToken() ] 进行令牌解析
routers/router.go —>(调用) controllers/user.go —>(调用) services/user.go —>(调用) dao/core.go
routers/router.go —>(调用) controllers/tags.go —>(调用) services/tags.go —>(调用) dao/core.go
routers/router.go —>(调用) controllers/cate.go —>(调用) services/cate.go —>(调用) dao/core.go
routers/router.go —>(调用) controllers/posts.go —>(调用) services/posts.go —>(调用) dao/core.go
routers/router.go —>(调用) controllers/comment.go —>(调用) services/comment.go —>(调用) dao/core.go
routers/router.go —>(调用) controllers/link.go —>(调用) services/link.go —>(调用) dao/core.go

四、项目启动方式
1.先启动后端服务：
D:\WorkSpace\GoFun\Golang\src\Blog>go run main.go
或者
D:\WorkSpace\GoFun\Golang\src\Blog>gowatch     //通过热加载的方式

2.再启动前端vue服务
a>在windows下安装npm软件
b>使用powershell终端进入Blog/admin目录
c>使用npm命令安装项目
npm install
d>使用npm命令运行项目
npm run serve

五、访问项目
http://192.168.2.2:8081
用户名：admin
密码：123456

————————————————————————————————————————————————————————————————————————————————————

参考:
http://blog.caixiaoxin.cn/?p=284