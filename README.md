
<h1 style="text-align: center">图片/名片系统</h1>


#### 项目简介

基于当前流行技术组合的前后端图片系统(包含了go后端、后台vue前端)：
Gin+Gorm+Casbin+Jwt+Redis+Mysql8+Vue 



#### 详细结构

```
- app 应用模块
    - controllers 控制器模块
      - front
        - wechatController.go 微信控制器
    - listen redis监听器
    - models 模型模块
    - service 服务模块
      - product_serive 商品服务
      - wechat_menu_serive 微信公众号菜单服务
      ......
- conf 公共配置
- docs swagger
- middleware 中间件
    - AuthCheck.go
	- cors.go 
	......
- pkg 程序应用包
  - app
  - base
  - casbin
  - jwt
  - redis
  - wechat
  .....
- routere 路由
- logs 日志存放
- runtime 资源目录
- config.yml yml配置
```
#### go后端配置、启动、部署
```
1、安装go>=1.15,这个可以https://studygolang.com/dl下载

2、开启mod： go env -w GO111MODULE=on

3、配置代理：go env -w GOPROXY=https://goproxy.cn,direct 这个让下载依赖速度更快

4。下载项目：git clone https://gitee.com/guchengwuyue/yshop-gin.git

5、go mod tidy 安装所需依赖

6、导入sql/yshop_go.sql,修改cconfig,yml 里数据库与redis配置

7、本地运行air -c .air.conf 启动即可，如果没有安装air的 直接go run main,go即可

8、线上部署： go build  打包然后上传服务器命令： nohup ./go-mall & 

9、配置nginx 反向代理即可
```
#### pc用户端配置、启动、部署
```
1、打开后端go代码目录pc-vue/,然后解压

2、npm install

3、./api/index.js 配置api地址

4。npm run dev 启动即可

5、线上部署：npm run build 然后上传./dist/下编译后的文件到web服务器即可

```
#### 功能说明
效果图：

| ![输入图片说明](https://d2wqe7ddhlev2l.cloudfront.net/picsite/pic_admin3.png) |
|--|

| ![输入图片说明](https://d2wqe7ddhlev2l.cloudfront.net/picsite/pic_admin2.png) |
|--|

| ![输入图片说明](https://d2wqe7ddhlev2l.cloudfront.net/picsite/pic_admin.png "suicai.png") |
|--|

| ![输入图片说明](https://d2wqe7ddhlev2l.cloudfront.net/picsite/pic_admin_login.png "suicai.png") |
|--|


#### 技术选型
* 1 后端使用技术
    * 1.1 gin
    * 1.2 jwt
    * 1.3 redis
    * 1.5 Mysql8
    * 1.6 Gorm
    * 1.7 copier
    * 1.8 ksuid
    * 1.9 Redis
    * 1.10 swagger
    * 1.11 Casbin
    * 1.12 viper
    * 1.13 zap
    * 1.14 wecchat
    * 1.15 gopay
        
* 前端使用技术
    * 2.1 Vue 全家桶
    * 2.2 Element

#### 特别鸣谢

- go-gin-example:https://github.com/EDDYCJY/go-gin-example
- gorm:https://gorm.io/
- casbin:https://casbin.org/
- wechat: https://github.com/silenceper/wechat
- gopay: https://github.com/go-pay/gopay
- vue:https://github.com/vuejs/vue
- element:https://github.com/ElemeFE/element

