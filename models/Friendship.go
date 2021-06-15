package models

/**
友情链接
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
var friendship Friendship

/**
获取列表【前端】
 */
func GetFriendshipListWithHome() []*Friendship {
	var friendshipList []*Friendship
	o := orm.NewOrm()
	o.QueryTable(friendship).Filter("status", 1).OrderBy("sort", "id").All(&friendshipList)
	return friendshipList
}

/**
根据ID获取详情
 */
func GetFriendship(id int) Friendship {
	var friendshipInfo Friendship
	o := orm.NewOrm()
	o.QueryTable(friendship).Filter("id", id).One(&friendshipInfo)
	return friendshipInfo
}

/**
后端获取列表
 */
func GetFriendshipListWithAdmin() []*Friendship {
	var friendshipList []*Friendship
	o := orm.NewOrm()
	o.QueryTable(friendship).All(&friendshipList)
	return friendshipList
}

/**
后端获取列表【带分页】
 */
func PageFriendshipWithAdmin(page, pageSize int) ([]Friendship, int64) {
	var list []Friendship
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(friendship)
	total, _ := qs.Count()
	qs.OrderBy("sort", "id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
获取总数
 */
func GetFriendshipCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(friendship)
	total, _ := qs.Count()
	return total
}

/**
根据标题获取详情
 */
func GetFriendshipInTitle(title string) Friendship {
	var friendshipInfo Friendship
	o := orm.NewOrm()
	o.QueryTable(friendship).Filter("title", title).One(&friendshipInfo)
	return friendshipInfo
}

/**
添加数据
 */
func InsertFriendship(friendship *Friendship) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(friendship)
	return id
}

/**
编辑数据
 */
func UpdateFriendship(friendship *Friendship) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(friendship)
	return id
}

/**
根据ID删除
 */
func DeleteFriendship(id int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(friendship).Filter("id", id).Delete()
	return returnId
}

/**
多条件删除
 */
func DeleteFriendshipWithWhere(friendship *Friendship) int64 {
	o := orm.NewOrm()
	returnId, _ := o.Delete(friendship)
	return returnId
}

/**
修改状态【后端】
 */
func UpdateFriendshipStatus(id, status int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(friendship).Filter("id", id).Update(orm.Params{
		"status":    status,
		"update_at": time.Now().Unix(),
	})
	return returnId
}
