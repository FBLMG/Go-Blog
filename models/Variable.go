package models

/**
变量表
 */

/**
引入模块
 */
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//初始化表
var variable Variable

/**
获取配置变量
 */
func GetVariable(name string) Variable {
	var variableInfo Variable
	o := orm.NewOrm()
	o.QueryTable(variable).Filter("name", name).One(&variableInfo)
	return variableInfo
}

/**
设置配置
 */
func SetVariable(name, desc, value string) int {
	//先获取是存在名称的配置信息
	var variableInfo Variable
	var returnId int
	o := orm.NewOrm()
	o.QueryTable(variable).Filter("name", name).One(&variableInfo)
	if variableInfo.Id > 0 {
		//已经存在该配置则直接修改改配置
		variableDate := Variable{Id: variableInfo.Id, Name: name, Desc: desc, Value: value, CreateAt: variableInfo.CreateAt, UpdateAt: time.Now().Unix()}
		updateId, _ := o.Update(&variableDate)
		returnId = int(updateId)
	} else {
		//改配置不存在则直接添加该配置
		variableDate := Variable{Name: name, Desc: desc, Value: value, CreateAt: time.Now().Unix()}
		insertId, _ := o.Insert(&variableDate)
		returnId = int(insertId)
	}
	return returnId
}
