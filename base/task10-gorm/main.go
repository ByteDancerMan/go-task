package main

import (
	"fmt"
	"go-task/base/task10-gorm/database"
	"go-task/base/task10-gorm/models"
	"go-task/base/task10-gorm/queries"
)

func main() {
	// 连接数据库
	db, err := database.ConnectDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	// // 创建一个用户
	// user := &models.User{
	// 	Username: "test",
	// 	Email:    "<EMAIL>",
	// 	Password: "password",
	// }

	// // 创建一篇文章
	// post := &models.Post{
	// 	Title:   "测试文章",
	// 	Content: "这是测试文章的内容",
	// }

	// // 插入一条示例数据
	// err = queries.CreateUserWithPosts(db, user, []*models.Post{post})
	// if err != nil {
	// 	fmt.Printf("插入数据失败: %v\n", err)
	// }

	// // 插入一条评论
	// comment := &models.Comment{
	// 	PostID:  1,
	// 	UserID:  1,
	// 	Content: "这是测试文章的评论",
	// }

	// err = queries.CreateComment(db, comment)
	// if err != nil {
	// 	fmt.Printf("插入评论失败: %v\n", err)
	// }

	// // 示例：查询ID为1的用户发布的所有文章及评论
	// posts, err := queries.QueryUserPostsWithComments(db, 1)
	// if err != nil {
	// 	fmt.Printf("查询用户文章失败: %v\n", err)
	// } else {
	// 	fmt.Printf("用户发布了 %d 篇文章\n", len(posts))
	// 	for _, post := range posts {
	// 		fmt.Printf("文章: %s, 评论数: %d\n", post.Title, len(post.Comments))
	// 	}
	// }
	
	// // 示例：查询评论数最多的文章
	// post, count, err := queries.QueryPostWithMostComments(db)
	// if err != nil {
	// 	fmt.Printf("查询热门文章失败: %v\n", err)
	// } else {
	// 	fmt.Printf("评论最多的文章: %s, 评论数: %d\n", post.Title, count)
	// }



	// 创建一个用户
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	// 先检查用户是否已存在，避免重复创建
	var existingUser models.User
	result := db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error != nil {
		// 用户不存在，创建新用户
		err = queries.CreateUser(db, user)
		if err != nil {
			fmt.Printf("创建用户失败: %v\n", err)
			return
		}
		fmt.Printf("创建用户成功，用户ID: %d\n", user.ID)
	} else {
		// 用户已存在，使用现有用户
		user = &existingUser
		fmt.Printf("使用现有用户，用户ID: %d\n", user.ID)
	}

	// 创建一篇文章
	post := &models.Post{
		Title:   "测试文章",
		Content: "这是测试文章的内容",
		UserID:  user.ID,
	}

	err = queries.CreatePost(db, post)
	if err != nil {
		fmt.Printf("创建文章失败: %v\n", err)
		return
	}
	fmt.Printf("创建文章成功，文章ID: %d\n", post.ID)

	// 验证用户的文章数量是否已更新
	var updatedUser models.User
	db.First(&updatedUser, user.ID)
	fmt.Printf("用户文章数量: %d\n", updatedUser.PostCount)

	// 创建几条评论
	comment1 := &models.Comment{
		Content: "第一条测试评论",
		PostID:  post.ID,
		UserID:  user.ID,
	}

	comment2 := &models.Comment{
		Content: "第二条测试评论",
		PostID:  post.ID,
		UserID:  user.ID,
	}

	err = queries.CreateComment(db, comment1)
	if err != nil {
		fmt.Printf("创建评论1失败: %v\n", err)
	}
	
	// 等待一段时间确保钩子函数执行完成
	// 验证文章的评论状态
	var updatedPost models.Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("创建一条评论后，文章评论状态: %s\n", updatedPost.CommentStatus)

	err = queries.CreateComment(db, comment2)
	if err != nil {
		fmt.Printf("创建评论2失败: %v\n", err)
	}

	// 再次检查文章状态
	db.First(&updatedPost, post.ID)
	fmt.Printf("创建两条评论后，文章评论状态: %s\n", updatedPost.CommentStatus)

	// 删除评论
	err = db.Delete(&models.Comment{}, comment1.ID).Error
	if err != nil {
		fmt.Printf("删除评论失败: %v\n", err)
	}

	// 再次检查文章状态
	db.First(&updatedPost, post.ID)
	fmt.Printf("删除一条评论后，文章评论状态: %s\n", updatedPost.CommentStatus)

	// 删除所有评论 - 逐个删除以触发钩子函数
	var comments []models.Comment
	db.Where("post_id = ?", post.ID).Find(&comments)
	for _, comment := range comments {
		db.Delete(&comment) // 这样会触发每个评论的 AfterDelete 钩子
	}

	// 最后检查文章状态
	db.First(&updatedPost, post.ID)
	fmt.Printf("删除所有评论后，文章评论状态: %s\n", updatedPost.CommentStatus)
}