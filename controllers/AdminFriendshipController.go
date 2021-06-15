package controllers

import (
	"personForum/models"
	"github.com/astaxie/beego"
	"time"
)

/**
后台-友情链接模块
 */
type AdminFriendshipController struct {
	AdminBaseController
}

/**
友情链接
 */

/**
友情链接-获取列表
 */
func (c *AdminFriendshipController) AdminFriendshipList() {
	//设置当前页面标题
	c.Data["WebTitle"] = "友情链接"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(5)
	//每页默认数量
	PageLimit := 10
	//获取友情链接列表
	page, err := c.GetInt("page") //获取页码
	if err != nil {
		page = 1
	}
	limit, err := c.GetInt("limit") //获取每页条数
	if err != nil {
		limit = PageLimit
	}
	result, count := models.PageFriendshipWithAdmin(page, limit)
	//重组友情链接数组
	friendsList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["Url"] = v.Url
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		if v.UpdateAt != 0 {
			row["UpdateAt"] = beego.Date(time.Unix(v.UpdateAt, 0), "Y-m-d H:s:i")
		} else {
			row["UpdateAt"] = ""
		}
		row["Status"] = int64(v.Status)
		friendsList[k] = row
	}
	c.Data["Friends"] = friendsList //返回的列表数组
	c.Data["PageCount"] = count     //获取的数据总条数
	c.Data["PagePage"] = page       //获取的页码
	c.Data["PageLimit"] = PageLimit //每页数量
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/friend/index.tpl"
}

/**
友情链接-添加【页面】
 */
func (c *AdminFriendshipController) AdminFriendshipAdd() {
	//设置当前页面标题
	c.Data["WebTitle"] = "添加友情链接"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(5)
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	c.Data["StatusList"] = StatusList
	//返回页面
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/friend/addFriend.tpl"
}

/**
友情链接-添加【方法】
 */
func (c *AdminFriendshipController) AdminFriendshipInsert() {
	//判断用户请求
	if c.isPost() {
		//获取用户请求
		title := c.GetString("title")           //备注
		url := c.GetString("url")               //链接
		describe := c.GetString("describe")     //描述
		sort, _ := c.GetInt("sort")             //序号
		status, statusErr := c.GetInt("status") //状态
		//判断用户参数
		if title == "" {
			c.ajaxMsg("标题不能为空", -1)
		}
		if url == "" {
			c.ajaxMsg("链接不能为空", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态不能为空", -1)
		}
		//添加数据
		friendship := new(models.Friendship)
		friendship.Title = title
		friendship.Url = url
		friendship.Describe = describe
		friendship.Sort = sort
		friendship.Status = status
		friendship.CreateAt = time.Now().Unix()
		returnId := models.InsertFriendship(friendship)
		//判断是否添加成功
		if returnId <= 0 {
			c.ajaxMsg("添加失败", -1)
		}
		//返回
		c.ajaxMsg("添加成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
友情链接-编辑【页面】
 */
func (c *AdminFriendshipController) AdminFriendshipEdit() {
	//设置当前页面标题
	c.Data["WebTitle"] = "编辑友情链接"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(5)
	//获取用户参数
	id, idErr := c.GetInt("id")
	//判断用户参数
	if idErr != nil {
		c.Data["Error"] = "ID不能为空"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取友情链接
	friendship := models.GetFriendship(id)
	if friendship.Id <= 0 {
		c.Data["Error"] = "友情链接获取失败"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	c.Data["StatusList"] = StatusList
	//组合返回数据
	c.Data["Friendship"] = friendship
	//返回数据
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/friend/editFriend.tpl"
}

/**
友情链接-编辑【方法】
 */
func (c *AdminFriendshipController) AdminFriendshipUpdate() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //id
		title := c.GetString("title")           //标题
		url := c.GetString("url")               //链接
		describe := c.GetString("describe")     //描述
		sort, _ := c.GetInt("sort")             //序号
		status, statusErr := c.GetInt("status") //状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if title == "" {
			c.ajaxMsg("标题获取失败", -1)
		}
		if url == "" {
			c.ajaxMsg("链接获取失败", -1)
		}
		if describe == "" {
			c.ajaxMsg("描述获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//获取友情链接
		exitData := models.GetFriendship(id)
		if exitData.Id <= 0 {
			c.ajaxMsg("数据获取失败", -1)
		}
		//编辑友情链接
		friendship := new(models.Friendship)
		friendship.Id = id
		friendship.Title = title
		friendship.Url = url
		friendship.Describe = describe
		friendship.Sort = sort
		friendship.Status = status
		friendship.CreateAt = exitData.CreateAt
		friendship.UpdateAt = time.Now().Unix()
		returnId := models.UpdateFriendship(friendship)
		//判断是否成功
		if returnId <= 0 {
			c.ajaxMsg("编辑失败", -1)
		}
		//返回
		c.ajaxMsg("编辑成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
友情链接-编辑状态【方法】
 */
func (c *AdminFriendshipController) AdminFriendshipUpdateStatus() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //id
		status, statusErr := c.GetInt("status") //状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//获取友情链接
		exitFriendship := models.GetFriendship(id)
		if exitFriendship.Id <= 0 {
			c.ajaxMsg("数据获取失败", -1)
		}
		//编辑数据
		returnId := models.UpdateFriendshipStatus(id, status)
		if returnId <= 0 {
			c.ajaxMsg("数据获取失败", -1)
		}
		//返回
		c.ajaxMsg("编辑成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
友情链接-删除【方法】
 */
func (c *AdminFriendshipController) AdminFriendshipDelete() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id") //id
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		//获取友情链接
		exitFriendship := models.GetFriendship(id)
		if exitFriendship.Id <= 0 {
			c.ajaxMsg("数据获取失败", -1)
		}
		//删除数据
		returnId := models.DeleteFriendship(id)
		if returnId <= 0 {
			c.ajaxMsg("删除失败", -1)
		}
		//返回
		c.ajaxMsg("删除成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}
