package models

/**
文章
 */

/**
引入模块
 */
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//初始化
var article Article

/**
前端获取文章列表(获取全部文章)
 */
func PageArticleWithHome(page, pageSize int) ([]Article, int64) {
	var list []Article
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	total, _ := qs.Count()
	qs.OrderBy("-id").Filter("status", 1).Limit(pageSize, offset).All(&list)
	return list, total
}

/**
前端获取置顶文章
 */
func GetArticleWithSign() []*Article {
	var articleList []*Article
	o := orm.NewOrm()
	o.QueryTable(article).Filter("sign", 1).Filter("status", 1).OrderBy("-sort", "-id").All(&articleList)
	return articleList
}

/**
前端获取文章列表(根据分类)
 */
func PageArticleWithHomeClassify(classifyId, page, pageSize int) ([]Article, int64) {
	var list []Article
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	qs = qs.Filter("type_id", classifyId).Filter("status", 1)
	total, _ := qs.Count()
	qs.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
前端获取文章列表(根据标签)
 */
func PageArticleWithHomeLabel(labelId, page, pageSize int) ([]Article, int64) {
	var list []Article
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	qs = qs.Filter("label_id", labelId).Filter("status", 1)
	total, _ := qs.Count()
	qs.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
根据文章分类获取文章总数【前端】
 */
func GetClassifyArticleCount(classifyId int) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	qs = qs.Filter("type_id", classifyId).Filter("status", 1)
	total, _ := qs.Count()
	return total
}

/**
根据文章标签获取文章总数【前端】
 */
func GetLabelArticleCount(labelId int) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	qs = qs.Filter("label_id", labelId).Filter("status", 1)
	total, _ := qs.Count()
	return total
}

/**
根据ID获取博客详情
 */
func GetArticle(id int) Article {
	var articleInfo Article
	o := orm.NewOrm()
	o.QueryTable(article).Filter("id", id).One(&articleInfo)
	return articleInfo
}

/**
文章获取列表
 */
func GetArticleWithAdmin() []*Article {
	var articleList []*Article
	o := orm.NewOrm()
	o.QueryTable(article).All(&articleList)
	return articleList
}

/**
前端获取相同分类下其他文章
 */
func GetArticleIdenticalWithHome(articleTypeId, id int) []*Article {
	var articleList []*Article
	o := orm.NewOrm()
	o.QueryTable(article).Filter("type_id", articleTypeId).Limit(2).All(&articleList)
	return articleList
}

/**
前端获取相同分类下其他文章【原生查询】
 */
func GetArticleIdenticalWithHomeNew(articleTypeId, id int) (dataList []Article) {
	var list []Article
	o := orm.NewOrm()
	num, _ := o.Raw("SELECT * FROM `article` WHERE `type_id` =? and `id` !=? and `status`=? limit ?", articleTypeId, id, 1, 2).QueryRows(&list)
	if num == 0 {

	}
	return list
}

/**
后端获取文章列表(获取全部文章)
 */
func PageArticleWithAdmin(page int, pageSize int, title string, typeId int, labelId int, searchSign int) ([]Article, int64) {
	var list []Article
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	//判断筛选
	if title != "" {
		qs = qs.Filter("title", title)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if labelId > 0 {
		qs = qs.Filter("label_id", labelId)
	}
	if searchSign > 0 {
		qs = qs.Filter("sign", searchSign)
	}
	//获取数据
	total, _ := qs.Count()
	qs.OrderBy("-sort", "-id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
获取文章总数
 */
func GetArticleCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	total, _ := qs.Count()
	return total
}

/**
根据文章分类ID获取文章总数【后端】
 */
func GetClassifyArticleCountWithAdmin(classifyId int) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	qs = qs.Filter("type_id", classifyId)
	total, _ := qs.Count()
	return total
}

/**
根据文章标签ID获取文章总数【后端】
 */
func GetLabelArticleCountWithAdmin(labelId int) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	qs = qs.Filter("label_id", labelId)
	total, _ := qs.Count()
	return total
}

/**
根据文章管理员获取文章总数【后端】
 */
func GetAdminArticleCountWithAdmin(adminId int) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(article)
	qs = qs.Filter("user_id", adminId)
	total, _ := qs.Count()
	return total
}

/**
添加博客【后端】
 */
func InsertArticle(article *Article) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(article)
	return id
}

/**
编辑博客
 */
func UpdateArticle(article *Article) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(article)
	return id
}

/**
根据ID删除博客
 */
func DeleteArticle(id int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(article).Filter("id", id).Delete()
	return returnId
}

/**
多条件删除博客
 */
func DeleteArticleWithWhere(article *Article) int64 {
	o := orm.NewOrm()
	returnId, _ := o.Delete(article)
	return returnId
}

/**
修改状态
 */
func UpdateArticleStatus(id, status int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(article).Filter("id", id).Update(orm.Params{
		"status":    status,
		"update_at": time.Now().Unix(),
	})
	return returnId
}

/**
修改置顶状态
 */
func UpdateArticleSign(id, sign int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(article).Filter("id", id).Update(orm.Params{
		"sign":      sign,
		"update_at": time.Now().Unix(),
	})
	return returnId
}
