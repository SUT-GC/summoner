package repo

type BlogInfoPO struct { // blog信息
	Id           int64  // 主键id
	UserID       int64  // 用户id
	Title        string // 标题
	Summary      string // 摘要
	Content      string // 内容
	ViewCount    int64  // 浏览计数
	CommentCount int64  // 评论计数
	Status       int    // 状态
	CreatedAt    int64  // 创建时间
	UpdatedAt    int64  // 更新时间
	DeletedAt    int64  // 删除时间
}

type UserInfoPO struct { // 用户信息
	Id             int64  // 主键id
	Email          string // email
	Nick           string // 昵称
	PassWorld      string // 密码
	AvatarImageUrl string // 头像
	Status         int    // 状态
	CreatedAt      int64  // 创建时间
	UpdatedAt      int64  // 更新时间
	DeletedAt      int64  // 删除时间
}
