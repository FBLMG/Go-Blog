package routers

import (
	"personForum/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//入口
	beego.Router("/", &controllers.HomeArticleController{}, "GET:ArticleIndex")

	/////////////////////////////////////      博客【前端】      ///////////////////////////////////////
	//博客列表【页面】
	beego.Router("/articleIndex", &controllers.HomeArticleController{}, "GET:ArticleIndex")
	//根据分类获取博客列表【页面】
	beego.Router("/articleType", &controllers.HomeArticleController{}, "GET:ArticleType")
	//根据标签获取博客列表【页面】
	beego.Router("/articleLabel", &controllers.HomeArticleController{}, "GET:ArticleLabel")
	//获取博客详情【页面】
	beego.Router("/articleInfo", &controllers.HomeArticleController{}, "GET:ArticleInfo")
	//博客评论【方法】
	beego.Router("/articleComment", &controllers.HomeArticleController{}, "POST:ArticleComment")

	/////////////////////////////////////      首页【后端】      ///////////////////////////////////////
	//首页
	beego.Router("/admin", &controllers.AdminIndexController{}, "GET:AdminIndex")

	/////////////////////////////////////      登陆操作【后端】      ///////////////////////////////////////
	//登陆【页面】
	beego.Router("/adminLogin", &controllers.AdminLoginController{}, "GET:AdminLogin")
	//登陆【方法】
	beego.Router("/adminDoLogin", &controllers.AdminLoginController{}, "*:AdminDoLogin")
	//退出登陆【方法】
	beego.Router("/adminLoginOut", &controllers.AdminLoginController{}, "*:AdminLoginOut")

	/////////////////////////////////////      博客【后端】      ///////////////////////////////////////
	//博客列表【页面】
	beego.Router("/adminArticleList", &controllers.AdminArticleController{}, "GET:AdminArticleList")
	//博客添加【页面】
	beego.Router("/adminArticleAdd", &controllers.AdminArticleController{}, "GET:AdminArticleAdd")
	//博客添加【页面】
	beego.Router("/adminArticleInsert", &controllers.AdminArticleController{}, "*:AdminArticleInsert")
	//博客编辑【页面】
	beego.Router("/adminArticleEdit", &controllers.AdminArticleController{}, "GET:AdminArticleEdit")
	//博客编辑【方法】
	beego.Router("/adminArticleUpdate", &controllers.AdminArticleController{}, "*:AdminArticleUpdate")
	//博客编辑状态【方法】
	beego.Router("/adminArticleUpdateStatus", &controllers.AdminArticleController{}, "*:AdminArticleUpdateStatus")
	//博客删除【方法】
	beego.Router("/adminArticleDelete", &controllers.AdminArticleController{}, "*:AdminArticleDelete")
	//博客编辑置顶状态【方法】
	beego.Router("/adminArticleUpdateSign", &controllers.AdminArticleController{}, "*:AdminArticleUpdateSign")
	/////////////////////////////////////      博客分类【后端】      ///////////////////////////////////////
	//博客分类列表【页面】
	beego.Router("/adminArticleTypeList", &controllers.AdminArticleController{}, "GET:AdminArticleTypeList")
	//博客分类添加【页面】
	beego.Router("/adminArticleTypeAdd", &controllers.AdminArticleController{}, "GET:AdminArticleTypeAdd")
	//博客分类新增【方法】
	beego.Router("/adminArticleTypeInsert", &controllers.AdminArticleController{}, "*:AdminArticleTypeInsert")
	//博客分类编辑【页面】
	beego.Router("/adminArticleTypeEdit", &controllers.AdminArticleController{}, "GET:AdminArticleTypeEdit")
	//博客分类编辑【方法】
	beego.Router("/adminArticleTypeUpdate", &controllers.AdminArticleController{}, "*:AdminArticleTypeUpdate")
	//博客分类更新状态【方法】
	beego.Router("/adminArticleTypeUpdateStatus", &controllers.AdminArticleController{}, "*:AdminArticleTypeUpdateStatus")
	//博客分类删除【方法】
	beego.Router("/adminArticleTypeDelete", &controllers.AdminArticleController{}, "*:AdminArticleTypeDelete")

	/////////////////////////////////////      博客标签【后端】      ///////////////////////////////////////
	//博客标签列表【页面】
	beego.Router("/adminArticleLabelList", &controllers.AdminArticleController{}, "GET:AdminArticleLabelList")
	//博客标签添加【页面】
	beego.Router("/adminArticleLabelAdd", &controllers.AdminArticleController{}, "GET:AdminArticleLabelAdd")
	//博客标签添加【方法】
	beego.Router("/adminArticleLabelInsert", &controllers.AdminArticleController{}, "*:AdminArticleLabelInsert")
	//博客标签编辑【页面】
	beego.Router("/adminArticleLabelEdit", &controllers.AdminArticleController{}, "GET:AdminArticleLabelEdit")
	//博客标签编辑【方法】
	beego.Router("/adminArticleLabelUpdate", &controllers.AdminArticleController{}, "*:AdminArticleLabelUpdate")
	//博客标签状态修改【方法】
	beego.Router("/adminArticleLabelUpdateStatus", &controllers.AdminArticleController{}, "*:AdminArticleLabelUpdateStatus")
	//博客标签删除【方法】
	beego.Router("/adminArticleLabelDelete", &controllers.AdminArticleController{}, "*:AdminArticleLabelDelete")

	/////////////////////////////////////      博客评论【后端】      ///////////////////////////////////////
	//修改博主评论状态【方法】
	beego.Router("/articleCommentGetList", &controllers.AdminArticleController{}, "GET:ArticleCommentGetList")
	beego.Router("/articleCommentUpdateStatus", &controllers.AdminArticleController{}, "*:ArticleCommentUpdateStatus")
	beego.Router("/articleCommentReply", &controllers.AdminArticleController{}, "*:ArticleCommentReply")

	/////////////////////////////////////      友情链接【后端】      ///////////////////////////////////////
	//友情链接列表【页面】
	beego.Router("/adminFriendshipList", &controllers.AdminFriendshipController{}, "GET:AdminFriendshipList")
	//友情链接添加【页面】
	beego.Router("/adminFriendshipAdd", &controllers.AdminFriendshipController{}, "GET:AdminFriendshipAdd")
	//友情链接添加【方法】
	beego.Router("/adminFriendshipInsert", &controllers.AdminFriendshipController{}, "*:AdminFriendshipInsert")
	//友情链接编辑【页面】
	beego.Router("/adminFriendshipEdit", &controllers.AdminFriendshipController{}, "GET:AdminFriendshipEdit")
	//友情链接编辑【方法】
	beego.Router("/adminFriendshipUpdate", &controllers.AdminFriendshipController{}, "*:AdminFriendshipUpdate")
	//友情链接更改状态【方法】
	beego.Router("/adminFriendshipUpdateStatus", &controllers.AdminFriendshipController{}, "*:AdminFriendshipUpdateStatus")
	//友情链接删除【方法】
	beego.Router("/adminFriendshipDelete", &controllers.AdminFriendshipController{}, "*:AdminFriendshipDelete")

	/////////////////////////////////////      用户【后端】      ///////////////////////////////////////
	//用户列表【页面】
	beego.Router("/adminUserList", &controllers.AdminUserController{}, "GET:AdminUserList")
	//用户添加【页面】
	beego.Router("/adminUserAdd", &controllers.AdminUserController{}, "GET:AdminUserAdd")
	//用户添加【方法】
	beego.Router("/adminUserInsert", &controllers.AdminUserController{}, "*:AdminUserInsert")
	//用户编辑【页面】
	beego.Router("/adminUserEdit", &controllers.AdminUserController{}, "GET:AdminUserEdit")
	//用户编辑【方法】
	beego.Router("/adminUserUpdate", &controllers.AdminUserController{}, "*:AdminUserUpdate")
	//用户修改状态【方法】
	beego.Router("/adminUserUpdateStatus", &controllers.AdminUserController{}, "*:AdminUserUpdateStatus")
	//用户删除【方法】
	beego.Router("/adminUserDelete", &controllers.AdminUserController{}, "*:AdminUserDelete")

	/////////////////////////////////////      统计【后端】      ///////////////////////////////////////
	//管理员登陆【页面】
	beego.Router("/adminStatisticUserLogin", &controllers.AdminStatisticController{}, "GET:AdminStatisticUserLogin")
	//导出管理员登陆数据【方法】
	beego.Router("/adminStatisticUserLoginExcel", &controllers.AdminStatisticController{}, "*:AdminStatisticUserLoginExcel")
	//删除管理员登陆数据【方法】
	beego.Router("/adminStatisticUserLoginDelete", &controllers.AdminStatisticController{}, "*:AdminStatisticUserLoginDelete")
	//文章访问【页面】
	beego.Router("/adminStatisticArticle", &controllers.AdminStatisticController{}, "GET:AdminStatisticArticle")
	//删除文章访问数据【方法】
	beego.Router("/adminStatisticArticleDelete", &controllers.AdminStatisticController{}, "*:AdminStatisticArticleDelete")
	//导出文章访问数据【方法】
	beego.Router("/adminStatisticArticleExcel", &controllers.AdminStatisticController{}, "*:AdminStatisticArticleExcel")

	/////////////////////////////////////      系统配置【后端】      ///////////////////////////////////////
	//获取底部文案【页面】
	beego.Router("/adminSettingGetBottomConfig", &controllers.AdminSettingController{}, "GET:AdminSettingGetBottomConfig")
	//配置底部文案【方法】
	beego.Router("/adminSettingSetBottomConfig", &controllers.AdminSettingController{}, "*:AdminSettingSetBottomConfig")
	//获取自我介绍【页面】
	beego.Router("/adminSettingGetUserConfig", &controllers.AdminSettingController{}, "GET:AdminSettingGetUserConfig")
	//配置自我介绍【方法】
	beego.Router("/adminSettingSetUserConfig", &controllers.AdminSettingController{}, "*:AdminSettingSetUserConfig")
	//获取广告【页面】
	beego.Router("/adminSettingGetAdConfig", &controllers.AdminSettingController{}, "GET:AdminSettingGetAdConfig")
	//配置广告【方法】
	beego.Router("/adminSettingSetAdConfig", &controllers.AdminSettingController{}, "*:AdminSettingSetAdConfig")
	//获取网站名称【页面】
	beego.Router("/adminSettingGetWebTitleConfig", &controllers.AdminSettingController{}, "GET:AdminSettingGetWebTitleConfig")
	//配置网站名称【方法】
	beego.Router("/adminSettingSetWebTitleConfig", &controllers.AdminSettingController{}, "*:AdminSettingSetWebTitleConfig")

	/////////////////////////////////////      上传处理类【后端】      ///////////////////////////////////////
	//配置封面上传
	beego.Router("/adminUploadSetting", &controllers.AdminUploadController{}, "*:AdminUploadSetting")
	//日志封面上传
	beego.Router("/adminUploadArticle", &controllers.AdminUploadController{}, "*:AdminUploadArticle")

	/*/////////////////////////////////////      【过滤的控制器】      ///////////////////////////////////////
	//首页
	beego.InsertFilter("/admin", beego.BeforeRouter, controllers.HasPermission)
	//博客
	beego.InsertFilter("/adminArticleList", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleAdd", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleInsert", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleEdit", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleUpdate", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleUpdateStatus", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleDelete", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleUpdateSign", beego.BeforeRouter, controllers.HasPermission)
	//博客分类
	beego.InsertFilter("/adminArticleTypeList", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleTypeAdd", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleTypeInsert", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleTypeEdit", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleTypeUpdate", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleTypeUpdateStatus", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleTypeDelete", beego.BeforeRouter, controllers.HasPermission)
	//博客标签
	beego.InsertFilter("/adminArticleLabelList", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleLabelAdd", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleLabelInsert", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleLabelEdit", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleLabelUpdate", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleLabelUpdateStatus", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminArticleLabelDelete", beego.BeforeRouter, controllers.HasPermission)
	//博客评论
	beego.InsertFilter("/articleCommentGetList", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/articleCommentUpdateStatus", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/articleCommentReply", beego.BeforeRouter, controllers.HasPermission)
	//友情链接
	beego.InsertFilter("/adminFriendshipList", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminFriendshipAdd", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminFriendshipInsert", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminFriendshipEdit", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminFriendshipUpdate", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminFriendshipUpdateStatus", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminFriendshipDelete", beego.BeforeRouter, controllers.HasPermission)
	//用户
	beego.InsertFilter("/adminUserList", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminUserAdd", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminUserInsert", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminUserEdit", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminUserUpdate", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminUserUpdateStatus", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminUserDelete", beego.BeforeRouter, controllers.HasPermission)
	//系统配置
	beego.InsertFilter("/adminSettingGetBottomConfig", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminSettingSetBottomConfig", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminSettingGetUserConfig", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminSettingSetUserConfig", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminSettingGetAdConfig", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminSettingSetAdConfig", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminSettingSetWebTitleConfig", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminSettingGetWebTitleConfig", beego.BeforeRouter, controllers.HasPermission)
	//统计
	beego.InsertFilter("/adminStatisticUserLogin", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminStatisticArticle", beego.BeforeRouter, controllers.HasPermission)
	//图片上传
	beego.InsertFilter("/adminUploadSetting", beego.BeforeRouter, controllers.HasPermission)
	beego.InsertFilter("/adminUploadArticle", beego.BeforeRouter, controllers.HasPermission)*/
}
