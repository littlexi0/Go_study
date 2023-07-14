// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func getting(c *gin.Context) {
// 	c.String(200, "hello %s", "world")
// 	c.JSON(http.StatusOK, gin.H{ //以json格式输出
// 		"name": "tom",
// 		"age":  "20",
// 	}) //200代表请求成功,http.StatusOK代表请求成功,gin.H是map的一个快捷方式
// }

// func welcome(c *gin.Context) {
// 	firstname := c.DefaultQuery("firstname", "Guest") //获取参数,如果没有则使用默认值
// 	// firstname := c.Query("firstname")
// 	lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

// 	c.String(http.StatusOK, "你好呀 %s %s", firstname, lastname)
// 	c.JSON(200, gin.H{
// 		firstname: "tom",
// 		lastname:  "jerry",
// 	})
// }

// func main() {
// 	router := gin.Default()

// 	// This handler will match /user/john but will not match /user/ or /user
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	// However, this one will match /user/john/ and also /user/john/send
// 	// If no other routers match /user/john, it will redirect to /user/john/
// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + " is " + action
// 		c.String(http.StatusOK, message)
// 	})

// 	// For each matched request Context will hold the route definition
// 	router.POST("/user/:name/*action", func(c *gin.Context) {
// 		b := c.FullPath() == "/user/:name/*action" // true
// 		c.String(http.StatusOK, "%t", b)
// 	})

// 	// This handler will add a new router for /user/groups.
// 	// Exact routes are resolved before param routes, regardless of the order they were defined.
// 	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
// 	router.GET("/user/groups", func(c *gin.Context) {
// 		c.String(http.StatusOK, "The available groups are [...]")
// 	})

// 	router.GET("/welcome", welcome)

// 	router.Run(":8080")
// }
