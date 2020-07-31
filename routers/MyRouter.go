package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"summoner/biz"
	"summoner/sconsts"
)

type MyRouter struct {
	Path    string
	Handler gin.HandlerFunc
}

type RouterContainer struct {
}

var blogService = biz.BlogServiceImpl{}

func (container *RouterContainer) IndexRouter() MyRouter {
	return MyRouter{sconsts.IndexPath, func(context *gin.Context) {
		context.HTML(http.StatusOK, sconsts.IndexTemplate, gin.H{})
	}}
}

func (container *RouterContainer) BlogSummaryListAssignRouter() MyRouter {
	return MyRouter{sconsts.BlogSummaryListAssignPath, func(context *gin.Context) {
		sorts := context.QueryArray("sort")
		sortTypes := make([]sconsts.SortType, 0)
		for _, sort := range sorts {
			if string(sconsts.SortByCommentCountDesc) == sort {
				sortTypes = append(sortTypes, sconsts.SortByCommentCountDesc)
			}
			if string(sconsts.SortByViewCountDesc) == sort {
				sortTypes = append(sortTypes, sconsts.SortByViewCountDesc)
			}
		}

		summaries, err := blogService.QueryAllBlogSummaryAssignOrder(sortTypes...)

		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
		} else {
			context.JSON(http.StatusOK, summaries)
		}
	}}
}

func (container *RouterContainer) BlogSummaryListDefaultRouter() MyRouter {
	return MyRouter{sconsts.BlogSummaryListDefaultPath, func(context *gin.Context) {
		summaries, err := blogService.QueryAllBlogSummaryDefaultOrder()

		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
		} else {
			context.JSON(http.StatusOK, summaries)
		}
	}}
}
