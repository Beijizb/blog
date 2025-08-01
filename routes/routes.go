package routes

import (
	"github.com/gin-gonic/gin"

	"blog/controllers"
	"blog/middleware"
)

func Register(r *gin.Engine) {
	// front routes
	r.GET("/", controllers.Home)
	r.GET("/post/:slug", controllers.ViewPost)
	r.GET("/category/:slug", controllers.CategoryPosts)

	admin := r.Group("/admin")
	{
		admin.GET("/login", controllers.ShowLogin)
		admin.POST("/login", controllers.Login)
		admin.GET("/logout", controllers.Logout)
	}

	protected := r.Group("/admin")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/dashboard", controllers.Dashboard)
		protected.GET("/posts", controllers.AdminPosts)
		protected.GET("/posts/new", controllers.NewPost)
		protected.POST("/posts", controllers.CreatePost)
		protected.GET("/posts/edit/:id", controllers.EditPost)
		protected.POST("/posts/update/:id", controllers.UpdatePost)
		protected.GET("/posts/delete/:id", controllers.DeletePost)

		protected.GET("/categories", controllers.Categories)
		protected.POST("/categories", controllers.CreateCategory)
		protected.GET("/categories/delete/:id", controllers.DeleteCategory)
	}
}
