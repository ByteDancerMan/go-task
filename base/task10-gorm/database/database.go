// database/database.go
package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"go-task/base/task10-gorm/models"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	// 自动迁移模型到数据库表
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		return nil, err
	}

	// 初始化已有文章的评论状态
	initializePostCommentStatus(db)
	
	return db, nil
}

// initializePostCommentStatus 初始化已有文章的评论状态
func initializePostCommentStatus(db *gorm.DB) {
	// 查询所有文章
	var posts []models.Post
	db.Find(&posts)
	
	for _, post := range posts {
		// 计算每篇文章的评论数量
		var count int64
		db.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&count)
		
		// 根据评论数量更新状态
		var status string
		if count > 0 {
			status = "有评论"
		} else {
			status = "无评论"
		}
		
		// 更新文章状态
		db.Model(&models.Post{}).Where("id = ?", post.ID).Update("comment_status", status)
	}
}