package controllers

import (
	"personForum/models"
	"time"
)

/**
日志函数
 */
type LogController struct {
}

/**
管理员登陆
 */
func (c *LogController) LogAdminLogin(adminId int, ip string) int {
	//实例化数据模型
	statistic := new(models.StatisticAdmin)
	//获取当前日期
	t := time.Now()
	createDate := t.Format("2006-01-02")
	//组合数据
	statistic.AdminId = adminId
	statistic.Ip = ip
	statistic.Address = ""
	statistic.CreateDate = createDate
	statistic.CreateAt = time.Now().Unix()
	//添加数据
	models.InsertStatisticAdmin(statistic)
	//返回
	return 1
}

/**
用户文章访问
 */
func (c *LogController) LogArticleInfo(articleId int, ip string) int {
	//实例化数据模型
	statistic := new(models.StatisticArticle)
	//获取当前日期
	t := time.Now()
	createDate := t.Format("2006-01-02")
	//组合数据
	statistic.ArticleId = articleId
	statistic.Ip = ip
	statistic.Address = ""
	statistic.CreateDate = createDate
	statistic.CreateAt = time.Now().Unix()
	//添加数据
	models.InsertStatisticArticle(statistic)
	//返回
	return 1
}
