package biz

import (
	"summoner/repo"
	"summoner/sconsts"
)

type BlogService interface {
	// 查询所有的Blog Summary 默认排序
	QueryAllBlogSummaryDefaultOrder() ([]BlogSummaryDTO, error)

	// 查询所有的Blog Summary 定制排序
	QueryAllBlogSummaryAssignOrder(sortType ...sconsts.SortType) ([]BlogSummaryDTO, error)

	// 查询blog信息 ByID
	QueryBlogById(blogId int64) (BlogDTO, error)

	// 添加Blog
	AddBlog(blog BlogDTO) (int64, error)
}

// blog dao
var blogDao = new(repo.BlogDaoImpl)

type BlogServiceImpl struct {
}

func (service *BlogServiceImpl) QueryAllBlogSummaryDefaultOrder() ([]BlogSummaryDTO, error) {
	blogSummaries := make([]BlogSummaryDTO, 0)

	ids, err := blogDao.QueryAllBlogIds()
	if ids == nil || err != nil {
		return blogSummaries, err
	}

	if len(ids) <= 0 {
		return blogSummaries, nil
	}

	blogPOs, err := blogDao.QueryBlogInfoByIds(ids...)
	if err != nil {
		return blogSummaries, err
	}

	for _, blogPO := range blogPOs {
		blogSummaries = append(blogSummaries, BlogPO2BlogSummaryDTO(blogPO))
	}

	return blogSummaries, nil
}

func (service *BlogServiceImpl) QueryAllBlogSummaryAssignOrder(sortType ...sconsts.SortType) ([]BlogSummaryDTO, error) {
	blogSummaries := make([]BlogSummaryDTO, 0)

	ids, err := blogDao.QueryAllBlogIds(sortType...)
	if ids == nil || err != nil {
		return blogSummaries, err
	}

	if len(ids) <= 0 {
		return blogSummaries, nil
	}

	blogPOs, err := blogDao.QueryBlogInfoByIds(ids...)
	if err != nil {
		return blogSummaries, err
	}

	for _, blogPO := range blogPOs {
		blogSummaries = append(blogSummaries, BlogPO2BlogSummaryDTO(blogPO))
	}

	return blogSummaries, nil
}

func (service *BlogServiceImpl) QueryBlogById(blogId int64) (bool, BlogDTO, error) {
	blogs, err := blogDao.QueryBlogInfoByIds(blogId)
	if err != nil {
		return false, BlogDTO{}, err
	}

	if len(blogs) <= 0 {
		return false, BlogDTO{}, nil
	}

	return true, BlogPO2BlogDTO(blogs[0]), nil
}

func (service *BlogServiceImpl) AddBlog(blog BlogDTO) (int64, error) {

	return blogDao.AddBlogInfo(BlogDTO2BlogPO(blog))
}
