package db

import (
	"fmt"
)

func DeleteBlog(blogid int) {
	_, err := Db.Exec("delete from comment where blogid=?", blogid)
	if err != nil {
		fmt.Println("DeleteBlog comment", err)
		return
	}
	_, err = Db.Exec("delete from article where id=?", blogid)
	if err != nil {
		fmt.Println("DeleteBlog article", err)
		return
	}
	return
}
