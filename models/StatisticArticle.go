package models

/**
用户访问文章-统计记录表
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
var statisticArticle StatisticArticle

/**
后端获取列表【带分页】
 */
func PageStatisticArticleWithAdmin(page int, pageSize int, searchStartDate string, searchEndDate string, searchArticleId int) ([]StatisticArticle, int64) {
	var list []StatisticArticle
	offset := (page - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(statisticArticle)
	//条件筛选
	//判断筛选条件
	if searchStartDate != "" {
		qs = qs.Filter("create_date__gte", searchStartDate)
	}
	if searchEndDate != "" {
		qs = qs.Filter("create_date__lte", searchEndDate)
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
根据条件获取数据【导出专用】
 */
func PageStatisticArticleList(searchStartDate string, searchEndDate string, searchArticleId int) ([]StatisticArticle, string) {
	var list []StatisticArticle
	o := orm.NewOrm()
	selectSql := "SELECT *  FROM `statistic_article` WHERE `id`!=0 "
	//判断开始日期
	if searchStartDate != "" {
		selectSql = selectSql + " AND `create_date`>='" + searchStartDate + "'"
	}
	if searchEndDate != "" {
		selectSql = selectSql + " AND `create_date`<='" + searchEndDate + "'"
	}
	//判断文章
	if searchArticleId > 0 {
		selectSql = selectSql + " AND `article_id`='" + strconv.Itoa(searchArticleId) + "'"
	}
	num, _ := o.Raw(selectSql).QueryRows(&list)
	if num == 0 {

	}
	return list, selectSql
}

/**
获取总数
 */
func GetStatisticArticleCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(statisticArticle)
	total, _ := qs.Count()
	return total
}

/**
添加数据
 */
func InsertStatisticArticle(statisticArticle *StatisticArticle) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(statisticArticle)
	return id
}

/**
多条件删除
 */
func DeleteStatisticArticleWithWhere(statisticArticle *StatisticArticle) int64 {
	o := orm.NewOrm()
	returnId, _ := o.Delete(statisticArticle)
	return returnId
}

/**
后端根据筛选删除数据
 */
func DeleteStatisticArticle(startDate string, endDate string, articleId int) string {
	deleteSql := "DELETE FROM `statistic_article` WHERE `id`!=0 "
	//判断开始日期
	if startDate != "" {
		deleteSql = deleteSql + " AND `create_date`>='" + startDate + "'"
	}
	if endDate != "" {
		deleteSql = deleteSql + " AND `create_date`<='" + endDate + "'"
	}
	//判断文章
	if articleId > 0 {
		deleteSql = deleteSql + " AND `article_id`='" + strconv.Itoa(articleId) + "'"
	}
	//删除数据
	o := orm.NewOrm()
	_, _ = o.Raw(deleteSql).Exec()
	return deleteSql
}
