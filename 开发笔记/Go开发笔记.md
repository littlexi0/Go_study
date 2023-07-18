# Go开发笔记

## 安装Go

1、将Go安装到计算机，直接官网[The Go Programming Language](https://go.dev/)下载就好啦，下载好后默认安装到C盘，然后将Go文件目录下的bin文件路径加到系统PATH里面

## 创建go.mod

1、创建一个后端文件夹，cmd输入`go mod init <your.module.name>`创建go.mod文件

## 初始化依赖

1、通过cmd输入`go get -u github.com/gin-gonic/gin`初始化Go的文件依赖，可以发现文件夹里面多了go.sum文件

## 简单运行

在example.go文件里面添加一个简单的测试代码

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(200, "hello %s", "world")
	c.JSON(http.StatusOK, gin.H{ //以json格式输出
		"name": "tom",
		"age":  "20",
	}) //200代表请求成功,http.StatusOK代表请求成功,gin.H是map的一个快捷方式
}

func main() {
	e := gin.Default() //创建一个默认的路由引擎
	e.GET("/hello", Hello)
	e.Run()
}
```

终端输入`go run example.go`,即可发现运行成功

通过默认端口8080访问即可

*感觉是能最快速启动的后端了*

## 链接数据库

安装gorm`go get -u gorm.io/gorm`

定义表结构体

```go
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
```

安装mysql`go get gorm.io/driver/mysql@latest`

导入包：`import "gorm.io/driver/mysql"`

在GORM v2 版本中，`*gorm.DB` 类型的对象没有 `Close()` 方法。因此，您不需要显式地关闭数据库连接。

## 数据库操作

`db.AutoMigrate(&User{})`创建table，其中User是之前定义的结构体

`db.Create(&newUser)`向数据库插入数据





## 分发TOKEN

安装jwt包：`go get github.com/dgrijalva/jwt-go`

## 引入config组件

安装viper`go get github.com/spf13/viper`

## 安装node-sass

node14对应的版本

`npm install node-sass@4.14.1`

## vue路由

安装路由

使用vue2安装对应的router3版本
使用vue3安装对应的router4版本

`npm install vue-router@3`

要在app.vue里面加入`<router-view></router-view> `











