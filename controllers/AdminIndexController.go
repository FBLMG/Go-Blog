package controllers

import (
	"personForum/models"
	"github.com/astaxie/beego"
	"time"
)

/**后台-首页模块**/
type AdminIndexController struct {
	AdminBaseController
}

/**
启动初始化
 */
func init() {
	//初始化Orm并注册数据库
	models.ConnectionDataBase()
}

/**
后台首页模块
 */
func (c *AdminIndexController) AdminIndex() {
	/**设置当前页面标题**/
	c.Data["WebTitle"] = "首页"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(1)
	//获取面板信息
	articleCount := models.GetArticleCount()    //文章数量
	classifyCount := models.GetClassifyCount()  //分类数量
	labelCount := models.GetArticleLabelCount() //标签数量
	friendCount := models.GetFriendshipCount()  //友链数量
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	//获取4篇待审核的评论
	newArticleFour := models.GetArticleCommentWithStatus()
	newArticleFourList := make([]map[string]interface{}, len(newArticleFour))
	for k, v := range newArticleFour {
		row := make(map[string]interface{})
		//重新赋值数组
		row["Id"] = v.Id
		row["UserNickname"] = v.UserNickname
		row["UserEmail"] = v.UserEmail
		row["Content"] = v.Content
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		newArticleFourList[k] = row
	}
	//获取置顶文章
	signArticleData, _ := models.PageArticleWithAdmin(1, 2, "", 0, 0, 1)
	signArticleDataList := make([]map[string]interface{}, len(signArticleData))
	for k, v := range signArticleData {
		row := make(map[string]interface{})
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["Image"] = imageUrl + v.Image
		row["Describe"] = v.Describe
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		signArticleDataList[k] = row
	}
	//组合返回数据
	c.Data["ArticleCount"] = articleCount         //文章数量
	c.Data["ClassifyCount"] = classifyCount       //分类数量
	c.Data["LabelCount"] = labelCount             //标签数量
	c.Data["FriendCount"] = friendCount           //友链数量
	c.Data["NewArticleFour"] = newArticleFourList //最新发表文章
	c.Data["NewArticleTwo"] = signArticleDataList //置顶文章
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/index/index.tpl"
}
