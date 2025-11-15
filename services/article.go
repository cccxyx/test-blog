package services

import (
	"errors"

	"github.com/dingdinglz/test-blog/database"
	"github.com/dingdinglz/test-blog/models"
	"gorm.io/gorm"
)

// CreateArticle 创建文章
func CreateArticle(title, content string, userID uint) (*models.Article, error) {
	db := database.GetDB()

	article := &models.Article{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	if err := db.Create(article).Error; err != nil {
		return nil, errors.New("创建文章失败")
	}

	// 预加载用户信息
	db.Preload("User").First(article, article.ID)

	return article, nil
}

// GetAllArticles 获取所有文章
func GetAllArticles() ([]models.Article, error) {
	db := database.GetDB()

	var articles []models.Article
	if err := db.Preload("User").Order("created_at desc").Find(&articles).Error; err != nil {
		return nil, errors.New("获取文章列表失败")
	}

	return articles, nil
}

// GetUserArticles 获取指定用户的文章
func GetUserArticles(userID uint) ([]models.Article, error) {
	db := database.GetDB()

	var articles []models.Article
	if err := db.Preload("User").Where("user_id = ?", userID).Order("created_at desc").Find(&articles).Error; err != nil {
		return nil, errors.New("获取用户文章列表失败")
	}

	return articles, nil
}

// GetArticleByID 根据ID获取文章
func GetArticleByID(articleID uint) (*models.Article, error) {
	db := database.GetDB()

	var article models.Article
	if err := db.Preload("User").First(&article, articleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("文章不存在")
		}
		return nil, errors.New("查询文章失败")
	}

	return &article, nil
}

// UpdateArticle 更新文章
func UpdateArticle(articleID, userID uint, title, content string) (*models.Article, error) {
	db := database.GetDB()

	// 查找文章
	var article models.Article
	if err := db.First(&article, articleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("文章不存在")
		}
		return nil, errors.New("查询文章失败")
	}

	// 检查权限
	if article.UserID != userID {
		return nil, errors.New("无权修改此文章")
	}

	// 更新文章
	article.Title = title
	article.Content = content

	if err := db.Save(&article).Error; err != nil {
		return nil, errors.New("更新文章失败")
	}

	// 预加载用户信息
	db.Preload("User").First(&article, article.ID)

	return &article, nil
}

// DeleteArticle 删除文章
func DeleteArticle(articleID, userID uint) error {
	db := database.GetDB()

	// 查找文章
	var article models.Article
	if err := db.First(&article, articleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("文章不存在")
		}
		return errors.New("查询文章失败")
	}

	// 检查权限
	if article.UserID != userID {
		return errors.New("无权删除此文章")
	}

	// 删除文章
	if err := db.Delete(&article).Error; err != nil {
		return errors.New("删除文章失败")
	}

	return nil
}
