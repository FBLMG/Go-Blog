package models

/**
管理员
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
var adminUser AdminUser

/**
根据ID获取详情
 */
func GetAdminUser(id int) AdminUser {
	var adminInfo AdminUser
	o := orm.NewOrm()
	o.QueryTable(adminUser).Filter("id", id).One(&adminInfo)
	return adminInfo
}

/**
后端获取列表
 */
func GetAdminUserWithAdmin() []*AdminUser {
	var adminUserList []*AdminUser
	o := orm.NewOrm()
	o.QueryTable(adminUser).All(&adminUserList)
	return adminUserList
}

/**
后端获取列表【带分页】
 */
func PageAdminUserWithAdmin(page, pageSize int) ([]AdminUser, int64) {
	var list []AdminUser
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(adminUser)
	total, _ := qs.Count()
	qs.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
获取总数
 */
func GetAdminUserCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(adminUser)
	total, _ := qs.Count()
	return total
}

/**
根据账号获取详情
 */
func GetAdminUserUsername(username string) AdminUser {
	var adminInfo AdminUser
	o := orm.NewOrm()
	o.QueryTable(adminUser).Filter("username", username).One(&adminInfo)
	return adminInfo
}

/**
添加数据
 */
func InsertAdminUser(adminUser *AdminUser) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(adminUser)
	return id
}

/**
编辑数据
 */
func UpdateAdminUser(adminUser *AdminUser) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(adminUser)
	return id
}

/**
根据ID删除
 */
func DeleteAdminUser(id int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(adminUser).Filter("id", id).Delete()
	return returnId
}

/**
多条件删除
 */
func DeleteAdminUserWithWhere(adminUser *AdminUser) int64 {
	o := orm.NewOrm()
	returnId, _ := o.Delete(adminUser)
	return returnId
}

/**
修改状态【后端】
 */
func UpdateAdminUserStatus(id, status int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(adminUser).Filter("id", id).Update(orm.Params{
		"status":    status,
		"update_at": time.Now().Unix(),
	})
	return returnId
}

/**
修改用户登陆状态【后端】
 */
func UpdateAdminUserLoginInfo(id int, tokenOverAt int64, token string) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(adminUser).Filter("id", id).Update(orm.Params{
		"token_over_at": tokenOverAt,
		"token":         token,
		"update_at":     time.Now().Unix(),
	})
	return returnId
}
