package db

//用户
type User_Yjm struct {
	ID       int    `sql:"id" json:"id" form:"id"`                   // 用户id
	Username string `sql:"username" json:"username" form:"username"` // 用户名
	Password string `sql:"password" json:"password" form:"password"` // 用户密码
	//Borrow_book string `sql:"borrow_book" json:"borrow_book" form:"borrow_book"` // 借书时间
	//Repay_book  string `sql:"repay_book" json:"repay_book" form:"repay_book"`    // 还书时间
}

//书籍
type Blog_Yjm struct {
	ID          int    `sql:"id" json:"id" form:"id"`                            // 博客id
	Username    string `sql:"username" json:"username" form:"username"`          // 博客作者
	BlogName    string `sql:"blogname" json:"blogname" form:"blogname"`          // 博客名称
	BlogContent string `sql:"blogcontent" json:"blogcontent" form:"blogcontent"` //博客内容
	Comment     string `sql:"comment" json:"comment" form:"comment"`             //评论
	Click       int    `sql:"click" json:"click" form:"click"`                   //点击量
	BlogTime    string `sql:"blogtime" json:"blogtime" form:"blogtime"`          //发布时间
	AlterTime   string `sql:"altertime" json:"altertime" form:"altertime"`       //修改时间
	Userid      int    `sql:"userid" json:"userid" form:"userid"`                //用户id
}

//用户评论
type User_Comment struct {
	Id       int    `sql:"id" json:"id" form:"id"`                   //评论id
	Username string `sql:"username" json:"username" form:"username"` //评论用户
	BlogName string `sql:"blogname" json:"blogname" form:"blogname"` // 博客名称
	Comment  string `sql:"comment" json:"comment" form:"comment"`    //评论
	Blogid   int    `sql:"blogid" json:"blogid" form:"blogid"`       //博客id
	Userid   int    `sql:"userid" json:"userid" form:"userid"`       //用户id
}

//查询评论存储
//type Comment struct {
//	Username string `sql:"username" json:"username" form:"username"` //评论用户
//	Comment  string `sql:"comment" json:"comment" form:"comment"`    //评论
//}
