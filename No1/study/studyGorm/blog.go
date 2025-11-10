package studygorm

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type Users struct {
	ID        uint64
	Email     string `gorm:"type:varchar(125);not null"`
	Name      string `gorm:"type:varchar(125);not null"`
	Password  string `gorm:"type:varchar(125);not null"`
	PostCount uint64 `gorm:"column:post_count"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Posts struct {
	ID            uint64
	Title         string `gorm:"type:varchar(255);not null"`
	Content       string `gorm:"type:text;not null"`
	UserId        uint64 `gorm:"column:user_id"`
	CommentCount  uint64 `gorm:"column:comment_count"`
	CommentStatus string `gorm:"type:varchar(50);default:'no_comments'"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

const (
	// 有评论
	CommentStatusNoComments = "no_comments"
	// 无评论
	CommentStatusHasComments = "has_comments"
)

type Comments struct {
	ID        uint64
	Content   string `gorm:"type:text;not null"`
	UserId    uint64 `gorm:"column:user_id"`
	PostId    uint64 `gorm:"column:post_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 查询出来的结果
type PostWithComments struct {
	Post     Posts
	Author   Users
	Comments []CommentWithAuthor
}
type CommentWithAuthor struct {
	Comment Comments
	Author  Users
}

// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func QueryPostsAndContnetByUserId(db *gorm.DB, name string) ([]PostWithComments, error) {
	// 提前构造好 返回参数
	var result []PostWithComments

	// 1. 先查询用户ID
	user := Users{}
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, fmt.Errorf("查询用户ID失败: %v", err)
	}

	// 2. 根据用户ID查询 []Posts
	posts := []Posts{}
	if err := db.Where("user_id = ?", user.ID).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("根据用户ID: %d,用户姓名: %s. 查询Post失败: %v", user.ID, user.Name, err)
	}

	if len(posts) == 0 {
		return result, nil
	}

	// 3. 根据 []Posts 查询每个 []Comments

	// for _, post := range posts {
	// 	// 查询每篇文章的所有评论
	// 	comments := []Comments{}
	// 	if err := db.Where("post_id = ?", post.ID).Find(&comments).Error; err != nil {
	// 		return nil, fmt.Errorf("根据用户ID: %d,用户姓名: %s, 文章ID: %d, 文章姓名: %s, 查询Content失败: %v", user.ID, user.Name, post.ID, post.Title, err)
	// 	}
	// 	// 查询每个评论的相关信息
	// 	var commentWithAuthor []CommentWithAuthor = make([]CommentWithAuthor, 0)
	// 	for _, comment := range comments {
	// 		// 为每个评论再查询评论作者的信息
	// 		var commentAuthor Users
	// 		if err := db.Where("id = ?", comment.UserId).Find(&commentAuthor).Error; err != nil {
	// 			// 如果评论作者不存在，跳过该评论或使用默认值
	// 			continue
	// 		}
	// 		commentWithAuthor = append(commentWithAuthor, CommentWithAuthor{Comment: comment, Author: commentAuthor})
	// 	}
	// 	// 每个进行装载
	// 	result = append(result, PostWithComments{Post: post, Author: user, Comments: commentWithAuthor})
	// }
	// return result, nil

	// 对上面进行优化 减少查询次数
	// 3. 批量查询所有文章的评论
	// 收集所有文章ID
	var postIds []uint64 = make([]uint64, len(posts))
	for i, post := range posts {
		postIds[i] = post.ID
	}

	var allComments []Comments
	if err := db.Where("post_id in ?", postIds).Find(&allComments).Error; err != nil {
		return nil, fmt.Errorf("查询评论失败: %v", err)
	}
	// 4. 批量查询所有评论作者
	authorIds := make(map[uint64]bool)
	for _, comment := range allComments {
		authorIds[comment.UserId] = true
	}

	var authorIDList []uint64
	for id := range authorIds {
		authorIDList = append(authorIDList, id)
	}

	var commentAuthors []Users
	if len(authorIDList) > 0 {
		if err := db.Where("id in ?", authorIDList).Find(&commentAuthors).Error; err != nil {
			return nil, fmt.Errorf("查询评论作者失败: %v", err)
		}
	}

	// 构建作者ID到作者信息的映射
	authorMap := make(map[uint64]Users)
	for _, author := range commentAuthors {
		authorMap[author.ID] = author
	}

	// 5. 按文章ID组织评论
	commentsByPostId := make(map[uint64][]CommentWithAuthor)
	for _, comment := range allComments {
		if author, exists := authorMap[comment.UserId]; exists {
			commentsByPostId[comment.PostId] = append(commentsByPostId[comment.PostId], CommentWithAuthor{
				Comment: comment,
				Author:  author,
			})
		}
	}

	// 6. 组装
	for _, post := range posts {
		result = append(result, PostWithComments{
			Post:     post,
			Author:   user,
			Comments: commentsByPostId[post.ID],
		})
	}
	return result, nil

}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryPostByCommentsMost(db *gorm.DB) (*Posts, int, error) {

	var post Posts = Posts{}
	var comments Comments = Comments{}

	var result struct {
		PostId       uint64
		CommentCount int
	}

	// 先查到评论数量最多的文章ID
	err := db.Model(&comments).
		Select("post_id, Count(*) as comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		return nil, 0, err
	}

	// 再根据文章ID查到文章信息
	err = db.Where("id = ?", result.PostId).First(&post).Error
	if err != nil {
		return nil, 0, err
	}
	return &post, result.CommentCount, nil
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Posts) AfterCreate(tx *gorm.DB) error {
	err := tx.Model(&Users{}).
		Where("id = ?", p.UserId).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error

	return err
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comments) AfterDelete(tx *gorm.DB) error {
	// 检查该文装的评论数量
	var commentCount int64
	err := tx.Model(&Comments{}).Where("post_id = ?", c.PostId).Count(&commentCount).Error
	if err != nil {
		return err
	}
	// 如果评论数量为0 更新文章评论状态为“无评论”
	if commentCount == 0 {
		err = tx.Model(&Posts{}).Where("id = ?", c.PostId).Update("comment_status", CommentStatusNoComments).Error
	}
	return err
}

func InitBlogTable(db *gorm.DB) {
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Posts{})
	db.AutoMigrate(&Comments{})
}

func InitBlogsData(db *gorm.DB) {
	// 插入User数据
	var users []Users = []Users{
		{
			ID:       1,
			Email:    "zhangsan@gmail.com",
			Name:     "张三",
			Password: "zhangsan_000",
		},
		{
			ID:       2,
			Email:    "lisi@gmail.com",
			Name:     "李四",
			Password: "lisi_000",
		},
	}

	db.Create(&users)

	// 插入 文章数据
	var post []Posts = []Posts{
		Posts{
			ID:      1,
			Title:   "1",
			Content: "111111",
			UserId:  1,
		},
		Posts{
			ID:      2,
			Title:   "2",
			Content: "222222",
			UserId:  1,
		},
		Posts{
			ID:      3,
			Title:   "3",
			Content: "333333",
			UserId:  2,
		},
	}

	db.Create(&post)

	// 插入评论
	comments := []Comments{
		{
			ID:      1,
			Content: "1 good",
			UserId:  1,
			PostId:  1,
		},
		{
			ID:      2,
			Content: "1 一般",
			UserId:  2,
			PostId:  1,
		},
		{
			ID:      3,
			Content: "1 还不错",
			UserId:  2,
			PostId:  1,
		},
		{
			ID:      4,
			Content: "2 还ok",
			UserId:  1,
			PostId:  2,
		},
	}

	db.Create(&comments)

}
