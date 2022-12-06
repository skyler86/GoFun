package routers

import (
	"Blog/controllers"
	"Blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	// 管理后台
	// 管理员登陆
	r.POST("/login", controllers.AuthLogin)

	// 管理员路由组
	userGroup := r.Group("/user/v1")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.POST("/create-user", controllers.CreateUser)
		userGroup.GET("/get-users", controllers.GetUsers)
		userGroup.POST("/del-user-by-id", controllers.DeleteUserById)
		userGroup.GET("/get-user-by-id", controllers.GetUserById)
		userGroup.POST("/update-user-by-id", controllers.UpdateUserById)
		userGroup.POST("/disable-user-by-id", controllers.DisableUserById)
		userGroup.POST("/enable-user-by-id", controllers.EnableUserById)

	}

	// 标签路由组
	tagsGroup := r.Group("/tags/v1")
	tagsGroup.Use(middleware.AuthMiddleware())
	{
		tagsGroup.POST("/create-tags", controllers.CreateTags)
		tagsGroup.GET("/get-tags-list", controllers.GetTagsList)
		tagsGroup.POST("/update-Tags-by-id", controllers.UpdateTagsById)
		tagsGroup.POST("/del-tags-by-id", controllers.DeleteTagsById)
	}

	// 分类路由组
	cateGroup := r.Group("/cate/v1")
	cateGroup.Use(middleware.AuthMiddleware())
	{
		cateGroup.POST("/create-cate", controllers.CreateCate)
		cateGroup.GET("/get-cate-list", controllers.GetCateList)
		cateGroup.POST("/update-cate-by-id", controllers.UpdateCateById)
		cateGroup.POST("/del-cate-by-id", controllers.DeleteCateById)
	}

	// 文章路由组
	postsGroup := r.Group("/posts/v1")
	postsGroup.Use(middleware.AuthMiddleware())
	{
		postsGroup.POST("/create-post", controllers.CreatePost)
		postsGroup.GET("/get-posts-list", controllers.GetPostsList)
		postsGroup.POST("/update-post-by-id", controllers.UpdatePostById)
		postsGroup.POST("/del-post-by-id", controllers.DeletePostById)
	}

	// 文章评论路由组
	commentGroup := r.Group("/comment/v1")
	{
		commentGroup.POST("/create-comment", controllers.CreateComment)
		commentGroup.GET("/get-comment-list", controllers.GetCommentList)
		commentGroup.POST("/del-comment-by-id", controllers.DelCommentById)
	}

	// 友联路由组
	linkGroup := r.Group("/link/v1")
	{
		linkGroup.POST("/create-link", controllers.CreateLink)
		linkGroup.GET("/get-link-list", controllers.GetLinkList)
		linkGroup.POST("/update-link-by-id", controllers.UpdateLinkById)
		linkGroup.POST("/del-link-by-id", controllers.DeleteLinkById)
	}

	r.Run(":8080")
}
