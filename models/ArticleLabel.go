package models

/**
文章-标签
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
var articleLabel ArticleLabel

/**
获取标签列表【前端】
 */
func GetArticleLabelListWithHome() []*ArticleLabel {
	o := orm.NewOrm()
	var articleLabelList []*ArticleLabel
	o.QueryTable(articleLabel).Filter("status", 1).OrderBy("sort", "id").All(&articleLabelList)
	return articleLabelList
}

/**
根据ID获取标签详情
 */
func GetArticleLabel(id int) ArticleLabel {
	var articleLabelInfo ArticleLabel
	o := orm.NewOrm()
	o.QueryTable(articleLabel).Filter("id", id).One(&articleLabelInfo)
	return articleLabelInfo
}

/**
后端获取标签列表
 */
func GetArticleLabelListWithAdmin() []*ArticleLabel {
	var articleLabelList []*ArticleLabel
	o := orm.NewOrm()
	o.QueryTable(articleLabel).All(&articleLabelList)
	return articleLabelList
}

/**
后端获取标签列表【带分页】
 */
func PageArticleLabelWithAdmin(page, pageSize int) ([]ArticleLabel, int64) {
	var list []ArticleLabel
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(articleLabel)
	total, _ := qs.Count()
	qs.OrderBy("sort", "id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
获取文章标签总数
 */
func GetArticleLabelCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(articleLabel)
	total, _ := qs.Count()
	return total
}

/**
根据标题获取标签详情
 */
func GetArticleLabelInTitle(title string) ArticleLabel {
	var articleLabelInfo ArticleLabel
	o := orm.NewOrm()
	o.QueryTable(articleLabel).Filter("title", title).One(&articleLabelInfo)
	return articleLabelInfo
}

/**
添加标签
 */
func InsertArticleLabel(articleLabel *ArticleLabel) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(articleLabel)
	return id
}

/**
编辑标签
 */
func UpdateArticleLabel(articleLabel *ArticleLabel) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(articleLabel)
	return id
}

/**
根据ID删除标签
 */
func DeleteArticleLabel(id int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(articleLabel).Filter("id", id).Delete()
	return returnId
}

/**
多条件删除分类
 */
func DeleteArticleLabelWithWhere(articleLabel *ArticleLabel) int64 {
	o := orm.NewOrm()
	returnId, _ := o.Delete(articleLabel)
	return returnId
}

/**
修改状态【后端】
 */
func UpdateArticleLabelStatus(id, status int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(articleLabel).Filter("id", id).Update(orm.Params{
		"status":    status,
		"update_at": time.Now().Unix(),
	})
	return returnId
}
