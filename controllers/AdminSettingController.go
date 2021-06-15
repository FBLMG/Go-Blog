package controllers

import (
	"personForum/models"
	"strconv"
)

/**
后台-配置模块
 */
type AdminSettingController struct {
	AdminBaseController
}

/**
网站备案
 */

/**
系统配置-获取底部文案
 */
func (c *AdminSettingController) AdminSettingGetBottomConfig() {
	//设置当前页面标题
	c.Data["WebTitle"] = "底部设置"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(7)
	//获取底部配置信息
	webTitle := models.GetVariable("webTitle").Value       //网站名称
	recordTitle := models.GetVariable("recordTitle").Value //备案号
	recordUrl := models.GetVariable("recordUrl").Value     //备案地址
	//赋值数据
	c.Data["webTitle"] = webTitle
	c.Data["recordTitle"] = recordTitle
	c.Data["recordUrl"] = recordUrl
	/** 页面返回**/
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/setting/bottom.tpl"
}

/**
系统配置-配置底部文案
 */
func (c *AdminSettingController) AdminSettingSetBottomConfig() {
	//获取前端传过来的参数
	webTitle := c.GetString("webTitle")       //网站名称
	recordTitle := c.GetString("recordTitle") //备案号
	recordUrl := c.GetString("recordUrl")     //备案地址
	//设置
	models.SetVariable("webTitle", "网站名称", webTitle)
	models.SetVariable("recordTitle", "备案号接", recordTitle)
	models.SetVariable("recordUrl", "备案地址", recordUrl)
	//返回
	c.ajaxMsg("编辑成功", 1)
}

/**
自我介绍
 */

/**
系统配置-获取自我介绍配置
 */
func (c *AdminSettingController) AdminSettingGetUserConfig() {
	//设置当前页面标题
	c.Data["WebTitle"] = "自我介绍配置"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(7)
	//获取自我介绍配置
	userTitle := models.GetVariable("userTitle").Value //自我介绍-标题
	userImage := models.GetVariable("userImage").Value //自我介绍-头像
	userDesc := models.GetVariable("userDesc").Value   //自我介绍-描述
	//判断图片是否空
	if userImage == "" {
		userImage = "static/img/imgUpload.jpeg"
	}
	//赋值数据
	c.Data["userTitle"] = userTitle
	c.Data["userImage"] = userImage
	c.Data["userDesc"] = userDesc
	//图片上传附件
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ThirdPartyJs"] = "admin/comm/uploadSettingImage.tpl" //图书上传封装方法
	c.LayoutSections["ThirdPartyImageJs"] = "admin/comm/uploadImageJs.tpl" //图片上传JS
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/setting/userConfig.tpl"
}

/**
系统配置-配置自我介绍
 */
func (c *AdminSettingController) AdminSettingSetUserConfig() {
	//获取前端传过来的参数
	userTitle := c.GetString("userTitle") //自我介绍-标题
	userImage := c.GetString("userImage") //自我介绍-头像
	userDesc := c.GetString("userDesc")   //自我介绍-描述
	//设置
	models.SetVariable("userTitle", "自我介绍-标题", userTitle)
	models.SetVariable("userImage", "自我介绍-头像", userImage)
	models.SetVariable("userDesc", "自我介绍-描述", userDesc)
	//返回
	c.ajaxMsg("编辑成功", 1)
}

/**
小程序广告
 */

/**
获取广告配置
 */
func (c *AdminSettingController) AdminSettingGetAdConfig() {
	//设置当前页面标题
	c.Data["WebTitle"] = "广告配置"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(7)
	//获取广告配置
	adTitle := models.GetVariable("adTitle").Value                    //广告-标题
	adImage := models.GetVariable("adImage").Value                    //广告-头像
	adStatus, _ := strconv.Atoi(models.GetVariable("adStatus").Value) //广告-状态【字符串转int】
	adUrl := models.GetVariable("adUrl").Value                        //广告-跳转地址
	//赋值数据
	c.Data["adTitle"] = adTitle
	c.Data["adImage"] = adImage
	c.Data["adStatus"] = adStatus
	c.Data["adUrl"] = adUrl
	//图片上传附件
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ThirdPartyJs"] = "admin/comm/uploadSettingImage.tpl" //图书上传封装方法
	c.LayoutSections["ThirdPartyImageJs"] = "admin/comm/uploadImageJs.tpl" //图片上传JS
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "显示"
	StatusList[2] = "隐藏"
	c.Data["StatusList"] = StatusList
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/setting/adConfig.tpl"
}

/**
设置广告配置
 */
func (c *AdminSettingController) AdminSettingSetAdConfig() {
	//获取前端传过来的参数
	adTitle := c.GetString("adTitle")   //广告-标题
	adImage := c.GetString("adImage")   //广告-头像
	adStatus := c.GetString("adStatus") //广告-状态
	adUrl := c.GetString("adUrl")       //广告-跳转地址
	//设置
	models.SetVariable("adTitle", "广告-标题", adTitle)
	models.SetVariable("adImage", "广告-头像", adImage)
	models.SetVariable("adStatus", "广告-状态", adStatus)
	models.SetVariable("adUrl", "广告-跳转地址", adUrl)
	//返回
	c.ajaxMsg("编辑成功", 1)
}

/**
网站名称
 */

/**
获取网站标题配置
 */
func (c *AdminSettingController) AdminSettingGetWebTitleConfig() {
	//设置当前页面标题
	c.Data["WebTitle"] = "广告配置"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(7)
	//获取广告配置
	webTopTitle := models.GetVariable("webTopTitle").Value //网站顶部栏标题
	//赋值数据
	c.Data["webTopTitle"] = webTopTitle
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/setting/webTitleConfig.tpl"
}

/**
设置网站标题配置
 */
func (c *AdminSettingController) AdminSettingSetWebTitleConfig() {
	//获取前端传过来的参数
	webTopTitle := c.GetString("webTopTitle") //广告-标题
	//设置
	models.SetVariable("webTopTitle", "广告配置", webTopTitle)
	//返回
	c.ajaxMsg("编辑成功", 1)
}
