package main

import (
	"fmt"
	"summoner/repo"
)

const htmlGlob = "template/*"

func main() {
	dao := repo.BlogDaoImpl{}
	r, _ := dao.QueryAllBlogIds(repo.SORT_BY_COMMENT_COUNT_DESC)
	blogs, err := dao.QueryBlogInfoByIds(r)
	fmt.Println(blogs, err)

	//router := gin.Default()
	//router.LoadHTMLGlob(htmlGlob)
	//
	//// 初始化router容器
	//routerContainer := routers.RouterContainer{}
	//
	//// 首页 router
	//indexRouter := routerContainer.IndexRouter()
	//router.GET(indexRouter.Path, indexRouter.Handler)
	//
	//// 文章列表 router
	//blogListRouter := routerContainer.BlogListRouter()
	//router.GET(blogListRouter.Path, blogListRouter.Handler)
	//
	//_ = router.Run(":8080")
}
