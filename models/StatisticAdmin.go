package models

/**
管理员登陆-统计记录表
 */

/**
引入模块
 */
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

//初始化表对象
var statisticAdmin StatisticAdmin

/**
后端获取列表【带分页】
 */
func PageStatisticAdminWithAdmin(page int, pageSize int, searchStartDate string, searchEndDate string, searchAdminId int) ([]StatisticAdmin, int64) {
	var list []StatisticAdmin
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(statisticAdmin)
	//判断筛选条件
	if searchStartDate != "" {
		qs = qs.Filter("create_date__gte", searchStartDate)
	}
	if searchEndDate != "" {
		qs = qs.Filter("create_date__lte", searchEndDate)
	}
	if searchAdminId > 0 {
		qs = qs.Filter("admin_id", searchAdminId)
	}
	//查询
	total, _ := qs.Count()
	qs.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
根据条件获取用户登录记录【导出专用】
 */
func PageStatisticAdminWithExcel(searchStartDate string, searchEndDate string, searchAdminId int) ([]StatisticAdmin) {
	var list []StatisticAdmin
	o := orm.NewOrm()
	qs := o.QueryTable(statisticAdmin)
	//判断筛选条件
	if searchStartDate != "" {
		qs = qs.Filter("create_date__gte", searchStartDate)
	}
	if searchEndDate != "" {
		qs = qs.Filter("create_date__lte", searchEndDate)
	}
	if searchAdminId > 0 {
		qs = qs.Filter("admin_id", searchAdminId)
	}
	//查询
	qs.OrderBy("-id").All(&list)
	return list
}

/**
获取总数
 */
func GetStatisticAdminCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(statisticAdmin)
	total, _ := qs.Count()
	return total
}

/**
添加数据
 */
func InsertStatisticAdmin(statisticAdmin *StatisticAdmin) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(statisticAdmin)
	return id
}

/**
多条件删除
 */
func DeleteStatisticAdminWithWhere(statisticAdmin *StatisticAdmin) int64 {
	o := orm.NewOrm()
	returnId, _ := o.Delete(statisticAdmin)
	return returnId
}

/**
根据条件删除数据
 */
func DeleteStatisticAdmin(searchStartDate string, searchEndDate string, searchAdminId int) string {
	deleteSql := "DELETE FROM `statistic_admin` WHERE `id`!=0 "
	//判断开始日期
	if searchStartDate != "" {
		deleteSql = deleteSql + " AND `create_date`>='" + searchStartDate + "'"
	}
	if searchEndDate != "" {
		deleteSql = deleteSql + " AND `create_date`<='" + searchEndDate + "'"
	}
	//判断文章
	if searchAdminId > 0 {
		deleteSql = deleteSql + " AND `admin_id`='" + strconv.Itoa(searchAdminId) + "'"
	}
	//删除数据
	o := orm.NewOrm()
	_, _ = o.Raw(deleteSql).Exec()
	return deleteSql
}
