package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;unique"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Posts    []Post `gorm:"foreignKey:UserID"`
	PostCount int   `gorm:"type:int;default:0"` // 添加文章数量统计字段
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(200);not null"`
	Content  string    `gorm:"type:text"`
	UserID   uint
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
	CommentCount int       `gorm:"type:int;default:0"` // 添加评论数量字段
	CommentStatus string   `gorm:"type:varchar(50);default:'无评论'"` // 添加评论状态字段
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint
	UserID  uint
	Post    Post `gorm:"foreignKey:PostID"`
	User    User `gorm:"foreignKey:UserID"`
}


// BeforeCreate 是 Post 模型的钩子函数，在创建文章时更新用户的文章数量统计
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 增加用户的 PostCount 字段
	err := tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
	return err
}

// AfterCreate 是 Comment 模型的钩子函数，在创建评论后更新文章的评论状态
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	// 更新文章的评论状态为"有评论"
	err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "有评论").Error
	return err
}

// AfterDelete 是 Comment 模型的钩子函数，在删除评论后检查文章的评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 查询文章当前的评论数量
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

	// 根据评论数量更新文章状态
	if count == 0 {
		err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
		return err
	}
	
	return nil
}