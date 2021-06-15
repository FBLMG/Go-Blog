package models

/**
文章-分类
 */

/**
引入模块
 */
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//初始化表对象
var classify Classify

/**
获取分类列表【前端】
 */
func GetClassifyListWithHome() []*Classify {
	var classifyList []*Classify
	o := orm.NewOrm()
	o.QueryTable(classify).Filter("status", 1).All(&classifyList)
	return classifyList
}

/**
根据ID获取分类详情
 */
func GetClassify(id int) Classify {
	var classifyInfo Classify
	o := orm.NewOrm()
	o.QueryTable(classify).Filter("id", id).OrderBy("sort", "id").One(&classifyInfo)
	return classifyInfo
}

/**
后端获取分类列表
 */
func GetClassifyListWithAdmin() []*Classify {
	var classifyList []*Classify
	o := orm.NewOrm()
	o.QueryTable(classify).All(&classifyList)
	return classifyList
}

/**
后端获取分类列表【带分页】
 */
func PageClassifyWithAdmin(page, pageSize int) ([]Classify, int64) {
	var list []Classify
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(classify)
	total, _ := qs.Count()
	qs.OrderBy("sort", "id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
获取文章分类总数
 */
func GetClassifyCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(classify)
	total, _ := qs.Count()
	return total
}

/**
根据标题获取分类详情
 */
func GetClassifyInTitle(title string) Classify {
	var classifyInfo Classify
	o := orm.NewOrm()
	o.QueryTable(classify).Filter("title", title).One(&classifyInfo)
	return classifyInfo
}

/**
添加分类
 */
func InsertClassify(classify *Classify) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(classify)
	return id
}

/**
编辑分类
 */
func UpdateClassify(classify *Classify) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(classify)
	return id
}

/**
根据ID删除分类
 */
func DeleteClassify(id int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(classify).Filter("id", id).Delete()
	return returnId
}

/**
多条件删除分类
 */
func DeleteClassifyWithWhere(classify *Classify) int64 {
	o := orm.NewOrm()
	returnId, _ := o.Delete(classify)
	return returnId
}

/**
修改状态【后端】
 */
func UpdateClassifyStatus(id, status int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(classify).Filter("id", id).Update(orm.Params{
		"status":    status,
		"update_at": time.Now().Unix(),
	})
	return returnId
}
