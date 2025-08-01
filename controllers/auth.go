package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"

	"blog/config"
	"blog/models"
)

func SetupSession(storeKey string, r *gin.Engine) {
	store := cookie.NewStore([]byte(storeKey))
	r.Use(sessions.Sessions("blogsession", store))
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.html", gin.H{"title": "Login"})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "admin/login.html", gin.H{"error": "Invalid credentials"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		c.HTML(http.StatusBadRequest, "admin/login.html", gin.H{"error": "Invalid credentials"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()
	c.Redirect(http.StatusFound, "/admin/dashboard")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/admin/login")
}
