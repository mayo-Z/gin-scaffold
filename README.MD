# 描述
此项目以gin框架为基础，构建了用户注册，用户登录，获取用户信息，用户退出，用户注销五个接口。以中间件的形式构建了翻译器，验证器，自定义返回方法

# 使用方法
此代码用到了mysql,redis,swagger，需提前配置好环境
1、用git将项目保存到本地
```go
git clone https://github.com/uptocorrupt/auth_frame.git
```
2、进config文件夹里修改配置文件

```go
server:
  port: 8090

datasource:
  driverName: mysql
  name: root
  password: root
  host: localhost
  port: 3306
  database: auth_frame
  charset: utf8

redis:
  size: 10
  network:  tcp
  address:  127.0.0.1:6379
  password: root
  key:  secret

swagger:
  title: auth_frame swagger API
  desc: This is a frame
  host: 127.0.0.1:8090
  base_path:
```
3、在mysql建一个数据库(库名如上“auth_frame”)

4、进项目根目录初始化swag

```powershell
>swag init
```
5、运行项目

6、在网页上输入host/swagger/index.html，测试和查看接口
如（[http://127.0.0.1:8090/swagger/index.html](http://127.0.0.1:8090/swagger/index.html)）
