package handlers

import (
	"strconv"

	"github.com/dingdinglz/test-blog/models"
	"github.com/dingdinglz/test-blog/services"
	"github.com/dingdinglz/test-blog/utils"
	"github.com/gin-gonic/gin"
)

// CreateArticleRequest 创建文章请求
type CreateArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Create 创建文章
func CreateArticle(c *gin.Context) {
	var req CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误: "+err.Error())
		return
	}

	// 从Context获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, "未授权")
		return
	}

	// 调用服务层
	article, err := services.CreateArticle(req.Title, req.Content, userID.(uint))
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	// 返回响应
	utils.Success(c, models.ArticleResponse{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		UserID:  article.UserID,
		Author: models.UserResponse{
			ID:       article.User.ID,
			Username: article.User.Username,
			Email:    article.User.Email,
		},
		CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, "创建成功")
}

// GetAll 获取所有文章
func GetAllArticles(c *gin.Context) {
	// 调用服务层
	articles, err := services.GetAllArticles()
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	// 构建响应
	var response []models.ArticleResponse
	for _, article := range articles {
		response = append(response, models.ArticleResponse{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			UserID:  article.UserID,
			Author: models.UserResponse{
				ID:       article.User.ID,
				Username: article.User.Username,
				Email:    article.User.Email,
			},
			CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	utils.Success(c, response, "success")
}

// GetByUser 获取指定用户的文章
func GetArticlesByUser(c *gin.Context) {
	// 获取用户ID参数
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		utils.Error(c, 400, "无效的用户ID")
		return
	}

	// 调用服务层
	articles, err := services.GetUserArticles(uint(userID))
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	// 构建响应
	var response []models.ArticleResponse
	for _, article := range articles {
		response = append(response, models.ArticleResponse{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			UserID:  article.UserID,
			Author: models.UserResponse{
				ID:       article.User.ID,
				Username: article.User.Username,
				Email:    article.User.Email,
			},
			CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	utils.Success(c, response, "success")
}

// GetArticleByID 根据ID获取文章
func GetArticleByID(c *gin.Context) {
	// 获取文章ID参数
	articleIDStr := c.Param("id")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 32)
	if err != nil {
		utils.Error(c, 400, "无效的文章ID")
		return
	}

	// 调用服务层
	article, err := services.GetArticleByID(uint(articleID))
	if err != nil {
		utils.Error(c, 404, err.Error())
		return
	}

	// 构建响应
	response := models.ArticleResponse{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		UserID:  article.UserID,
		Author: models.UserResponse{
			ID:       article.User.ID,
			Username: article.User.Username,
			Email:    article.User.Email,
		},
		CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.Success(c, response, "success")
}

// Update 更新文章
func UpdateArticle(c *gin.Context) {
	// 获取文章ID参数
	articleIDStr := c.Param("id")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 32)
	if err != nil {
		utils.Error(c, 400, "无效的文章ID")
		return
	}

	var req UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误: "+err.Error())
		return
	}

	// 从Context获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, "未授权")
		return
	}

	// 调用服务层
	article, err := services.UpdateArticle(uint(articleID), userID.(uint), req.Title, req.Content)
	if err != nil {
		utils.Error(c, 403, err.Error())
		return
	}

	// 返回响应
	utils.Success(c, models.ArticleResponse{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		UserID:  article.UserID,
		Author: models.UserResponse{
			ID:       article.User.ID,
			Username: article.User.Username,
			Email:    article.User.Email,
		},
		CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, "更新成功")
}

// Delete 删除文章
func DeleteArticle(c *gin.Context) {
	// 获取文章ID参数
	articleIDStr := c.Param("id")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 32)
	if err != nil {
		utils.Error(c, 400, "无效的文章ID")
		return
	}

	// 从Context获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, "未授权")
		return
	}

	// 调用服务层
	err = services.DeleteArticle(uint(articleID), userID.(uint))
	if err != nil {
		utils.Error(c, 403, err.Error())
		return
	}

	utils.Success(c, nil, "删除成功")
}
