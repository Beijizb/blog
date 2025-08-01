package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"blog/config"
	"blog/models"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/dashboard.html", nil)
}

func AdminPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Preload("Category").Find(&posts)
	c.HTML(http.StatusOK, "admin/posts.html", gin.H{"posts": posts})
}

func NewPost(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.HTML(http.StatusOK, "admin/form.html", gin.H{"categories": categories})
}

func CreatePost(c *gin.Context) {
	title := c.PostForm("title")
	slug := c.PostForm("slug")
	content := c.PostForm("content")
	categoryID := c.PostForm("category_id")

	var post models.Post
	post.Title = title
	post.Slug = slug
	post.Content = content
	post.CategoryID = parseUint(categoryID)
	post.UserID = getUserID(c)

	file, _ := c.FormFile("cover")
	if file != nil {
		filename := filepath.Base(file.Filename)
		path := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(file, path); err == nil {
			post.Cover = "/" + path
		}
	}
	config.DB.Create(&post)
	c.Redirect(http.StatusFound, "/admin/posts")
}

func EditPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if config.DB.First(&post, id).Error != nil {
		c.String(http.StatusNotFound, "Post not found")
		return
	}
	var categories []models.Category
	config.DB.Find(&categories)
	c.HTML(http.StatusOK, "admin/form.html", gin.H{"post": post, "categories": categories})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if config.DB.First(&post, id).Error != nil {
		c.String(http.StatusNotFound, "Post not found")
		return
	}
	post.Title = c.PostForm("title")
	post.Slug = c.PostForm("slug")
	post.Content = c.PostForm("content")
	post.CategoryID = parseUint(c.PostForm("category_id"))
	post.UpdatedAt = time.Now()

	file, _ := c.FormFile("cover")
	if file != nil {
		filename := filepath.Base(file.Filename)
		path := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(file, path); err == nil {
			post.Cover = "/" + path
		}
	}

	config.DB.Save(&post)
	c.Redirect(http.StatusFound, "/admin/posts")
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Post{}, id)
	c.Redirect(http.StatusFound, "/admin/posts")
}

func Categories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.HTML(http.StatusOK, "admin/categories.html", gin.H{"categories": categories})
}

func CreateCategory(c *gin.Context) {
	name := c.PostForm("name")
	slug := c.PostForm("slug")
	config.DB.Create(&models.Category{Name: name, Slug: slug})
	c.Redirect(http.StatusFound, "/admin/categories")
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Category{}, id)
	c.Redirect(http.StatusFound, "/admin/categories")
}

func parseUint(s string) uint {
	var i uint
	fmt.Sscanf(s, "%d", &i)
	return i
}

func getUserID(c *gin.Context) uint {
	session := sessions.Default(c)
	id := session.Get("user_id")
	if id == nil {
		return 0
	}
	return id.(uint)
}
