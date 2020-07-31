package biz

import "time"

type BlogSummaryDTO struct { // blog 的摘要信息
	Id           int64
	Title        string
	Summary      string
	ViewCount    int64
	CommentCount int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type BlogDTO struct { // blog信息
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
