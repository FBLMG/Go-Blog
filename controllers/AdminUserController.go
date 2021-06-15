package controllers

import (
	"personForum/models"
	"github.com/astaxie/beego"
	"time"
)

/**
后台-用户模块
*/
type AdminUserController struct {
	AdminBaseController
}

/**
后台用户-获取列表【页面】
 */
func (c *AdminUserController) AdminUserList() {
	//设置当前页面标题
	c.Data["WebTitle"] = "账号管理"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(8)
	//每页默认数量
	PageLimit := 10
	//获取管理员账号列表
	page, err := c.GetInt("page") //获取页码
	if err != nil {
		page = 1
	}
	limit, err := c.GetInt("limit") //获取每页条数
	if err != nil {
		limit = PageLimit
	}
	result, count := models.PageAdminUserWithAdmin(page, limit)
	//重组友情链接数组
	adminList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//重新赋值数组
		row["Id"] = v.Id
		row["Username"] = v.Username
		row["Status"] = int64(v.Status)
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		if v.UpdateAt != 0 {
			row["UpdateAt"] = beego.Date(time.Unix(v.UpdateAt, 0), "Y-m-d H:s:i")
		} else {
			row["UpdateAt"] = ""
		}
		if v.TokenOverAt != 0 {
			row["TokenOverAt"] = beego.Date(time.Unix(v.TokenOverAt, 0), "Y-m-d H:s:i")
		} else {
			row["TokenOverAt"] = ""
		}
		adminList[k] = row
	}
	//组合返回数据
	c.Data["Admin"] = adminList     //返回的列表数组
	c.Data["PageCount"] = count     //获取的数据总条数
	c.Data["PagePage"] = page       //获取的页码
	c.Data["PageLimit"] = PageLimit //每页数量
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/user/index.tpl"
}

/**
后台用户-添加【页面】
 */
func (c *AdminUserController) AdminUserAdd() {
	//设置当前页面标题
	c.Data["WebTitle"] = "添加账户"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(8)
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	c.Data["StatusList"] = StatusList
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/user/addAdmin.tpl"
}

/**
后台用户-添加【方法】
 */
func (c *AdminUserController) AdminUserInsert() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		nickname := c.GetString("nickname")     //昵称
		username := c.GetString("username")     //用户名
		password := c.GetString("password")     //用户密码
		status, errStatus := c.GetInt("status") //状态
		//判断用户参数
		if nickname == "" {
			c.ajaxMsg("昵称获取失败", -1)
		}
		if username == "" {
			c.ajaxMsg("账号获取失败", -1)
		}
		if password == "" {
			c.ajaxMsg("密码获取失败", -1)
		}
		if errStatus != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//判断账户是否存在
		exitData := models.GetAdminUserUsername(username)
		if exitData.Id > 0 {
			c.ajaxMsg("账户已存在，请输入其他", -1)
		}
		//获取加密后的密码
		md5Password := c.encryption(password, username)
		//添加账户
		adminUser := new(models.AdminUser)
		adminUser.Nickname = nickname
		adminUser.Username = username
		adminUser.Password = md5Password
		adminUser.Status = status
		adminUser.CreateAt = time.Now().Unix()
		returnId := models.InsertAdminUser(adminUser)
		//判断是否成功
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
后台用户-编辑【页面】
 */
func (c *AdminUserController) AdminUserEdit() {
	//设置当前页面标题
	c.Data["WebTitle"] = "编辑账户"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(8)
	//获取用户参数
	id, idErr := c.GetInt("id")
	//判断用户参数
	if idErr != nil {
		c.Data["Error"] = "用户ID不能为空"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取用户
	adminUser := models.GetAdminUser(id)
	if adminUser.Id <= 0 {
		c.Data["Error"] = "用户获取失败"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	c.Data["adminUser"] = adminUser
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	c.Data["StatusList"] = StatusList
	//返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/user/editAdmin.tpl"
}

/**

后台用户-编辑【方法】
 */
func (c *AdminUserController) AdminUserUpdate() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //Id
		nickname := c.GetString("nickname")     //昵称
		username := c.GetString("username")     //用户名
		password := c.GetString("password")     //用户密码
		status, errStatus := c.GetInt("status") //状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if nickname == "" {
			c.ajaxMsg("昵称获取失败", -1)
		}
		if username == "" {
			c.ajaxMsg("账号获取失败", -1)
		}
		if password == "" {
			c.ajaxMsg("密码获取失败", -1)
		}
		if errStatus != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//获取用户信息
		exitData := models.GetAdminUser(id)
		if exitData.Id <= 0 {
			c.ajaxMsg("用户获取失败", -1)
		}
		//判断是否存在相同的用户名
		exitUsername := models.GetAdminUserUsername(username)
		if (exitUsername.Id > 0) && (exitUsername.Id != id) {
			c.ajaxMsg("存在相同的用户名", -1)
		}
		//判断密码是否相同
		updatePassword := ""
		if (password == exitData.Password) && (username == exitData.Username) {
			updatePassword = exitData.Password
		} else {
			updatePassword = c.encryption(password, username)
		}
		//修改用户
		adminUser := new(models.AdminUser)
		adminUser.Id = id
		adminUser.Nickname = nickname
		adminUser.Username = username
		adminUser.Password = updatePassword
		adminUser.Status = status
		adminUser.UpdateAt = time.Now().Unix()
		adminUser.CreateAt = exitData.CreateAt
		returnId := models.UpdateAdminUser(adminUser)
		if returnId <= 0 {
			c.ajaxMsg("账户编辑失败", -1)
		}
		//返回
		c.ajaxMsg("编辑成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
后台用户-修改用户状态【方法】
 */
func (c *AdminUserController) AdminUserUpdateStatus() {
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
		//获取用户
		exitUser := models.GetAdminUser(id)
		if exitUser.Id <= 0 {
			c.ajaxMsg("用户获取失败", -1)
		}
		//更新用户状态
		returnId := models.UpdateAdminUserStatus(id, status)
		if returnId <= 0 {
			c.ajaxMsg("用户编辑失败", -1)
		}
		//返回
		c.ajaxMsg("用户编辑成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
后台用户-删除用户【方法】
 */
func (c *AdminUserController) AdminUserDelete() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id") //Id
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("id获取失败", -1)
		}
		//获取用户
		exitUser := models.GetAdminUser(id)
		if exitUser.Id <= 0 {
			c.ajaxMsg("用户获取失败", -1)
		}
		//判断该用户是否有关联文章
		exitArticle := models.GetAdminArticleCountWithAdmin(id)
		if exitArticle > 0 {
			c.ajaxMsg("请删除该管理员发表的文章再删除", -1)
		}
		//删除用户
		returnId := models.DeleteAdminUser(id)
		if returnId <= 0 {
			c.ajaxMsg("删除失败", -1)
		}
		//返回
		c.ajaxMsg("删除成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}
