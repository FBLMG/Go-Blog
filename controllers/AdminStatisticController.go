package controllers

import (
	"personForum/models"
	"github.com/astaxie/beego"
	"time"
	"github.com/tealeg/xlsx"
)

/**
后台-统计模块
*/
type AdminStatisticController struct {
	AdminBaseController
}

/**
统计-用户登陆
 */
func (c *AdminStatisticController) AdminStatisticUserLogin() {
	//设置当前页面标题
	c.Data["WebTitle"] = "用户登陆统计"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(9)
	//默认每页数量
	PageLimit := 10
	//获取用户参数
	page, err := c.GetInt("page") //获取页码
	if err != nil {
		page = 1
	}
	limit, err := c.GetInt("limit") //获取每页条数
	if err != nil {
		limit = PageLimit
	}
	//获取用户搜索参数
	searchStartDate := c.GetString("searchStartDate")            //搜索-开始日期
	searchEndDate := c.GetString("searchEndDate")                //搜索-结束日期
	searchAdminId, searchAdminIdErr := c.GetInt("searchAdminId") //搜索-管理员Id
	if searchStartDate == "" {
		searchStartDate = ""
	}
	if searchEndDate == "" {
		searchEndDate = ""
	}
	if searchAdminIdErr != nil {
		searchAdminId = 0
	}
	//获取管理员
	adminUser := models.GetAdminUserWithAdmin()
	//获取统计数据
	result, count := models.PageStatisticAdminWithAdmin(page, limit, searchStartDate, searchEndDate, searchAdminId)
	//重组文章统计数组
	statisticList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//循环获取登陆用户名
		adminUserName := ""
		for _, vUser := range adminUser {
			if vUser.Id == v.AdminId {
				adminUserName = vUser.Nickname
				continue
			}
		}
		//重新赋值数组
		row["Id"] = v.Id
		row["AdminId"] = v.AdminId
		row["AdminUserName"] = adminUserName
		row["Address"] = v.Address
		row["Ip"] = v.Ip
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		statisticList[k] = row
	}
	//返回用户搜索参数
	getInput := make(map[string]interface{})
	getInput["searchStartDate"] = searchStartDate //搜索-开始日期
	getInput["searchEndDate"] = searchEndDate     //搜索-结束日期
	getInput["searchAdminId"] = searchAdminId     //搜索-管理员Id
	c.Data["getInput"] = getInput
	//其他信息
	c.Data["AdminUser"] = adminUser         //管理员列表
	c.Data["StatisticList"] = statisticList //返回的列表数组
	c.Data["PageCount"] = count             //获取的数据总条数
	c.Data["PagePage"] = page               //获取的页码
	c.Data["PageLimit"] = PageLimit         //每页数量
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/statistic/statisticAdmin.tpl"
}

/**
统计-删除用户登录数据
 */
func (c *AdminStatisticController) AdminStatisticUserLoginDelete() {
	//判断用户请求
	if c.isPost() {
		//获取参数
		searchStartDate := c.GetString("searchStartDate")            //搜索-开始日期
		searchEndDate := c.GetString("searchEndDate")                //搜索-结束日期
		searchAdminId, searchAdminIdErr := c.GetInt("searchAdminId") //搜索-管理员Id
		if searchStartDate == "" {
			searchStartDate = ""
		}
		if searchEndDate == "" {
			searchEndDate = ""
		}
		if searchAdminIdErr != nil {
			searchAdminId = 0
		}
		//删除数据
		models.DeleteStatisticAdmin(searchStartDate, searchEndDate, searchAdminId)
		//返回
		c.ajaxMsg("删除成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
统计-用户登录数据导出
 */
func (c *AdminStatisticController) AdminStatisticUserLoginExcel() {
	//获取用户搜索参数
	searchStartDate := c.GetString("searchStartDate")            //搜索-开始日期
	searchEndDate := c.GetString("searchEndDate")                //搜索-结束日期
	searchAdminId, searchAdminIdErr := c.GetInt("searchAdminId") //搜索-管理员Id
	if searchStartDate == "" {
		searchStartDate = ""
	}
	if searchEndDate == "" {
		searchEndDate = ""
	}
	if searchAdminIdErr != nil {
		searchAdminId = 0
	}
	//设置保存路径
	filePath := "static/file/AdminStatisticUserLogin.xlsx"
	//创建文件
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	//设置表格宽度从第一列到第三列,宽度设置为30
	sheet.SetColWidth(0, 2, 30.0)
	//创建表头
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "登录账户"
	cell = row.AddCell()
	cell.Value = "登录IP"
	cell = row.AddCell()
	cell.Value = "登录时间"
	//获取管理员
	adminUser := models.GetAdminUserWithAdmin()
	//获取统计数据
	result := models.PageStatisticAdminWithExcel(searchStartDate, searchEndDate, searchAdminId)
	//重组文章统计数组
	for _, v := range result {
		//循环获取登陆用户名
		adminUserName := ""
		for _, vUser := range adminUser {
			if vUser.Id == v.AdminId {
				adminUserName = vUser.Nickname
				continue
			}
		}
		//获取时间
		CreateAt := beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		//追加数据
		row = sheet.AddRow()
		sheet.SetColWidth(0, 2, 30.0)
		cell = row.AddCell()
		cell.Value = adminUserName
		cell = row.AddCell()
		cell.Value = v.Ip
		cell = row.AddCell()
		cell.Value = CreateAt
		file.Save(filePath)
	}
	// 导出表格
	c.Ctx.Output.Download(filePath, "用户登录统计.xlsx")
}

/**
统计-文章访问
 */
func (c *AdminStatisticController) AdminStatisticArticle() {
	//设置当前页面标题
	c.Data["WebTitle"] = "文章访问统计"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(9)
	//默认每页数量
	PageLimit := 10
	//获取用户参数
	page, err := c.GetInt("page") //获取页码
	if err != nil {
		page = 1
	}
	limit, err := c.GetInt("limit") //获取每页条数
	if err != nil {
		limit = PageLimit
	}
	//获取用户搜索参数
	searchStartDate := c.GetString("searchStartDate")                  //搜索-开始日期
	searchEndDate := c.GetString("searchEndDate")                      //搜索-结束日期
	searchArticleId, searchArticleIdErr := c.GetInt("searchArticleId") //搜索-文章Id
	if searchStartDate == "" {
		searchStartDate = ""
	}
	if searchEndDate == "" {
		searchEndDate = ""
	}
	if searchArticleIdErr != nil {
		searchArticleId = 0
	}
	//获取文章
	article := models.GetArticleWithAdmin()
	//获取统计数据
	result, count := models.PageStatisticArticleWithAdmin(page, limit, searchStartDate, searchEndDate, searchArticleId)
	//重组文章统计数组
	statisticList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//获取文章名称
		ArticleTitle := ""
		for _, vArticle := range article {
			if vArticle.Id == v.ArticleId {
				ArticleTitle = vArticle.Title
				continue
			}
		}
		//重新赋值数组
		row["Id"] = v.Id
		row["ArticleId"] = v.ArticleId
		row["ArticleTitle"] = ArticleTitle
		row["Ip"] = v.Ip
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		statisticList[k] = row
	}
	//返回用户搜索参数
	getInput := make(map[string]interface{})
	getInput["searchStartDate"] = searchStartDate //搜索-开始日期
	getInput["searchEndDate"] = searchEndDate     //搜索-结束日期
	getInput["searchArticleId"] = searchArticleId //搜索-文章Id
	c.Data["getInput"] = getInput
	//组合返回数据
	c.Data["Article"] = article             //文章列表
	c.Data["StatisticList"] = statisticList //返回的列表数组
	c.Data["PageCount"] = count             //获取的数据总条数
	c.Data["PagePage"] = page               //获取的页码
	c.Data["PageLimit"] = PageLimit         //每页数量
	/** 页面返回**/
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/statistic/statisticArticle.tpl"
}

/**
统计-文章数据删除
 */
func (c *AdminStatisticController) AdminStatisticArticleDelete() {
	//判断用户请求
	if c.isPost() {
		//获取参数
		searchStartDate := c.GetString("searchStartDate")                  //搜索-开始日期
		searchEndDate := c.GetString("searchEndDate")                      //搜索-结束日期
		searchArticleId, searchArticleIdErr := c.GetInt("searchArticleId") //搜索-文章Id
		if searchStartDate == "" {
			searchStartDate = ""
		}
		if searchEndDate == "" {
			searchEndDate = ""
		}
		if searchArticleIdErr != nil {
			searchArticleId = 0
		}
		//删除数据
		models.DeleteStatisticArticle(searchStartDate, searchEndDate, searchArticleId)
		//返回
		c.ajaxMsg("删除成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
统计-文章数据导出
 */
func (c *AdminStatisticController) AdminStatisticArticleExcel() {
	//获取参数
	searchStartDate := c.GetString("searchStartDate")                  //搜索-开始日期
	searchEndDate := c.GetString("searchEndDate")                      //搜索-结束日期
	searchArticleId, searchArticleIdErr := c.GetInt("searchArticleId") //搜索-文章Id
	//判断过滤参数
	if searchStartDate == "" {
		searchStartDate = ""
	}
	if searchEndDate == "" {
		searchEndDate = ""
	}
	if searchArticleIdErr != nil {
		searchArticleId = 0
	}
	//获取文章
	article := models.GetArticleWithAdmin()
	//获取统计数据
	result, _ := models.PageStatisticArticleList(searchStartDate, searchEndDate, searchArticleId)
	//设置保存路径
	filePath := "static/file/AdminStatisticArticle.xlsx"
	//创建文件
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	//设置表格宽度从第一列到第三列,宽度设置为30
	sheet.SetColWidth(0, 2, 50.0)
	//创建表头
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "访问文章"
	cell = row.AddCell()
	cell.Value = "访问IP"
	cell = row.AddCell()
	cell.Value = "访问时间"
	//循环塞进容器
	for _, v := range result {
		//获取文章名称
		ArticleTitle := ""
		for _, vArticle := range article {
			if vArticle.Id == v.ArticleId {
				ArticleTitle = vArticle.Title
				continue
			}
		}
		//格式化时间
		CreateAt := beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		//追加数据
		row = sheet.AddRow()
		sheet.SetColWidth(0, 2, 50.0)
		cell = row.AddCell()
		cell.Value = ArticleTitle
		cell = row.AddCell()
		cell.Value = v.Ip
		cell = row.AddCell()
		cell.Value = CreateAt
		file.Save(filePath)
	}
	// 导出表格
	c.Ctx.Output.Download(filePath, "文章访问统计.xlsx")
}
