# 描述
这是一个结构清晰的gin脚手架，有dao,dto,controller层。

此项目以gin框架为基础，构建了用户注册，用户登录，获取用户信息，用户退出，用户注销五个接口。以中间件的形式构建了翻译器，验证器，自定义返回方法

# 使用方法
此代码用到了mysql,redis,swagger，需提前配置好环境

1、用git将项目保存到本地

```
git clone https://github.com/uptocorrupt/gin-scaffold.git
```

2、进config文件夹里修改配置文件dev.yml:server、datasource、redis、swagger.并将其重命名为application.yml

3、在mysql建一个数据库

4、进项目根目录初始化swag

```
swag init
```
5、运行项目

6、在网页上输入host/swagger/index.html，测试和查看接口
如（[http://127.0.0.1:8090/swagger/index.html](http://127.0.0.1:8090/swagger/index.html)）
