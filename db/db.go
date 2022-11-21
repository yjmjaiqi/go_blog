package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db = Connect()

func Connect() *sql.DB {
	Db, err := sql.Open("mysql", "root:123456@(localhost:3306)/register?charset='utf8'&&parseTime=true")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return nil
	}
	user := `create table if not exists user(
		id int primary key auto_increment,
		username varchar(30),
		password varchar(30)
	);`
	article := `create table if not exists article(
		id int primary key auto_increment,
		username varchar(30) default '',
		blogname varchar(255) default '',
		content longtext,
		comment longtext,
		click  int(11) default 0,
		blogtime varchar(30) default '',
		altertime varchar(30) default '',
        userid int(11),
        foreign key (userid) references user(id)
	);`
	comment := `create table if not exists comment(
		id int primary key auto_increment,
		username varchar(30) default '',
		blogname varchar(255) default '',
		comment longtext,
        blogid int(11),
        userid int(11),
        foreign key (blogid) references article(id),
        foreign key (userid) references user(id)
	);`
	_, err = Db.Exec(user)
	if err != nil {
		fmt.Println("Db.Exec(user)", err)
		return nil
	}
	_, err = Db.Exec(article)
	if err != nil {
		fmt.Println("Db.Exec(article)", err)
		return nil
	}
	_, err = Db.Exec(comment)
	if err != nil {
		fmt.Println("Db.Exec(comment)", err)
		return nil
	}
	return Db
}
