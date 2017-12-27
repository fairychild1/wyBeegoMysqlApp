package main

//import (
//	_ "wyBeegoMysqlApp/routers"
//	"github.com/astaxie/beego"
//)
//
//func main() {
//	beego.Run()
//}

import (
	"github.com/astaxie/beego"
	"wyBeegoMysqlApp/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	beego.Run()
}

