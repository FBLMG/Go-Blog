package controllers

import (
	"github.com/astaxie/beego"
	"personForum/models"
	"strings"
	"time"
)

/**
前端-文章模块
 */
type HomeArticleController struct {
	HomeBaseController
	LogController
}

/**
前台-获取文章列表
 */
func (c *HomeArticleController) ArticleIndex() {
	//每页默认数量
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
	//根据页码获取数据
	result, count := models.PageArticleWithHome(page, limit)
	//获取管理员
	adminUser := models.GetAdminUserWithAdmin()
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	//获取文章分类
	articleTypeList := models.GetClassifyListWithAdmin()
	//重新组合返回数据
	articleList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//获取作者
		adminUserName := ""
		for _, vAdmin := range adminUser {
			if vAdmin.Id == v.UserId {
				adminUserName = vAdmin.Nickname
				continue
			}
		}
		//获取文章分类
		articleType := ""
		for _, vType := range articleTypeList {
			if vType.Id == v.TypeId {
				articleType = vType.Title
				continue
			}
		}
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["articleType"] = articleType
		row["Describe"] = v.Describe
		row["Image"] = imageUrl + v.Image
		row["Status"] = v.Status
		row["Sign"] = v.Sign
		row["Content"] = v.Content
		row["Thumbs"] = v.Thumbs
		row["Hiss"] = v.Hiss
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		row["Author"] = adminUserName
		articleList[k] = row
	}
	//获取公共基础数据
	baseData := c.getBaseData()
	//组合返回数据
	c.Data["BaseData"] = baseData   //公共基础数据
	c.Data["Article"] = articleList //返回的列表数组
	c.Data["PageCount"] = count     //获取的数据总条数
	c.Data["PagePage"] = page       //获取的页码
	c.Data["PageLimit"] = PageLimit //每页数据
	c.Data["ContentCss"] = ""       //容器样式
	//页面返回
	c.Layout = "home/layouts/layout.tpl"
	c.TplName = "home/article/index.tpl"
}

/**
前台-根据分类Id获取文章列表
 */
func (c *HomeArticleController) ArticleType() {
	//每页默认数量
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
	//获取分类Id
	articleTypeId, articleTypeIdErr := c.GetInt("articleTypeId") //分类Id
	if articleTypeIdErr != nil {
		articleTypeId = 0
	}
	//根据页码获取数据
	result, count := models.PageArticleWithHomeClassify(articleTypeId, page, limit)
	//获取管理员
	adminUser := models.GetAdminUserWithAdmin()
	//获取文章分类
	articleTypeList := models.GetClassifyListWithAdmin()
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	//重新组合返回数据
	articleList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//获取作者
		adminUserName := ""
		for _, vAdmin := range adminUser {
			if vAdmin.Id == v.UserId {
				adminUserName = vAdmin.Nickname
				continue
			}
		}
		//获取文章分类
		articleType := ""
		for _, vType := range articleTypeList {
			if vType.Id == v.TypeId {
				articleType = vType.Title
				continue
			}
		}
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["articleType"] = articleType
		row["Describe"] = v.Describe
		row["Image"] = imageUrl + v.Image
		row["Status"] = v.Status
		row["Sign"] = v.Sign
		row["Content"] = v.Content
		row["Thumbs"] = v.Thumbs
		row["Hiss"] = v.Hiss
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		row["Author"] = adminUserName
		articleList[k] = row
	}
	//获取公共基础数据
	baseData := c.getBaseData()
	//组合返回数据
	c.Data["BaseData"] = baseData           //公共基础数据
	c.Data["Article"] = articleList         //返回的列表数组
	c.Data["PageCount"] = count             //获取的数据总条数
	c.Data["PagePage"] = page               //获取的页码
	c.Data["PageLimit"] = PageLimit         //每页数据
	c.Data["articleTypeId"] = articleTypeId //分类Id
	c.Data["ContentCss"] = ""               //容器样式
	//页面返回
	c.Layout = "home/layouts/layout.tpl"
	c.TplName = "home/article/typeList.tpl"
}

/**
前台-根据标签Id获取文章列表
 */
func (c *HomeArticleController) ArticleLabel() {
	//每页默认数量
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
	//获取标签Id
	labelId, labelIdErr := c.GetInt("labelId") //分类Id
	if labelIdErr != nil {
		labelId = 0
	}
	//根据页码获取数据
	result, count := models.PageArticleWithHomeLabel(labelId, page, limit)
	//获取管理员
	adminUser := models.GetAdminUserWithAdmin()
	//获取文章标签
	articleLabelList := models.GetArticleLabelListWithAdmin()
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	//重新组合返回数据
	articleList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//优化文章描述
		desc := strings.Count(v.Describe, "") - 1
		if desc > 20 {
			rs := []rune(v.Describe)
			row["Describe"] = string(rs[0:20]) + "..."
		} else {
			row["Describe"] = v.Describe
		}
		//获取作者
		adminUserName := ""
		for _, vAdmin := range adminUser {
			if vAdmin.Id == v.UserId {
				adminUserName = vAdmin.Nickname
				continue
			}
		}
		//获取文章分类
		articleLabel := ""
		for _, vType := range articleLabelList {
			if vType.Id == v.LabelId {
				articleLabel = vType.Title
				continue
			}
		}
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["articleLabel"] = articleLabel
		row["Describe"] = v.Describe
		row["Image"] = imageUrl + v.Image
		row["Status"] = v.Status
		row["Sign"] = v.Sign
		row["Content"] = v.Content
		row["Thumbs"] = v.Thumbs
		row["Hiss"] = v.Hiss
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		row["Author"] = adminUserName
		articleList[k] = row
	}
	//获取公共基础数据
	baseData := c.getBaseData()
	//组合返回数据
	c.Data["BaseData"] = baseData   //公共基础数据
	c.Data["Article"] = articleList //返回的列表数组
	c.Data["PageCount"] = count     //获取的数据总条数
	c.Data["PagePage"] = page       //获取的页码
	c.Data["PageLimit"] = PageLimit //每页数据
	c.Data["labelId"] = labelId     //标签Id
	c.Data["ContentCss"] = ""       //容器样式
	//页面返回
	c.Layout = "home/layouts/layout.tpl"
	c.TplName = "home/article/labelList.tpl"
}

/**
获取文章详情
 */
func (c *HomeArticleController) ArticleInfo() {
	//获取用户参数
	id, idErr := c.GetInt("id")
	if idErr != nil {
		c.Data["Error"] = "ID不能为空"
		c.TplName = "home/error/error.tpl"
		return
	}
	//获取文章
	articleInfo := models.GetArticle(id)
	if articleInfo.Id <= 0 {
		c.Data["Error"] = "文章获取失败"
		c.TplName = "home/error/error.tpl"
		return
	}
	//判断是否开启状态
	if articleInfo.Status != 1 {
		c.Data["Error"] = "文章获取失败"
		c.TplName = "home/error/error.tpl"
		return
	}
	//获取作者
	adminUser := models.GetAdminUser(articleInfo.UserId)
	//获取分类
	articleType := models.GetClassify(articleInfo.TypeId)
	//获取标签
	articleLabel := models.GetArticleLabel(articleInfo.LabelId)
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	//处理数据
	article := make(map[string]interface{})
	article["Id"] = articleInfo.Id
	article["Title"] = articleInfo.Title
	article["AdminUser"] = adminUser.Nickname
	article["Image"] = imageUrl + articleInfo.Image
	article["Describe"] = articleInfo.Describe
	article["Content"] = articleInfo.Content
	article["Thumbs"] = articleInfo.Thumbs
	article["Hiss"] = articleInfo.Hiss
	article["ArticleType"] = articleType.Title
	article["ArticleLabel"] = articleLabel.Title
	article["CreateAt"] = beego.Date(time.Unix(articleInfo.CreateAt, 0), "Y-m-d H:s:i")
	//获取公共基础数据
	baseData := c.getBaseData()
	//获取相同分类下文章
	identicalArticle := models.GetArticleIdenticalWithHomeNew(articleInfo.TypeId, articleInfo.Id)
	identicalArticleList := make([]map[string]interface{}, len(identicalArticle))
	for k, v := range identicalArticle {
		row := make(map[string]interface{})
		//获取作者
		adminUser := models.GetAdminUser(articleInfo.UserId)
		//获取分类
		articleType := models.GetClassify(articleInfo.TypeId)
		//组合数据
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["Image"] = imageUrl + v.Image
		row["AdminUser"] = adminUser.Nickname
		row["ArticleType"] = articleType.Title
		row["ArticleTypeId"] = v.TypeId
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		//判断是否可以加入
		identicalArticleList[k] = row
	}
	//每页默认数量
	PageLimit := 4
	//获取用户参数
	page, err := c.GetInt("page") //获取页码
	if err != nil {
		page = 1
	}
	limit, err := c.GetInt("limit") //获取每页条数
	if err != nil {
		limit = PageLimit
	}
	//获取评论
	articleCommentResult, articleCommentCount := models.PageArticleCommentWithHome(page, limit, id)
	articleCommentList := make([]map[string]interface{}, len(articleCommentResult))
	for k, v := range articleCommentResult {
		row := make(map[string]interface{})
		//处理评论内容，追加被评论人信息
		if v.CommentId > 0 {
			row["Content"] = "@" + v.CommentUserNickname + "  " + v.Content
		} else {
			row["Content"] = v.Content
		}
		row["Id"] = v.Id
		row["UserNickname"] = v.UserNickname
		row["UserEmail"] = v.UserEmail
		row["AdminStatus"] = v.AdminStatus
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		articleCommentList[k] = row
	}
	//打点
	ip := c.GetIpInfo()
	c.LogArticleInfo(id, ip)
	//组合返回数据
	c.Data["BaseData"] = baseData                         //公共基础数据
	c.Data["ArticleInfo"] = article                       //文章内容
	c.Data["ContentCss"] = "themelazer-content-area"      //容器样式
	c.Data["IdenticalArticleList"] = identicalArticleList //相同分类文章
	c.Data["ArticleCommentList"] = articleCommentList     //评论列表
	c.Data["PageCount"] = articleCommentCount             //评论列表数据总条数
	c.Data["PagePage"] = page                             //获取的页码
	c.Data["PageLimit"] = PageLimit                       //每页数据
	//页面返回
	c.Layout = "home/layouts/layout.tpl"
	c.TplName = "home/article/info.tpl"
}

/**
评论文章
 */
func (c *HomeArticleController) ArticleComment() {
	//判断用户请求
	if c.isPost() {
		//获取用户评论状态
		commentToken := c.Ctx.GetCookie("commentToken")
		if commentToken == "" {
			c.Ctx.SetCookie("commentToken", "10086", 300)
		} else {
			c.ajaxMsg("短时间内不要频繁评论", -1)
		}
		//获取用户参数
		articleId, articleIdErr := c.GetInt("articleId")  //文章Id
		commentId, commentIdErr := c.GetInt("commentId")  //评论Id
		userNickname := c.GetString("userNickname")       //用户昵称
		userEmail := c.GetString("userEmail")             //用户邮箱
		content := c.GetString("content")                 //用户评论内容
		commentUserName := c.GetString("commentUserName") //评论用户昵称
		//判断用户参数
		if articleIdErr != nil {
			c.ajaxMsg("文章Id获取失败", -1)
		}
		if commentIdErr != nil {
			commentId = 0
		}
		if userNickname == "" {
			c.ajaxMsg("用户昵称获取失败", -1)
		}
		if userEmail == "" {
			c.ajaxMsg("用户邮箱获取失败", -1)
		}
		if content == "" {
			c.ajaxMsg("用户内容获取失败", -1)
		}
		//处理内容【去除前缀】
		if commentUserName != "" {
			content = strings.Trim(content, commentUserName)
		}
		//获取文章
		article := models.GetArticle(articleId)
		if article.Id <= 0 {
			c.ajaxMsg("文章获取失败", -1)
		}
		//判断是否有被评论人信息
		CommentUserNickname := ""
		if commentId > 0 {
			commentInfo := models.GetArticleComment(commentId)
			if commentInfo.Id > 0 {
				CommentUserNickname = commentInfo.UserNickname
			}
		}
		//添加评论
		articleComment := new(models.ArticleComment)
		articleComment.ArticleId = articleId
		articleComment.CommentId = commentId
		articleComment.CommentUserNickname = CommentUserNickname
		articleComment.UserNickname = userNickname
		articleComment.UserEmail = userEmail
		articleComment.Content = content
		articleComment.Status = 2
		articleComment.Sign = 2
		articleComment.AdminStatus = 2
		articleComment.CreateAt = time.Now().Unix()
		returnId := models.InsertArticleComment(articleComment)
		if returnId <= 0 {
			c.ajaxMsg("评论失败！", -1)
		}
		//返回
		c.ajaxMsg("评论成功，请等待博主审核！", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}
