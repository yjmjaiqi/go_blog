package db

import (
	"fmt"
)

//更新评论
func UpdateArticle(username string, comment string, id int) {
	str := username + ":  " + comment
	fmt.Println(str)
	_, err := Db.Exec("update article set comment=concat (comment,?,';') where id=?", str, id)
	if err != nil {
		fmt.Println("UpdateArticle update", err)
		return
	}
	return
}

//点击量
func UpdateClick(id int) {
	_, err := Db.Exec("update article set click=click+1  where id=?", id)
	if err != nil {
		fmt.Println("UpdateClick update", err)
		return
	}
	return
}

//修改博客
func UpdateBlog(blogid int, blogName string, blogname string, content string, comment string) {
	_, err := Db.Exec("update article set blogname=?,content=?,comment=?,altertime=current_timestamp where id=?", blogname, content, comment, blogid)
	if err != nil {
		fmt.Println("UpdateBlog update", err)
		return
	}
	_, err = Db.Exec("update comment set blogname=?,comment=? where blogname=? and id=?", blogname, comment, blogName, blogid)
	if err != nil {
		fmt.Println("UpdateBlog comment update", err)
		return
	}
	return
}
