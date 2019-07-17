package models

import (
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type User struct{ //用户表
	Id 			int
	Name 		string`orm:"size(20);unique"` 		//用户名
	PassWord 	string`orm:"size(20)"`		//登陆密码
	Email 		string`orm:"size(50)"`   	//邮箱
	Active 		bool`orm:"default(false)"`    	//是否激活
	Power 		int`orm:"default(0)"`			//权限设置  0 表示未激活  1表示激活
	Address		[]*Address `orm:"reverse(many)"`

}

type Address struct { //地址表
	Id int
	Receiver string `orm:"size(20)"`  			//收件人
	Addr string		`orm:"size(50)"`  			//收件地址
	Zipcode string  `orm:"size(20)"` 			//邮编
	Phone string	 `orm:"size(20)"` 			//联系方式
	Isdefault bool	 `orm:"defalt(false)"`			//是否默认 0 为非默认  1为默认
	User *User `orm:"rel(fk)"` 	//用户ID

}


func init() {
	//先在mysql中新建一个数据库(务必设置charset=utf8),然后登录mysql中对应的数据库,其中登录时,需要对参数 mysql进行手动导包:_空格
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/pyg")
   orm.RegisterModel(new(User),new(Address))
   orm.RunSyncdb("default",false,true)
}
