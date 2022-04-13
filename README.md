# 师资培训管理系统

#### 基于beego 2.0.1 框架和AdminLte前端框架开发的高性能教师资格培训管理系统，在beego框架的基础上，封装了后台系统的分页功能，excel数据导出功能等丰富常用的扩展，基于MVC模式，html界面随心定义，相较于某些后台复杂代码生成的前端html元素，使用原生的html原生作为前端显示，更加的灵活自由。欢迎大家使用。

## 安装

### 安装方式 (GO MOD方式安装,已移除 GOPATH方式安装说明，需要的请查看 tag v1.0.1)

#### 1、安装beego v2.0.1和bee v2.0.2

参考[Beego](https://beego.me/docs/install/)和[Bee](https://beego.me/docs/install/bee.md)安装手册

#### 2、配置数据库

```
将目录中init.sql文件导入sql数据库

更改根目录下的config.yaml文件内的数据库连接信息
```

#### 3、安装项目依赖

```
项目根目录下 go mod tidy 将自动下载依赖包
```

### 通过上面方式安装后,接下来

#### 运行系统

```
直接运行go run main.go，或者使用bee run在项目下运行，开始进行调试开发
```

#### 访问后台

访问`/admin/index/index`，默认超级管理员的账号密码都为`super_admin`。

## 补充

本项目在beego@v2.0.1的框架基础上完善了很多丰富的常用后台功能，分页封装、excel数据一键导出等功能，目前没有开发手册，相信大家一看代码就可知道功能怎么使用，如果大家需要详细的使用手册，我可为大家写一份详细的功能使用介绍，此外，如果有需要php语言的laravel版本的后台管理系统，可以使用[laravel-admin](https://github.com/yuxingfei/laravel-admin)。

## 联系我

邮箱：1903249375@qq.com
