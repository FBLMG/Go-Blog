package controllers

/**
引入模块
 */
import (
	"personForum/models"
	"time"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
)

/**
后台-文章模块
 */
type AdminArticleController struct {
	AdminBaseController
}

/**
文章分类
 */

/**
博客分类-获取列表【页面】
 */
func (c *AdminArticleController) AdminArticleTypeList() {

	/**设置当前页面标题**/
	c.Data["WebTitle"] = "博客分类"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(3)
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
	result, count := models.PageClassifyWithAdmin(page, limit)
	//初始化数据
	classifyList := make([]map[string]interface{}, len(result))
	//循环数据赋值
	for k, v := range result {
		row := make(map[string]interface{})
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		if v.UpdateAt != 0 {
			row["UpdateAt"] = beego.Date(time.Unix(v.UpdateAt, 0), "Y-m-d H:s:i")
		} else {
			row["UpdateAt"] = ""
		}
		row["Status"] = int64(v.Status)
		classifyList[k] = row
	}
	//组合返回数据
	c.Data["Classify"] = classifyList //返回的列表数组
	c.Data["PageCount"] = count       //获取的数据总条数
	c.Data["PagePage"] = page         //获取的页码
	c.Data["PageLimit"] = PageLimit   //每页数量
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/classify/index.tpl"
}

/**
博客分类-添加分类页面【页面】
 */
func (c *AdminArticleController) AdminArticleTypeAdd() {
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	/**设置当前页面标题**/
	c.Data["WebTitle"] = "添加分类"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(3)
	c.Data["StatusList"] = StatusList
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/classify/addClassify.tpl"
}

/**
博客分类-添加分类【方法】
 */
func (c *AdminArticleController) AdminArticleTypeInsert() {
	//判断用户请求
	if c.isPost() {
		//获取前端传过来的参数
		title := c.GetString("title", "") //标题
		status, err := c.GetInt("status") //状态
		sort, _ := c.GetInt("sort")       //序号
		//判断用户输入
		if title == "" {
			c.ajaxMsg("博客分类标题不能为空", -1)
		}
		if err != nil {
			c.ajaxMsg("博客分类状态不能为空", -1)
		}
		//判断是否存在相同分类
		exitTitle := models.GetClassifyInTitle(title)
		if exitTitle.Id > 0 {
			c.ajaxMsg("存在相同分类", -1)
		}
		//添加博客分类
		classify := new(models.Classify)
		classify.Status = status
		classify.Title = title
		classify.Sort = sort
		classify.CreateAt = time.Now().Unix()
		id := models.InsertClassify(classify)
		if id <= 0 {
			c.ajaxMsg("添加博客分类失败", -1)
		}
		//返回
		c.ajaxMsg("添加博客分类成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客分类-编辑分类页面【页面】
 */
func (c *AdminArticleController) AdminArticleTypeEdit() {
	/**设置当前页面标题**/
	c.Data["WebTitle"] = "编辑博客分类"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(3)
	//获取用户参数
	id, idErr := c.GetInt("id")
	if idErr != nil {
		c.Data["Error"] = "博客分类ID不能为空"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取博客分类信息
	articleType := models.GetClassify(id)
	if articleType.Id <= 0 {
		c.Data["Error"] = "博客分类获取失败"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	//设置返回
	c.Data["articleType"] = articleType
	c.Data["StatusList"] = StatusList
	/** 页面返回**/
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/classify/editClassify.tpl"
}

/**
博客分类-编辑分类【方法】
 */
func (c *AdminArticleController) AdminArticleTypeUpdate() {
	//判断请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //Id
		title := c.GetString("title")           //标题
		sort, _ := c.GetInt("sort")             //序号
		status, statusErr := c.GetInt("status") //状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("id获取失败", -1)
		}
		if title == "" {
			c.ajaxMsg("标题获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//判断分类标题是否已经存在
		exitType := models.GetClassifyInTitle(title)
		if exitType.Id > 0 && exitType.Id != id {
			c.ajaxMsg("分类已存在", -1)
		}
		//编辑分类数据
		articleType := new(models.Classify)
		articleType.Id = id
		articleType.Title = title
		articleType.Sort = sort
		articleType.Status = status
		articleType.CreateAt = exitType.CreateAt
		articleType.UpdateAt = time.Now().Unix()
		returnId := models.UpdateClassify(articleType)
		if returnId <= 0 {
			c.ajaxMsg("博客分类编辑失败", -1)
		}
		//返回
		c.ajaxMsg("编辑博客分类成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客分类-更新状态【方法】
 */
func (c *AdminArticleController) AdminArticleTypeUpdateStatus() {
	//判断是否是Post请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //分类Id
		status, statusErr := c.GetInt("status") //状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//修改分类状态
		returnId := models.UpdateClassifyStatus(id, status)
		if returnId <= 0 {
			c.ajaxMsg("状态修改失败", -1)
		}
		//返回
		c.ajaxMsg("状态修改成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客分类-删除分类【方法】
 */
func (c *AdminArticleController) AdminArticleTypeDelete() {
	//判断请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id") //Id
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		//获取分类数据
		articleTypeData := models.GetClassify(id)
		if articleTypeData.Id <= 0 {
			c.ajaxMsg("分类获取失败", -1)
		}
		//判断是否存在关联
		exitCount := models.GetClassifyArticleCount(id)
		if exitCount > 0 {
			c.ajaxMsg("存在关联数据，不可删除", -1)
		}
		//删除分类
		returnId := models.DeleteClassify(id)
		if returnId <= 0 {
			c.ajaxMsg("分类删除失败", -1)
		}
		//成功返回
		c.ajaxMsg("分类删除成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
文章标签
 */

/**
博客标签-获取标签列表【页面】
 */
func (c *AdminArticleController) AdminArticleLabelList() {
	//赋值表头
	c.Data["WebTitle"] = "博客标签"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(4)
	//设置默认每页数量
	PageLimit := 10
	//获取页码参数
	page, err := c.GetInt("page") //获取页码
	if err != nil {
		page = 1
	}
	limit, err := c.GetInt("limit") //获取每页条数
	if err != nil {
		limit = PageLimit
	}
	//获取数据
	result, count := models.PageArticleLabelWithAdmin(page, limit)
	//重组博客标签数组
	labelList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		if v.UpdateAt != 0 {
			row["UpdateAt"] = beego.Date(time.Unix(v.UpdateAt, 0), "Y-m-d H:s:i")
		} else {
			row["UpdateAt"] = ""
		}
		row["Status"] = int64(v.Status)
		labelList[k] = row
	}
	//组合返回数据
	c.Data["Label"] = labelList     //返回的列表数组
	c.Data["PageCount"] = count     //获取的数据总条数
	c.Data["PagePage"] = page       //获取的页码
	c.Data["PageLimit"] = PageLimit //每页数量
	/** 页面返回**/
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/label/index.tpl"
}

/**
博客标签-添加【页面】
 */
func (c *AdminArticleController) AdminArticleLabelAdd() {
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	c.Data["StatusList"] = StatusList
	//赋值表头
	c.Data["WebTitle"] = "博客标签"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(4)
	//返回页面
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/label/addLabel.tpl"
}

/**
博客标签-添加【方法】
 */
func (c *AdminArticleController) AdminArticleLabelInsert() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		title := c.GetString("title")           //标题
		status, statusErr := c.GetInt("status") //状态
		sort, _ := c.GetInt("sort")             //序号
		//判断用户参数
		if title == "" {
			c.ajaxMsg("标题获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//判断是否存在该标签
		exitData := models.GetArticleLabelInTitle(title)
		if exitData.Id > 0 {
			c.ajaxMsg("该标签已存在", -1)
		}
		//添加数据
		label := new(models.ArticleLabel)
		label.Title = title
		label.Status = status
		label.Sort = sort
		label.CreateAt = time.Now().Unix()
		returnId := models.InsertArticleLabel(label)
		if returnId <= 0 {
			c.ajaxMsg("标签添加失败", -1)
		}
		//返回
		c.ajaxMsg("标签添加成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客标签-编辑【页面】
 */
func (c *AdminArticleController) AdminArticleLabelEdit() {
	//设置表头
	c.Data["WebTitle"] = "编辑博客标签"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(4)
	//获取用户参数
	id, idErr := c.GetInt("id")
	if idErr != nil {
		c.Data["Error"] = "博客标签ID不能为空"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取标签
	articleLabel := models.GetArticleLabel(id)
	if articleLabel.Id <= 0 {
		c.Data["Error"] = "博客标签获取失败"
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
	c.Data["articleLabel"] = articleLabel
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/label/editLabel.tpl"
}

/**
博客标签-方法【方法】
 */
func (c *AdminArticleController) AdminArticleLabelUpdate() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //Id
		title := c.GetString("title")           //标题
		status, statusErr := c.GetInt("status") //状态
		sort, _ := c.GetInt("sort")             //序号
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if title == "" {
			c.ajaxMsg("标题获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//判断标题是否存在
		exitTitle := models.GetArticleLabelInTitle(title)
		if exitTitle.Id > 0 && exitTitle.Id != id {
			c.ajaxMsg("该标签已存在", -1)
		}
		//编辑数据
		articleLabel := new(models.ArticleLabel)
		articleLabel.Id = id
		articleLabel.Title = title
		articleLabel.Sort = sort
		articleLabel.Status = status
		articleLabel.CreateAt = exitTitle.CreateAt
		articleLabel.UpdateAt = time.Now().Unix()
		returnId := models.UpdateArticleLabel(articleLabel)
		//判断编辑是否成功
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
博客标签-修改状态
 */
func (c *AdminArticleController) AdminArticleLabelUpdateStatus() {
	//获取用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //Id
		status, statusErr := c.GetInt("status") //状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//获取标签
		exitData := models.GetArticleLabel(id)
		if exitData.Id <= 0 {
			c.ajaxMsg("标签获取失败", -1)
		}
		//修改标签状态
		returnId := models.UpdateArticleLabelStatus(id, status)
		if returnId <= 0 {
			c.ajaxMsg("状态修改失败", -1)
		}
		//返回
		c.ajaxMsg("状态修改成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客标签-删除【方法】
 */
func (c *AdminArticleController) AdminArticleLabelDelete() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id") //Id
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		//获取标签
		exitData := models.GetArticleLabel(id)
		if exitData.Id <= 0 {
			c.ajaxMsg("标签获取失败", -1)
		}
		//判断是否存在关联
		exitCount := models.GetLabelArticleCount(id)
		if exitCount > 0 {
			c.ajaxMsg("存在关联数据，不可删除", -1)
		}
		//删除标签
		returnId := models.DeleteArticleLabel(id)
		if returnId <= 0 {
			c.ajaxMsg("标签删除失败", -1)
		}
		//返回
		c.ajaxMsg("删除成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
文章
 */

/**
博客-获取文章列表【页面】
 */
func (c *AdminArticleController) AdminArticleList() {
	//设置当前页面标题
	c.Data["WebTitle"] = "文章管理"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(2)
	//设置默认每页数量
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
	//获取用户参数
	searchTitle := c.GetString("searchTitle")                    //搜索-标题
	searchTypeId, searchTypeIdErr := c.GetInt("searchTypeId")    //搜索-分类Id
	searchLabelId, searchLabelIdErr := c.GetInt("searchLabelId") //搜索-标签Id
	searchSign, searchSignErr := c.GetInt("searchSign")          //搜索-置顶
	if searchTitle == "" {
		searchTitle = ""
	}
	if searchTypeIdErr != nil {
		searchTypeId = 0
	}
	if searchLabelIdErr != nil {
		searchLabelId = 0
	}
	if searchSignErr != nil {
		searchSign = 0
	}
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	//获取文章列表
	result, count := models.PageArticleWithAdmin(page, limit, searchTitle, searchTypeId, searchLabelId, searchSign)
	//获取管理员
	adminUser := models.GetAdminUserWithAdmin()
	//重新组合返回数据
	articleList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//获取文章类别
		articleClassify := models.GetClassify(v.TypeId)
		//获取文章标签
		articleLabel := models.GetArticleLabel(v.LabelId)
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
		//重新赋值数组
		row["Id"] = v.Id
		row["Title"] = v.Title
		row["Image"] = imageUrl + v.Image
		row["Status"] = v.Status
		row["Sign"] = v.Sign
		row["Content"] = v.Content
		row["Thumbs"] = v.Thumbs
		row["Hiss"] = v.Hiss
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		row["ClassifyTitle"] = articleClassify.Title
		row["LabelTitle"] = articleLabel.Title
		row["Author"] = adminUserName
		articleList[k] = row
	}
	c.Data["Article"] = articleList //返回的列表数组
	c.Data["PageCount"] = count     //获取的数据总条数
	c.Data["PagePage"] = page       //获取的页码
	c.Data["PageLimit"] = PageLimit //每页数据
	//获取博客分类列表
	articleType := models.GetClassifyListWithAdmin()
	c.Data["ArticleType"] = articleType
	//获取博客标签列表
	articleLabel := models.GetArticleLabelListWithHome()
	c.Data["ArticleLabel"] = articleLabel
	//返回用户搜索参数
	getInput := make(map[string]interface{})
	getInput["searchTitle"] = searchTitle     //搜索-标题
	getInput["searchTypeId"] = searchTypeId   //搜索-分类Id
	getInput["searchLabelId"] = searchLabelId //搜索-标签Id
	getInput["searchSign"] = searchSign       //搜索-置顶
	c.Data["getInput"] = getInput
	//获取置顶列表
	SignList := make(map[int]string, 3)
	SignList[0] = "全部"
	SignList[1] = "是"
	SignList[2] = "否"
	c.Data["SignList"] = SignList
	//样式列表
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ThirdPartyCss"] = "admin/comm/buttonCss.tpl" //
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/article/index.tpl"
}

/**
博客-添加文章【页面】
 */
func (c *AdminArticleController) AdminArticleAdd() {
	//设置当前页面标题
	c.Data["WebTitle"] = "文章管理"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(2)
	//获取博客分类列表
	articleType := models.GetClassifyListWithAdmin()
	c.Data["ArticleType"] = articleType
	//获取博客标签列表
	articleLabel := models.GetArticleLabelListWithAdmin()
	c.Data["ArticleLabel"] = articleLabel
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	c.Data["StatusList"] = StatusList
	//图片上传附件
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ThirdPartyJs"] = "admin/comm/uploadArticleImage.tpl" //图书上传封装方法
	c.LayoutSections["ThirdPartyImageJs"] = "admin/comm/uploadImageJs.tpl" //图片上传JS
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/article/addArticle.tpl"
}

/**
博客-添加文章【方法】
 */
func (c *AdminArticleController) AdminArticleInsert() {
	//判断用户请求
	if c.isPost() {
		//获取前端用户传过来的值
		blogTitle := c.GetString("blogTitle")         //标题
		blogDesc := c.GetString("blogDesc")           //描述
		typeId, errTypeId := c.GetInt("typeId")       //分类ID
		labelId, errLabelId := c.GetInt("labelId")    //标签ID
		blogImageUrl := c.GetString("blogImageUrl")   //封面
		blogUrl := c.GetString("blogUrl")             //原文地址
		container := c.GetString("container")         //内容
		status, errStatus := c.GetInt("status")       //状态
		blogSign, errBlogSign := c.GetInt("blogSign") //置顶
		sort, _ := c.GetInt("sort")                   //排序
		//判断用户输入
		if blogTitle == "" {
			c.ajaxMsg("博客标题不能为空", -1)
		}
		if blogDesc == "" {
			c.ajaxMsg("博客描述不能为空", -1)
		}
		if blogImageUrl == "" {
			c.ajaxMsg("博客封面地址不能为空", -1)
		}
		if container == "" {
			c.ajaxMsg("博客内容不能为空", -1)
		}
		if errStatus != nil {
			c.ajaxMsg("博客状态不能为空", -1)
		}
		if errBlogSign != nil {
			c.ajaxMsg("博客置顶不能为空", -1)
		}
		//获取博客分类
		if errTypeId != nil {
			c.ajaxMsg("博客分类ID不能为空", -1)
		}
		exitType := models.GetClassify(typeId)
		if exitType.Id <= 0 {
			c.ajaxMsg("博客分类获取失败", -1)
		}
		//获取博客标签
		if errLabelId != nil {
			c.ajaxMsg("博客标签ID不能为空", -1)
		}
		exitLabel := models.GetArticleLabel(labelId)
		if exitLabel.Id <= 0 {
			c.ajaxMsg("博客标签获取失败", -1)
		}
		//获取当前登陆用户ID
		tokenId := c.Ctx.GetCookie("tokenId")
		adminId, _ := strconv.Atoi(tokenId)
		//发表博客
		article := new(models.Article)
		article.Title = blogTitle
		article.Describe = blogDesc
		article.TypeId = typeId
		article.LabelId = labelId
		article.Image = blogImageUrl
		article.Url = blogUrl
		article.Content = container
		article.Status = status
		article.Sign = blogSign
		article.Sort = sort
		article.CreateAt = time.Now().Unix()
		article.UserId = adminId
		article.Thumbs = 0
		article.Hiss = 0
		returnId := models.InsertArticle(article)
		if returnId <= 0 {
			c.ajaxMsg("博客发表失败", -1)
		}
		//返回
		c.ajaxMsg("博客发表成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客-编辑文章【页面】
 */
func (c *AdminArticleController) AdminArticleEdit() {
	//设置当前页面标题
	c.Data["WebTitle"] = "文章管理"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(2)
	//获取博客分类列表
	articleType := models.GetClassifyListWithAdmin()
	c.Data["ArticleType"] = articleType
	//获取博客标签列表
	articleLabel := models.GetArticleLabelListWithHome()
	c.Data["ArticleLabel"] = articleLabel
	//获取用户参数
	id, idErr := c.GetInt("id")
	if idErr != nil {
		c.Data["Error"] = "ID不能为空"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取文章详情
	article := models.GetArticle(id)
	if article.Id <= 0 {
		c.Data["Error"] = "文章获取失败"
		c.Layout = "admin/layouts/layout.tpl"
		c.TplName = "admin/error/error.tpl"
		return
	}
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	c.Data["imageUrl"] = imageUrl
	//组合返回数据
	c.Data["Article"] = article
	//获取状态列表
	StatusList := make(map[int]string, 2)
	StatusList[1] = "正常"
	StatusList[2] = "禁用"
	c.Data["StatusList"] = StatusList
	//图片上传附件
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ThirdPartyJs"] = "admin/comm/uploadArticleImage.tpl" //图书上传封装方法
	c.LayoutSections["ThirdPartyImageJs"] = "admin/comm/uploadImageJs.tpl" //图片上传JS
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/article/editArticle.tpl"
}

/**
博客-编辑文章【方法】
 */
func (c *AdminArticleController) AdminArticleUpdate() {
	//判断用户请求
	if c.isPost() {
		//获取前端用户传过来的值
		blogTitle := c.GetString("blogTitle")         //标题
		blogDesc := c.GetString("blogDesc")           //描述
		typeId, errTypeId := c.GetInt("typeId")       //分类ID
		labelId, errLabelId := c.GetInt("labelId")    //标签ID
		blogImageUrl := c.GetString("blogImageUrl")   //封面
		blogUrl := c.GetString("blogUrl")             //原文地址
		container := c.GetString("container")         //内容
		status, errStatus := c.GetInt("status")       //状态
		blogSign, errBlogSign := c.GetInt("blogSign") //置顶
		sort, _ := c.GetInt("sort")                   //排序
		//判断用户输入
		if blogTitle == "" {
			c.ajaxMsg("博客标题不能为空", -1)
		}
		if blogDesc == "" {
			c.ajaxMsg("博客描述不能为空", -1)
		}
		if blogImageUrl == "" {
			c.ajaxMsg("博客封面地址不能为空", -1)
		}
		if container == "" {
			c.ajaxMsg("博客内容不能为空", -1)
		}
		if errStatus != nil {
			c.ajaxMsg("博客状态不能为空", -1)
		}
		if errBlogSign != nil {
			c.ajaxMsg("博客z不能为空", -1)
		}
		//获取博客
		id, errId := c.GetInt("blogId")
		if errId != nil {
			c.ajaxMsg("博客ID不能为空", -1)
		}
		exitBlog := models.GetArticle(id)
		if exitBlog.Id <= 0 {
			c.ajaxMsg("博客获取失败", -1)
		}
		//获取博客分类
		if errTypeId != nil {
			c.ajaxMsg("博客分类ID不能为空", -1)
		}
		exitType := models.GetClassify(typeId)
		if exitType.Id <= 0 {
			c.ajaxMsg("博客分类获取失败", -1)
		}
		//获取博客标签
		if errLabelId != nil {
			c.ajaxMsg("博客标签ID不能为空", -1)
		}
		exitLabel := models.GetArticleLabel(labelId)
		if exitLabel.Id <= 0 {
			c.ajaxMsg("博客标签获取失败", -1)
		}
		//编辑博客
		article := new(models.Article)
		article.Id = id
		article.Title = blogTitle
		article.Describe = blogDesc
		article.TypeId = typeId
		article.LabelId = labelId
		article.Image = blogImageUrl
		article.Url = blogUrl
		article.Content = container
		article.Status = status
		article.Sign = blogSign
		article.Sort = sort
		article.CreateAt = exitBlog.CreateAt
		article.UpdateAt = time.Now().Unix()
		article.UserId = exitBlog.UserId
		article.Thumbs = exitBlog.Thumbs
		article.Hiss = exitBlog.Hiss
		returnId := models.UpdateArticle(article)
		if returnId <= 0 {
			c.ajaxMsg("博客编辑失败", -1)
		}
		//返回
		c.ajaxMsg("博客编辑成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客-修改状态【方法】
 */
func (c *AdminArticleController) AdminArticleUpdateStatus() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")             //Id
		status, statusErr := c.GetInt("status") //状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//获取文章
		exitData := models.GetArticle(id)
		if exitData.Id <= 0 {
			c.ajaxMsg("文章获取失败", -1)
		}
		//修改状态
		returnId := models.UpdateArticleStatus(id, status)
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
博客-删除【方法】
 */
func (c *AdminArticleController) AdminArticleDelete() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id") //Id
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		//获取文章
		exitData := models.GetArticle(id)
		if exitData.Id <= 0 {
			c.ajaxMsg("文章获取失败", -1)
		}
		//删除文章
		returnId := models.DeleteArticle(id)
		if returnId <= 0 {
			c.ajaxMsg("文章删除失败", -1)
		}
		//返回
		c.ajaxMsg("文章删除成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
博客-修改置顶状态【方法】
 */
func (c *AdminArticleController) AdminArticleUpdateSign() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")       //Id
		sign, signErr := c.GetInt("sign") //置顶状态
		//判断用户参数
		if idErr != nil {
			c.ajaxMsg("Id获取失败", -1)
		}
		if signErr != nil {
			c.ajaxMsg("置顶状态获取失败", -1)
		}
		//获取文章
		exitData := models.GetArticle(id)
		if exitData.Id <= 0 {
			c.ajaxMsg("文章获取失败", -1)
		}
		//修改状态
		returnId := models.UpdateArticleSign(id, sign)
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
文章评论
 */

/**
文章评论-获取评论列表
 */
func (c *AdminArticleController) ArticleCommentGetList() {
	//设置当前页面标题
	c.Data["WebTitle"] = "博客评论"
	//获取左侧导航点击状态
	c.Data["LeftActive"] = c.getLeftActive(6)
	//获取参数
	searchStatus, searchStatusErr := c.GetInt("searchStatus")
	if searchStatusErr != nil {
		searchStatus = 0
	}
	searchAdminStatus, searchAdminStatusErr := c.GetInt("searchAdminStatus")
	if searchAdminStatusErr != nil {
		searchAdminStatus = 0
	}
	searchCommentId, searchCommentIdErr := c.GetInt("searchCommentId")
	if searchCommentIdErr != nil {
		searchCommentId = 0
	}
	searchArticleId, searchArticleIdErr := c.GetInt("searchArticleId")
	if searchArticleIdErr != nil {
		searchArticleId = 0
	}
	//获取文章
	articleList := models.GetArticleWithAdmin()
	articleListSelect := make([]map[string]interface{}, len(articleList))
	for k, v := range articleList {
		row := make(map[string]interface{})
		//处理标题
		ArticleTitle := strings.Count(v.Title, "") - 1
		if ArticleTitle > 20 {
			rs := []rune(v.Title)
			row["Title"] = string(rs[0:20]) + "..."
		} else {
			row["Title"] = v.Title
		}
		//重新赋值数组
		row["Id"] = v.Id
		articleListSelect[k] = row
	}
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
	//获取评论列表
	result, count := models.PageArticleCommentWithAdmin(page, limit, searchStatus, searchAdminStatus, searchCommentId, searchArticleId)
	//重组评论数组
	articleCommentList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		//获取文章标题
		articleTitle := ""
		for _, vArticle := range articleList {
			if vArticle.Id == v.ArticleId {
				articleTitle = vArticle.Title
				continue
			}
		}
		//处理标题
		ArticleTitle := strings.Count(articleTitle, "") - 1
		if ArticleTitle > 20 {
			rs := []rune(articleTitle)
			row["ArticleTitle"] = string(rs[0:20]) + "..."
		} else {
			row["ArticleTitle"] = articleTitle
		}
		//重新赋值数组
		row["Id"] = v.Id
		row["ArticleId"] = v.ArticleId
		row["CommentId"] = v.CommentId
		row["UserNickname"] = v.UserNickname
		row["UserEmail"] = v.UserEmail
		row["Content"] = v.Content
		row["AdminStatus"] = v.AdminStatus
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		row["Status"] = int64(v.Status)
		articleCommentList[k] = row
	}
	//获取状态列表
	statusList := [4]string{"全部", "显示", "待审核", "隐藏"}
	//是否是管理员
	adminStatus := [3]string{"全部", "是", "否"}
	//组合返回数据
	c.Data["ArticleCommentList"] = articleCommentList //评论列表
	c.Data["PageCount"] = count                       //获取的数据总条数
	c.Data["PagePage"] = page                         //获取的页码
	c.Data["PageLimit"] = PageLimit                   //每页数量
	c.Data["StatusList"] = statusList                 //状态列表
	c.Data["AdminStatus"] = adminStatus               //是否是管理员
	c.Data["ArticleList"] = articleListSelect         //文章列表
	//返回用户搜索参数
	getInput := make(map[string]interface{})
	getInput["SearchStatus"] = searchStatus           //搜索-状态
	getInput["SearchAdminStatus"] = searchAdminStatus //搜索-是否管理员
	getInput["SearchArticleId"] = searchArticleId     //搜索-文章Id
	c.Data["getInput"] = getInput                     //用户搜索参数
	//页面返回
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/articleComment/index.tpl"
	//
}

/**
文章评论-修改状态
 */
func (c *AdminArticleController) ArticleCommentUpdateStatus() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		id, idErr := c.GetInt("id")
		if idErr != nil {
			c.ajaxMsg("id获取失败", -1)
		}
		status, statusErr := c.GetInt("status")
		if statusErr != nil {
			c.ajaxMsg("状态获取失败", -1)
		}
		//获取评论
		commentInfo := models.GetArticleComment(id)
		if commentInfo.Id <= 0 {
			c.ajaxMsg("评论获取失败", -1)
		}
		//修改评论状态
		returnId := models.UpdateArticleCommentStatus(id, status)
		if returnId <= 0 {
			c.ajaxMsg("修改失败", -1)
		}
		//返回
		c.ajaxMsg("修改成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}

/**
文章评论-评论
 */
func (c *AdminArticleController) ArticleCommentReply() {
	//判断用户请求
	if c.isPost() {
		//获取用户参数
		articleId, articleIdErr := c.GetInt("articleId") //文章ID
		commentId, commentIdErr := c.GetInt("commentId") //评论ID
		content := c.GetString("content")                //内容
		//判断用户参数
		if articleIdErr != nil {
			c.ajaxMsg("文章Id获取失败", -1)
		}
		if commentIdErr != nil {
			commentId = 0
		}
		if content == "" {
			c.ajaxMsg("内容获取失败", -1)
		}
		//初始化数据
		userNickname := "佛布朗斯基"
		userEmail := "657271672@qq.com"
		CommentUserNickname := ""
		//是否有被评论人信息
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
		articleComment.Status = 1
		articleComment.Sign = 1
		articleComment.AdminStatus = 1
		articleComment.CreateAt = time.Now().Unix()
		returnId := models.InsertArticleComment(articleComment)
		if returnId <= 0 {
			c.ajaxMsg("评论失败！", -1)
		}
		//返回
		c.ajaxMsg("评论成功", 1)
	} else {
		c.ajaxMsg("非法请求", -2)
	}
}
