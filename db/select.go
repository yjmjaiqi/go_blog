package db

import (
	"fmt"
	"math"
)

//用户登录查询判断
func GetSelectUser(user string, pass string) (username string, password string, id int) {
	rows, err := Db.Query("select id,username,password from user where username=? and password=?", user, pass)
	if err != nil {
		fmt.Println("GetSelectUser()", err)
		return
	}
	for rows.Next() {
		user := User_Yjm{}
		err = rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			fmt.Println("GetSelectUser rows.Scan", err)
			return
		}
		username = user.Username
		password = user.Password
		id = user.ID
	}
	return
}

//blog查询页面渲染评论用户
func SelectArticle(id int) (blog []Blog_Yjm) {
	rows, err := Db.Query("select * from article where userid=? order by id desc limit 3", id)
	if err != nil {
		fmt.Println("SelectArticle()", err)
		return
	}
	for rows.Next() {
		Blog := Blog_Yjm{}
		err = rows.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime, &Blog.AlterTime, &Blog.Userid)
		if err != nil {
			fmt.Println("SelectArticle() rows.Scan", err)
			return
		}
		blog = append(blog, Blog)
	}
	return
}

//博客圈
func SelectBlogs() (blog []Blog_Yjm) {
	row, err := Db.Query("select id,username,blogname,content,comment,click,blogtime from article order by blogtime desc limit 3")
	if err != nil {
		fmt.Println("SelectComment()", err)
		return
	}
	for row.Next() {
		Blog := Blog_Yjm{}
		err = row.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime)
		if err != nil {
			fmt.Println("SelectComment() rows.Scan", err)
			return
		}
		blog = append(blog, Blog)
	}
	return
}

//  blog 查询
func SearchBlog(blogname string, userid int) (blog []Blog_Yjm) {
	fmt.Println(blogname, userid)
	rows, err := Db.Query("select * from article where userid=? and blogname like '%"+blogname+"%'  order by id desc limit 3", userid)
	if err != nil {
		fmt.Println("SelectBlog()", err)
		return
	}
	for rows.Next() {
		Blog := Blog_Yjm{}
		err = rows.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime, &Blog.AlterTime, &Blog.Userid)
		if err != nil {
			fmt.Println("SelectBlog() rows.Scan", err)
			return
		}
		blog = append(blog, Blog)
	}
	return
}

//  blogs 查询
func SearchBlogs(blogname string) (blog []Blog_Yjm) {
	row, err := Db.Query("select id,username,blogname,content,comment,click,blogtime from article where blogname like '%" + blogname + "%' order by blogtime desc limit 3")
	if err != nil {
		fmt.Println("SelectComment()", err)
		return
	}
	for row.Next() {
		Blog := Blog_Yjm{}
		err = row.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime)
		if err != nil {
			fmt.Println("SelectComment() rows.Scan", err)
			return
		}
		blog = append(blog, Blog)
	}
	return
}

//查看评论
func SelectComment(blogid int) (blog []Blog_Yjm) {
	rows, err := Db.Query("select id,blogname,comment from article where id=?", blogid)
	if err != nil {
		fmt.Println("SelectComment()", err)
		return
	}
	for rows.Next() {
		Blog := Blog_Yjm{}
		err = rows.Scan(&Blog.ID, &Blog.BlogName, &Blog.Comment)
		if err != nil {
			fmt.Println("SelectComment() rows.Scan", err)
			return
		}
		blog = append(blog, Blog)
	}
	return
}

//博客篇目
var Amount int

func BlogAmount(id int) (amount int, num float64) {
	rows, err := Db.Query("select count(id)as id from article where userid=?", id)
	if err != nil {
		fmt.Println("BlogAmount()", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			fmt.Println("SelectArticle() rows.Scan", err)
			return
		}
	}
	Amount = amount
	num = math.Round(float64(amount/3) + 0.5)
	return
}

//博客总篇目
var Amounts int

func BlogAmounts() (amount int, num float64) {
	rows, err := Db.Query("select count(id)as id from article")
	if err != nil {
		fmt.Println("BlogAmount()", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			fmt.Println("SelectArticle() rows.Scan", err)
			return
		}
	}
	Amounts = amount
	num = math.Round(float64(amount/3) + 0.5)
	return
}

//  博客
//分页
func NextPage(id int, nums int) (blog []Blog_Yjm) {
	if nums >= 1 && (nums*3) <= (Amount+2) {
		rows, err := Db.Query("select * from article where userid=? order by id desc limit ?,3", id, (nums-1)*3)
		if err != nil {
			fmt.Println("SelectArticle()", err)
			return
		}
		for rows.Next() {
			Blog := Blog_Yjm{}
			err = rows.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime, &Blog.AlterTime, &Blog.Userid)
			if err != nil {
				fmt.Println("SelectArticle() rows.Scan", err)
				return
			}
			blog = append(blog, Blog)
		}
	} else {
		rows, err := Db.Query("select * from article where userid=? order by id desc limit 3", id)
		if err != nil {
			fmt.Println("SelectArticle()", err)
			return
		}
		for rows.Next() {
			Blog := Blog_Yjm{}
			err = rows.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime, &Blog.AlterTime, &Blog.Userid)
			if err != nil {
				fmt.Println("SelectArticle() rows.Scan", err)
				return
			}
			blog = append(blog, Blog)
		}
	}
	return
}

//分页 博客圈
func NextPages(nums int) (blog []Blog_Yjm) {
	if nums >= 1 && (nums*3) <= (Amounts+2) {
		rows, err := Db.Query("select * from article  order by id desc limit ?,3", (nums-1)*3)
		if err != nil {
			fmt.Println("SelectArticle()", err)
			return
		}
		for rows.Next() {
			Blog := Blog_Yjm{}
			err = rows.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime, &Blog.AlterTime, &Blog.Userid)
			if err != nil {
				fmt.Println("SelectArticle() rows.Scan", err)
				return
			}
			blog = append(blog, Blog)
		}
	} else {
		rows, err := Db.Query("select * from article order by id desc limit 3")
		if err != nil {
			fmt.Println("SelectArticle()", err)
			return
		}
		for rows.Next() {
			Blog := Blog_Yjm{}
			err = rows.Scan(&Blog.ID, &Blog.Username, &Blog.BlogName, &Blog.BlogContent, &Blog.Comment, &Blog.Click, &Blog.BlogTime, &Blog.AlterTime, &Blog.Userid)
			if err != nil {
				fmt.Println("SelectArticle() rows.Scan", err)
				return
			}
			blog = append(blog, Blog)
		}
	}
	return
}
