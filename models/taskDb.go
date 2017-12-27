package models
import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/astaxie/beego"
	//"time"
)

func init() {
	// set default database
	host :=beego.AppConfig.String("mysqlHost")
	port :=beego.AppConfig.String("mysqlPort")
	username :=beego.AppConfig.String("mysqlUser")
	password :=beego.AppConfig.String("mysqlPasswd")
	db := beego.AppConfig.String("mysqlDb")
	conn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn, 30)
	//orm.RegisterDataBase("default", "mysql", "wy:123456@/wy_test_db?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(TaskDb))

	// create table
	orm.RunSyncdb("default", false, true)
}


type TaskDb struct {
	Id   int64  // Unique identifier
	Title string `orm:"size(50)"` // Description
	Done  bool   // Is this task done?
}

func ListTasksInDb() (*([]*TaskDb),error){
	var tasks []*TaskDb
	o := orm.NewOrm()
	_, err := o.QueryTable("task_db").All(&tasks)
	return &tasks,err
}

func AddTask(title string) error {
	o := orm.NewOrm()
	task := TaskDb{
		Title : title,
		Done : false,
	}
	_,err := o.Insert(&task)
	if err != nil {
		fmt.Println("添加task的信息到数据库失败")
		return err
	}
	return nil
}

func UpdateTask(task *TaskDb) error {
	o := orm.NewOrm()
	_, err := o.Update(task)
	if err != nil {
		fmt.Println("在数据库更新task的信息失败")
		return err
	}
	return nil
}

func GetTask(id int64) error {
	o :=orm.NewOrm()
	var tasks []orm.Params
	_, err := o.QueryTable("task_db").Filter("id", id).Values(&tasks)
	if err != nil {
		fmt.Printf("数据库里找不到该id %d\n",id)
	}
	return err
}



