package repo

import (
	"summoner/utils"
	"time"
)

type BlogInfoPO struct { // blog信息
	Id           int64     // 主键id
	UserID       int64     // 用户id
	Title        string    // 标题
	Summary      string    // 摘要
	Content      string    // 内容
	ViewCount    int64     // 浏览计数
	CommentCount int64     // 评论计数
	Status       int       // 状态
	CreatedAt    time.Time // 创建时间
	UpdatedAt    time.Time // 更新时间
	DeletedAt    time.Time // 删除时间
}

type UserInfoPO struct { // 用户信息
	Id             int64     // 主键id
	Email          string    // email
	Nick           string    // 昵称
	PassWorld      string    // 密码
	AvatarImageUrl string    // 头像
	Status         int       // 状态
	CreatedAt      time.Time // 创建时间
	UpdatedAt      time.Time // 更新时间
	DeletedAt      time.Time // 删除时间
}

type TagInfoPO struct { // 标签信息
	Id        int64     // 主键id
	Name      string    // 标签名
	Status    int       // 状态
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt time.Time // 删除时间
}

type BlogRelationTagInfoPO struct { // blog和标签的关联关系
	Id        int64 // 主键id
	BlogId    int64 // blogID
	TagId     int64 // tagId
	Status    int   // 状态
	CreatedAt int64 // 创建时间
	UpdatedAt int64 // 更新时间
	DeletedAt int64 // 删除时间
}

type CommentInfoPO struct { // 评论信息
	Id        int64  // 评论id
	UserId    int64  // 用户id
	BlogId    int64  // blog id
	ParentId  int64  // 父评论id
	Content   string // 内容信息
	Status    int    // 状态
	CreatedAt int64  // 创建时间
	UpdatedAt int64  // 更新时间
	DeletedAt int64  // 删除时间
}

type SortType int32

const (
	SORT_BY_CREATED_DESC       SortType = iota
	SORT_BY_VIEW_COUNT_DESC    SortType = iota
	SORT_BY_COMMENT_COUNT_DESC SortType = iota
)

type BlogDao interface {
	// 查询所有的 BlogId
	QueryAllBlogIds(sortType ...SortType) []int64

	// 根据Tag查询所有的 BlogId
	QueryAllBlogIdsWithTagFilter(tagIds []int64, sortType ...SortType) []int64

	// 根据id查询所有的blog信息
	QueryBlogInfoByIds(blogIds []int64) []BlogInfoPO
}

type BlogDaoImpl struct {
}

func (blogDaoImpl *BlogDaoImpl) QueryBlogInfoByIds(blogIds []int64) ([]BlogInfoPO, error) {
	r := make([]BlogInfoPO, 0)
	if blogIds == nil || len(blogIds) <= 0 {
		return r, nil
	}

	sql := "select id, user_id, title, summary, content, view_count, comment_count, status, created_at, updated_at, deleted_at from blog_info where status = 1 and id in "
	sql += "("
	sql += utils.PlaceHolders(len(blogIds))
	sql += ")"

	rows, err := GetDB().Query(sql, 1)

	if err != nil {
		return r, err
	}

	for rows.Next() {
		blogInfo := new(BlogInfoPO)
		err := rows.Scan(&blogInfo.Id, &blogInfo.UserID, &blogInfo.Title, &blogInfo.Summary, &blogInfo.Content, &blogInfo.ViewCount, &blogInfo.Content, &blogInfo.Status, &blogInfo.CreatedAt, &blogInfo.UpdatedAt, &blogInfo.DeletedAt)
		if err != nil {
			return r, err
		}

		r = append(r, *blogInfo)
	}

	return r, nil
}

func (blogDapImpl *BlogDaoImpl) QueryAllBlogIds(sortTypes ...SortType) ([]int64, error) {
	sql := "select id from blog_info where status = 1"

	sql += " order by created_at desc"

	for _, sortType := range sortTypes {
		switch sortType {
		case SORT_BY_VIEW_COUNT_DESC:
			sql += ", view_count desc"
		case SORT_BY_COMMENT_COUNT_DESC:
			sql += ", comment_count desc"
		}
	}

	rows, err := GetDB().Query(sql)

	if err != nil {
		return []int64{}, err
	}

	ids := make([]int64, 0)
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		ids = append(ids, id)
	}

	return ids, nil
}

/*
rows, err := GetDB().Query(

*/
