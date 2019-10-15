package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "user/routers"
)

func main() {
	orm.RegisterDataBase("default","mysql","root:daomei,815@/beego_student?charset=utf8")
	//o:=orm.NewOrm()
	//var user1 models.User
	//user1.Id=1
	//o.Delete(&user1)
	beego.Run()
}

