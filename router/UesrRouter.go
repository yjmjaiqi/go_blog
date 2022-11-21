package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/go_blog/db"
	"net/http"
	"strconv"
)

//用户注册get/post
func GetUserRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "userRegister.html", nil)
}
func PostUserRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")
	fmt.Println(username, password, rePassword)
	user, _, _ := db.GetSelectUser(username, password)
	if password == rePassword && user == "" {
		db.InsertUser(username, password)
		c.HTML(http.StatusOK, "userLogin.html", nil)
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "注册失败，密码不匹配",
			"res":     "用户注册",
			"href":    "/",
		})
	}
}

//用户登录get/post
func GetUserLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "userLogin.html", nil)
}

var userName string
var passWord string

func PostUserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, passWord)
	user, pass, id := db.GetSelectUser(username, password)
	userName = user
	passWord = pass
	fmt.Println(user, pass)
	amount, num := db.BlogAmount(id)
	blog := db.SelectArticle(id)
	if username == user && password == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "登录失败，用户名或密码不匹配",
			"res":     "用户登录",
			"href":    "/userLogin",
		})
	}
}

//博客页
func Blog(c *gin.Context) {
	ids, _ := strconv.Atoi(c.Param("id"))
	user, pass, id := db.GetSelectUser(userName, passWord)
	db.UpdateClick(ids)
	amount, num := db.BlogAmount(id)
	blog := db.SelectArticle(id)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "",
			"href":    "/userLogin",
		})
	}
}

//增加博文get/post
func GetAddArticle(c *gin.Context) {
	user, pass, _ := db.GetSelectUser(userName, passWord)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "addArticle.html", gin.H{
			"user": user,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "发表失败",
			"href":    "/blog",
		})
	}
}
func PostAddArticle(c *gin.Context) {
	blogname := c.PostForm("blogname")
	content := c.PostForm("content")
	user, pass, id := db.GetSelectUser(userName, passWord)
	db.InsertArticle(user, blogname, content, id)
	amount, num := db.BlogAmount(id)
	blog := db.SelectArticle(id)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "发表失败",
			"res":     "",
			"href":    "/blog",
		})
	}
}

//评论
func Comment(c *gin.Context) {
	blogid, _ := strconv.Atoi(c.PostForm("id"))
	blogname := c.PostForm("blogname")
	comment := c.PostForm("comment")
	fmt.Println(blogid, blogname, comment)
	user, pass, id := db.GetSelectUser(userName, passWord)
	db.InsertComment(user, blogname, comment, blogid, id)
	db.UpdateArticle(user, comment, blogid)
	amount, num := db.BlogAmount(id)
	blog := db.SelectArticle(id)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "评论失败",
			"href":    "/blog",
		})
	}
}

//博客圈
func Blogs(c *gin.Context) {
	user, pass, _ := db.GetSelectUser(userName, passWord)
	ids, _ := strconv.Atoi(c.Param("id"))
	db.UpdateClick(ids)
	amounts, num := db.BlogAmounts()
	blog := db.SelectBlogs()
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blogs.html", gin.H{
			"amounts": amounts,
			"num":     num,
			"blog":    blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "进入博客圈失败",
			"res":     "",
			"href":    "/blog",
		})
	}
}

//评论
func BlogsComment(c *gin.Context) {
	blogid, _ := strconv.Atoi(c.PostForm("id"))
	blogname := c.PostForm("blogname")
	comment := c.PostForm("comment")
	fmt.Println(blogid, blogname, comment)
	user, pass, id := db.GetSelectUser(userName, passWord)
	db.InsertComment(user, blogname, comment, blogid, id)
	db.UpdateArticle(user, comment, blogid)
	amounts, num := db.BlogAmounts()
	blog := db.SelectBlogs()
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blogs.html", gin.H{
			"amounts": amounts,
			"num":     num,
			"blog":    blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "评论失败",
			"href":    "/blogs",
		})
	}
}

// 博客 模糊 查询
func SearchBlog(c *gin.Context) {
	blogname := c.PostForm("blogname")
	fmt.Println(blogname)
	user, pass, id := db.GetSelectUser(userName, passWord)
	blog := db.SearchBlog(blogname, id)
	amount, num := db.BlogAmount(id)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "查询失败",
			"href":    "/blog",
		})
	}
}

// 博客圈 模糊 查询
func SearchBlogs(c *gin.Context) {
	blogname := c.PostForm("blogname")
	user, pass, _ := db.GetSelectUser(userName, passWord)
	blog := db.SearchBlogs(blogname)
	amounts, num := db.BlogAmounts()
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blogs.html", gin.H{
			"amounts": amounts,
			"num":     num,
			"blog":    blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "查询失败",
			"href":    "/blogs",
		})
	}
}

//修改博客 get/post
func GetAlterBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) //博客id(点击量)
	user, pass, _ := db.GetSelectUser(userName, passWord)
	comment := db.SelectComment(id)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "alterBlog.html", gin.H{
			"user":    user,
			"id":      id,
			"comment": comment,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "",
			"href":    "/blog",
		})
	}
}
func PostAlterBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id")) //博客id(修改博客)
	blogName := c.PostForm("blogName")
	blogname := c.PostForm("blogname")
	content := c.PostForm("content")
	comment := c.PostForm("comment")
	db.UpdateBlog(id, blogName, blogname, content, comment)
	db.UpdateClick(id)
	user, pass, userid := db.GetSelectUser(userName, passWord)
	amount, num := db.BlogAmount(userid)
	blog := db.SelectArticle(userid)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "",
			"href":    "/alterBlog",
		})
	}
}

// 删除博客
func DeleteBlog(c *gin.Context) {
	deleteId, _ := strconv.Atoi(c.Param("id"))
	db.DeleteBlog(deleteId)
	user, pass, id := db.GetSelectUser(userName, passWord)
	db.UpdateClick(deleteId)
	amount, num := db.BlogAmount(id)
	blog := db.SelectArticle(id)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "",
			"href":    "/userLogin",
		})
	}
}

//  博客 分页
func NextPage(c *gin.Context) {
	nums, _ := strconv.Atoi(c.Param("id"))
	user, pass, userid := db.GetSelectUser(userName, passWord)
	amount, num := db.BlogAmount(userid)
	blog := db.NextPage(userid, nums)
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"user":   user,
			"amount": amount,
			"num":    num,
			"blog":   blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "",
			"href":    "/blog",
		})
	}
}

//  博客圈 分页
func NextPages(c *gin.Context) {
	nums, _ := strconv.Atoi(c.Param("id"))
	user, pass, _ := db.GetSelectUser(userName, passWord)
	blog := db.NextPages(nums)
	amounts, num := db.BlogAmounts()
	if userName == user && passWord == pass {
		c.HTML(http.StatusOK, "blogs.html", gin.H{
			"amounts": amounts,
			"num":     num,
			"blog":    blog,
		})
	} else {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"content": "未获取到相关内容",
			"res":     "",
			"href":    "/blogs",
		})
	}
}
