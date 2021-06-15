package controllers

import (
	"log"
	"path"
	"time"
	"fmt"
)

type AdminUploadController struct {
	AdminBaseController
}

/**
后端上传处理类
 */

/**
配置图片上传
 */
func (c *AdminUploadController) AdminUploadSetting() {
	//获取图片信息
	f, h, err := c.GetFile("file")
	if err != nil {
		log.Fatal("获取文件错误 ", err)
		c.ajaxMsg("文件获取失败", -1)
	}
	//清除临时文件
	defer f.Close()
	//获取文件后缀
	var fileSuffix string
	fileSuffix = path.Ext(h.Filename)
	//获取当前时间戳
	timeNow := time.Now().Unix()
	//获取随机数
	randCode := c.randCode()
	//设置保存后的路径
	imageUrl := "static/upload/setting/" + randCode + fmt.Sprintf("%d", timeNow) + fileSuffix
	//保存文件
	c.SaveToFile("file", imageUrl) // 保存位置在 static/upload, 没有文件夹要先创建
	//返回
	c.ajaxMsg(imageUrl, 1)
}

/**
日志封面图片上传
 */
func (c *AdminUploadController) AdminUploadArticle() {
	//获取图片信息
	f, h, err := c.GetFile("file")
	if err != nil {
		log.Fatal("获取文件错误 ", err)
		c.ajaxMsg("文件获取失败", -1)
	}
	//清除临时文件
	defer f.Close()
	//获取文件后缀
	var fileSuffix string
	fileSuffix = path.Ext(h.Filename)
	//获取当前时间戳
	timeNow := time.Now().Unix()
	//获取随机数
	randCode := c.randCode()
	//设置保存后的路径
	imageUrl := "static/upload/article/" + randCode + fmt.Sprintf("%d", timeNow) + fileSuffix
	//保存文件
	c.SaveToFile("file", imageUrl) // 保存位置在 static/upload, 没有文件夹要先创建
	//返回
	c.ajaxMsg(imageUrl, 1)
}
