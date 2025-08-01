package controllers

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"

	"blog/config"
	"blog/models"
)

func Home(c *gin.Context) {
	var posts []models.Post
	config.DB.Order("created_at desc").Preload("Category").Find(&posts)
	c.HTML(http.StatusOK, "front/index.html", gin.H{"posts": posts})
}

func ViewPost(c *gin.Context) {
	slug := c.Param("slug")
	var post models.Post
	if config.DB.Preload("Category").Where("slug = ?", slug).First(&post).Error != nil {
		c.String(http.StatusNotFound, "Post not found")
		return
	}
	md := goldmark.New()
	var buf bytes.Buffer
	if err := md.Convert([]byte(post.Content), &buf); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "front/post.html", gin.H{"post": post, "content": buf.String()})
}

func CategoryPosts(c *gin.Context) {
	slug := c.Param("slug")
	var category models.Category
	if config.DB.Where("slug = ?", slug).First(&category).Error != nil {
		c.String(http.StatusNotFound, "Category not found")
		return
	}
	var posts []models.Post
	config.DB.Where("category_id = ?", category.ID).Preload("Category").Find(&posts)
	c.HTML(http.StatusOK, "front/category.html", gin.H{"category": category, "posts": posts})
}
