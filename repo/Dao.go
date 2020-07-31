package repo

import (
	"strings"
	"summoner/sconsts"
	"summoner/utils"
)

type BlogDao interface {
	// 查询所有的 BlogId
	QueryAllBlogIds(sortType ...sconsts.SortType) []int64

	// 根据Tag查询所有的 BlogId
	QueryAllBlogIdsWithTagFilter(tagIds []int64, sortType ...sconsts.SortType) []int64

	// 根据id查询所有的blog信息
	QueryBlogInfoByIds(blogIds []int64) []BlogInfoPO
}

type BlogDaoImpl struct {
}

func (blogDaoImpl *BlogDaoImpl) AddBlogInfo(blogInfo BlogInfoPO) (int64, error) {
	sql := "insert into blog_info ( user_id, title, summary, content, view_count, comment_count, status ) values ( ?, ?, ?, ?, ?, ?, ? )"
	result, err := GetDB().Exec(sql, blogInfo.Title, blogInfo.Summary, blogInfo.Content, blogInfo.ViewCount, blogInfo.CommentCount, blogInfo.Status)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (blogDaoImpl *BlogDaoImpl) QueryBlogInfoByIds(blogIds ...int64) ([]BlogInfoPO, error) {
	r := make([]BlogInfoPO, 0)
	if blogIds == nil || len(blogIds) <= 0 {
		return r, nil
	}

	sql := "select id, user_id, title, summary, content, view_count, comment_count, status, created_at, updated_at, deleted_at from blog_info where status = 1 and id in "
	sql += "("
	sql += strings.Join(utils.ChangeIntListToStringList(blogIds), ",")
	sql += ")"

	rows, err := GetDB().Query(sql)

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

func (blogDaoImpl *BlogDaoImpl) QueryAllBlogIds(sortTypes ...sconsts.SortType) ([]int64, error) {
	sql := "select id from blog_info where status = 1"

	sql += " order by created_at desc"

	for _, sortType := range sortTypes {
		switch sortType {
		case sconsts.SortByViewCountDesc:
			sql += ", view_count desc"
		case sconsts.SortByCommentCountDesc:
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
