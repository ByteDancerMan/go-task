package queries

import (
	"gorm.io/gorm"
	"go-task/base/task10-gorm/models"
)


// QueryUserPostsWithComments 查询用户发布的所有文章及对应评论
func QueryUserPostsWithComments(db *gorm.DB, userID uint) ([]models.Post, error) {
	var posts []models.Post
	
	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	
	return posts, nil
}

// QueryPostWithMostComments 查询评论数最多的文章
func QueryPostWithMostComments(db *gorm.DB) (*models.Post, int64, error) {
	var post models.Post
	var result struct {
		PostID       uint
		CommentCount int64
	}
	
	err := db.Model(&models.Comment{}).
		Select("post_id, COUNT(*) as comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error
		
	if err != nil {
		return nil, 0, err
	}
	
	err = db.First(&post, result.PostID).Error
	if err != nil {
		return nil, 0, err
	}
	
	return &post, result.CommentCount, nil
}

// CreateUser 创建新用户
func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// CreatePost 创建新文章
func CreatePost(db *gorm.DB, post *models.Post) error {
	return db.Create(post).Error
}

// CreateComment 创建新评论
func CreateComment(db *gorm.DB, comment *models.Comment) error {
	return db.Create(comment).Error
}

// CreateUserWithPosts 创建用户及其文章
func CreateUserWithPosts(db *gorm.DB, user *models.User, posts []*models.Post) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		
		// 创建文章
		for i := range posts {
			posts[i].UserID = user.ID
			if err := tx.Create(&posts[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}