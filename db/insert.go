package db

import (
	"fmt"
)

//用户注册
func InsertUser(userName string, Password string) {
	_, err := Db.Exec("insert into user (username,password) values (?,?)", userName, Password)
	if err != nil {
		fmt.Println("InsertUser", err)
		return
	}
	return
}

//增加博文
func InsertArticle(username string, blogname string, content string, id int) {
	_, err := Db.Exec("insert into article(username,blogname,content,comment,blogtime,userid) values(?,?,?,'',current_timestamp,?)", username, blogname, content, id)
	if err != nil {
		fmt.Println("InsertArticle insert", err)
		return
	}
	return
}

//插入评论
func InsertComment(username string, blogname string, comment string, blogid int, userid int) {
	str := username + ":  " + comment
	_, err := Db.Exec("insert into comment (username,blogname,comment,blogid,userid) values(?,?,?,?,?)", username, blogname, str, blogid, userid)
	if err != nil {
		fmt.Println("InsertComment insert", err)
		return
	}
	return
}
