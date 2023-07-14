package main

import (
	"mymodule/common"
	"mymodule/route"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = route.CollectRoute(r)
	panic(r.Run())

}
