package controllers

import (
	"time"
	"fmt"
	"personForum/models"
)

/**
后台-用户登陆操作
 */
type AdminLoginController struct {
	AdminBaseController
	LogController
}

/**
登陆【页面】
 */
func (c *AdminLoginController) AdminLogin() {
	//设置当前页面标题
	c.Data["WebTitle"] = "登陆"
	//页面返回
	c.TplName = "admin/login/login.tpl"
}

/**
用户登陆【方法】
 */
func (c *AdminLoginController) AdminDoLogin() {
	//判断用户请求
	if c.isPost() {
		//获取用户请求
		username := c.GetString("username")
		password := c.GetString("password")
		//判断用户输入
		if username == "" {
			c.ajaxMsg("用户账号不能为空", -1)
		}
		if password == "" {
			c.ajaxMsg("用户密码不能为空", -1)
		}
		//判断是否存在该用户
		exitAdmin := models.GetAdminUserUsername(username)
		if exitAdmin.Id <= 0 {
			c.ajaxMsg("用户获取失败", -1)
		}
		//判断用户名跟密码是否一致
		passwordToUser := c.encryption(password, username)
		if passwordToUser != exitAdmin.Password {
			c.ajaxMsg("用户密码错误", -1)
		}
		//判断账号是否被禁用
		if exitAdmin.Status == 2 {
			c.ajaxMsg("账号被禁用,无法登陆", -1)
		}
		//设置用户登陆秘钥
		tokenId := exitAdmin.Id
		token := generatingToken()
		tokenLastTime := time.Now().Unix() + 4200
		//更需要用户登陆信息
		models.UpdateAdminUserLoginInfo(tokenId, tokenLastTime, token)
		//设置cookie
		c.Ctx.SetCookie("tokenId", fmt.Sprintf("%d", tokenId), 4200) // 设置cookie
		c.Ctx.SetCookie("token", token, 4200)                        // 设置cookie
		//打点
		ip := c.GetIpInfo()
		c.LogAdminLogin(tokenId, ip)
		//返回
		c.ajaxMsg("登陆成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
退出登陆【方法】
 */
func (c *AdminLoginController) AdminLoginOut() {
	//删除cookie
	c.Ctx.SetCookie("tokenId", "", -1)
	c.Ctx.SetCookie("token", "", -1)
	//跳转返回登陆页
	c.Redirect("/adminLogin", 302)
}
