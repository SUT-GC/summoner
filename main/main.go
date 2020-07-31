package main

import (
	"github.com/gin-gonic/gin"
	"summoner/routers"
)

const htmlGlob = "template/*"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob(htmlGlob)

	// 初始化router容器
	routerContainer := routers.RouterContainer{}

	// 首页 router
	indexRouter := routerContainer.IndexRouter()
	router.GET(indexRouter.Path, indexRouter.Handler)

	// 文章列表 router
	blogListRouter := routerContainer.BlogSummaryListDefaultRouter()
	router.GET(blogListRouter.Path, blogListRouter.Handler)

	// 文章列表排序 router
	blogListOrderRouter := routerContainer.BlogSummaryListAssignRouter()
	router.GET(blogListOrderRouter.Path, blogListOrderRouter.Handler)

	err := router.Run(":8080")

	panic(err)
}
