package repo

import "time"

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
