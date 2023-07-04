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







