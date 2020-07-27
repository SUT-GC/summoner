package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"summoner/consts"
)

type MyRouter struct {
	Path    string
	Handler gin.HandlerFunc
}

type RouterContainer struct {
}

func (container *RouterContainer) IndexRouter() MyRouter {
	return MyRouter{consts.IndexPath, func(context *gin.Context) {
		context.HTML(http.StatusOK, consts.IndexTemplate, gin.H{})
	}}
}

func (container *RouterContainer) BlogListRouter() MyRouter {
	return MyRouter{consts.BlogListPath, func(context *gin.Context) {
		sort := context.DefaultQuery("sort", "default")
		fmt.Println(sort)
		context.JSON(http.StatusOK, gin.H{"sort": sort})
	}}
}
