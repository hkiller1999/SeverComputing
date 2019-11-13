### 1、概述
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）
### 2、任务要求
编程 web 应用程序 cloudgo-io。 请在项目 README.MD 给出完成任务的证据！

基本要求

1、支持静态文件服务

2、支持简单 js 访问

3、提交表单，并输出一个表格

4、对 `/unknown` 给出开发中的提示，返回码·`5xx`

###测试
按照[潘老师blog](https://blog.csdn.net/pmlpml/article/details/78539261)的提示，我首先对一些包进行了了解。
首先
```
go get github.com/codegangsta/negroni
go get github.com/gorilla/mux
go get github.com/unrolled/render 
 ```
导入这些包以后使用`get install`和`get build`对源码进行编译，使用命令`get run main.go`运行main.go函数
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113223039628.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hraWxsZXIxOTk5,size_16,color_FFFFFF,t_70)

##### 静态文件服务
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113223707849.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hraWxsZXIxOTk5,size_16,color_FFFFFF,t_70)
##### js访问
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113223738417.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hraWxsZXIxOTk5,size_16,color_FFFFFF,t_70)
##### 提交表单，并输出一个表格
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113223851359.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hraWxsZXIxOTk5,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113223909747.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hraWxsZXIxOTk5,size_16,color_FFFFFF,t_70)
##### 对 `/unknown` 给出开发中的提示，返回码·`5xx`![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113223945385.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hraWxsZXIxOTk5,size_16,color_FFFFFF,t_70)

