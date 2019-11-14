package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-redis/redis"

	//"os"

	//"fmt"
	"strconv"
	"strings"
	"user/models"
	// "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"


)

//  UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("form_add",c.FormAdd)
	c.Mapping("Update",c.Update)
	c.Mapping("updateUser",c.UpdateUser)
	c.Mapping("Index",c.Index)
	c.Mapping("login",c.Login)
	c.Mapping("loginTest",c.LoginTest)
	c.Mapping("logout",c.Logout)
}
func (c *UserController) Login(){
	c.TplName="login.html"
}
func (c *UserController) LoginTest() {
	//c.TplName="login.html"

	client:=getRedisConnected()

	name := c.Input().Get("username")
	password := c.Input().Get("password")

	passwordInRedis,_ := client.Get(name).Result()

	if passwordInRedis == password {
		fmt.Println("这次登陆是从redis中查的密码")
		c.SetSession("userName", name)
		c.Redirect("/index", 302)
	} else  {

		o := orm.NewOrm()

		var user models.User
		qs := o.QueryTable(user)
		err1 := qs.Filter("name", name).Filter("password", password).One(&user)

		if err1 == nil {
			//fmt.Println(user.name,user.Password)
			c.SetSession("userName", name)
			c.Redirect("/index", 302)

			//使用redis缓存用户的登录名和密码

			err1 := client.Set(name, password, 0).Err()
			if err1 != nil {
				panic(err1)
			}

			val, err2 := client.Get(name).Result()
			if err2!= nil {
				panic(err2)
			}
			fmt.Println(name, val)

		} else if err1 == orm.ErrNoRows {
			str := "用户名或密码输入错误!"
			c.Data["info"] = str
			c.TplName = "login.html"
			//fmt.Println("用户名或密码输入错误")
			//fmt.Println(user.Age)

		} else if err1 == orm.ErrMissPK {
			fmt.Println("找不到主键")
			c.Redirect("/login", 302)
		}

	}
}
func (c *UserController) Logout(){
	c.DelSession("userName")
	c.Redirect("/",302)
}
func (c *UserController) Index(){

	c.Data["userName"]=c.GetSession("userName")
	if c.Data["userName"]==nil{
		c.Redirect("/login",302)
	}
	c.TplName="index.html"
}
func (c *UserController) FormAdd(){
	c.Data["userName"]=c.GetSession("userName")
	c.TplName="form_add.html"
	if c.Data["userName"]==nil{
		c.Redirect("/login",302)
	}
}
func (c *UserController) Update(){
	id := c.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(id)
	data ,_:= models.GetUserById(int64(intid))
	c.Data["list"] = data
	c.Data["userName"]=c.GetSession("userName")
	c.TplName="form_update.html"
	if c.Data["userName"]==nil{
		c.Redirect("/login",302)
	}
}
func (c *UserController) UpdateUser(){
	c.Data["userName"]=c.GetSession("userName")
	if c.Data["userName"]==nil{
		c.Redirect("/login",302)
	}
	if c.Ctx.Request.Method == "POST" {
		id := c.Input().Get("id")
		intid, _ := strconv.Atoi(id)
		u := models.User{Id: int64(intid)}
		if err := c.ParseForm(&u); err != nil {
			c.Redirect("/update/"+id , 302)
		}
		if err := models.UpdateUserById(&u); err == nil {

			client:=getRedisConnected()
			errRedisUpdate:=client.Set(u.Name,u.Password,0).Err()
			if errRedisUpdate!=nil{
				panic(errRedisUpdate)
			}

			c.Redirect("/getAll", 302)

		} else {
			c.Redirect("/update/"+id , 302)
		}
	}
	c.TplName = "table.html"
	}


// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	c.TplName="form_add.html"
	c.Data["userName"]=c.GetSession("userName")
	if c.Data["userName"]==nil{
		c.Redirect("/login",302)
	}
	//c.Redirect("form.html",302)
	var v models.User
	//json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := c.ParseForm(&v); err != nil {
		fmt.Println("转换model失败")
		c.Ctx.WriteString("转换model失败")
        fmt.Println(err)

	}
	//c.Ctx.WriteString(v.Name)
	//fmt.Println(v.Name)
	id, err := models.AddUser(&v)
	if err == nil && id > 0 {
		c.Data["userName"]=c.GetSession("userName")
		c.Redirect("/getAll", 302)
	} else if err!=nil{
		fmt.Println("第二次err添加失败")
		//c.Ctx.WriteString("第二次err添加失败")
		fmt.Println(err)
	}
	//fmt.Println(v.Name)
	//if _, err := models.AddUser(&v); err == nil {
		//c.Ctx.Output.SetStatus(201)
		//c.Data["json"] = v
	//} else {
		//c.Data["json"] = err.Error()
	//}
	//c.ServeJSON()
	c.Data["userName"]=c.GetSession("userName")
	c.Redirect("/getAll",302)
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	//下面这个limit是限制返回条数
	var limit int64 = 100
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
		c.Data["userName"]=c.GetSession("userName")
		if c.Data["userName"]==nil{
			c.Redirect("/login",302)
		}
	}
	//c.ServeJSON()
	c.TplName = "table.html"
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	//c.TplName="form_update.html"
	c.Redirect("/form_update",302)
	//idStr := c.Ctx.Input.Param(":id")
	idStr:=c.Input().Get("id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	fmt.Println(id)
	v := models.User{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUserById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]

func (c *UserController) Delete() {

	 idStr := c.Ctx.Input.Param(":id")

     //idStr:=c.Input().Get("id")

	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUser(id); err == nil {
		//c.Data["json"] = "OK"
		c.Redirect("/getAll",302)
	} else {
		//c.Data["json"] = err.Error()
		//c.Redirect("/",302)
		c.Ctx.WriteString("删除失败！")
		c.Ctx.WriteString(idStr)

	}
	//c.ServeJSON()
	//c.TplName="table.html"
}

