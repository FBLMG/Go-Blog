package models

/**
文章评论
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
var articleComment ArticleComment

/**
前端获取文章评论列表
 */
func PageArticleCommentWithHome(page, pageSize, articleId int) ([]ArticleComment, int64) {
	var list []ArticleComment
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(articleComment).Filter("status", 1).Filter("article_id", articleId)
	total, _ := qs.Count()
	qs.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
获取待审核评论
 */
func GetArticleCommentWithStatus() []*ArticleComment {
	var articleCommentList []*ArticleComment
	o := orm.NewOrm()
	o.QueryTable(articleComment).Filter("status", 2).Limit(4).All(&articleCommentList)
	return articleCommentList
}

/**
根据ID获取博客评论
 */
func GetArticleComment(id int) ArticleComment {
	var articleCommentInfo ArticleComment
	o := orm.NewOrm()
	o.QueryTable(articleComment).Filter("id", id).One(&articleCommentInfo)
	return articleCommentInfo
}

/**
后端获取评论列表
 */
func PageArticleCommentWithAdmin(page, pageSize, status, adminStatus, id, searchArticleId int) ([]ArticleComment, int64) {
	var list []ArticleComment
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(articleComment)
	//判断筛选
	if status > 0 {
		qs = qs.Filter("status", status)
	}
	if adminStatus > 0 {
		qs = qs.Filter("admin_status", adminStatus)
	}
	if id > 0 {
		qs = qs.Filter("id", id)
	}
	if searchArticleId > 0 {
		qs = qs.Filter("article_id", searchArticleId)
	}
	//获取数据
	total, _ := qs.Count()
	qs.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

/**
添加评论【后端】
 */
func InsertArticleComment(articleComment *ArticleComment) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(articleComment)
	return id
}

/**
修改状态
 */
func UpdateArticleCommentStatus(id, status int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(articleComment).Filter("id", id).Update(orm.Params{
		"status":    status,
		"update_at": time.Now().Unix(),
	})
	return returnId
}

/**
修改置顶状态
 */
func UpdateArticleCommentSign(id, sign int) int64 {
	o := orm.NewOrm()
	returnId, _ := o.QueryTable(articleComment).Filter("id", id).Update(orm.Params{
		"sign":      sign,
		"update_at": time.Now().Unix(),
	})
	return returnId
}
