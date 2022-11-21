package main

import (
	"github.com/gin-gonic/gin"
	"go_code/go_blog/db"
	"go_code/go_blog/router"
)

func main() {
	db.Connect()
	Router := gin.Default()
	Router.LoadHTMLGlob("./html/*")
	Router.Static("../css", "./static/css/")
	Router.Static("../js", "./static/js/")
	Router.Static("../img", "./static/img/")
	//注册
	Router.GET("/", router.GetUserRegister)
	Router.POST("/userRegister", router.PostUserRegister)
	//登录
	Router.GET("/userLogin", router.GetUserLogin)
	Router.POST("/userLogin", router.PostUserLogin)
	//博客
	Router.GET("/blog", router.Blog)
	Router.GET("/blog/:id", router.Blog)
	Router.GET("/blogs", router.Blogs)
	Router.GET("/blogs/:id", router.Blogs)
	Router.GET("/addArticle", router.GetAddArticle)
	Router.POST("/addArticle", router.PostAddArticle)
	//评论
	Router.POST("/comment", router.Comment)
	Router.POST("/blogsComment", router.BlogsComment)
	// 查询
	Router.POST("/searchBlog", router.SearchBlog)
	Router.POST("/searchBlogs", router.SearchBlogs)
	//修改blog
	Router.GET("/alterBlog/:id", router.GetAlterBlog)
	Router.POST("/alterBlog", router.PostAlterBlog)
	//分页
	Router.GET("/page/:id", router.NextPage)   //用户博客
	Router.GET("/pages/:id", router.NextPages) //所有博客
	//删除博客
	Router.GET("/delete/:id", router.DeleteBlog)
	Router.Run()
}
