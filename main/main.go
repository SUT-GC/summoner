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
	blogListRouter := routerContainer.BlogListRouter()
	router.GET(blogListRouter.Path, blogListRouter.Handler)

	_ = router.Run(":8080")
}
