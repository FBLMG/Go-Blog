package controllers

//引入模块
import (
	"github.com/astaxie/beego"
	"personForum/models"
	"time"
	"strconv"
)

/**
前端基础配置
 */
type HomeBaseController struct {
	beego.Controller
}

/**
是否POST提交
 */
func (self *HomeBaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//ajax返回
func (self *HomeBaseController) ajaxMsg(msg string, status int) {
	out := make(map[string]interface{})
	out["status"] = status
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

/**
获取当前访问Ip
 */
func (self *HomeBaseController) GetIpInfo() string {
	/*	ip := strings.Split(self.Ctx.Request.RemoteAddr, ":")[0]
		return ip*/
	ip := self.Ctx.Input.IP()
	return ip
}

/**
获取基础数据-基础共用
 */
func (self *HomeBaseController) getBaseData() (map[string]interface{}) {
	//获取分类数据
	articleType := models.GetClassifyListWithHome()
	//获取标签数据
	articleLabel := models.GetArticleLabelListWithHome()
	//获取友链数据
	friendship := models.GetFriendshipListWithHome()
	//获取图片域名
	imageUrl := beego.AppConfig.String("webUrl")
	//获取置顶文章
	signArticleList := models.GetArticleWithSign()
	signArticle := make([]map[string]interface{}, len(signArticleList))
	for k, v := range signArticleList {
		row := make(map[string]interface{})
		//获取分类
		articleTypeTitle := ""
		for _, vType := range articleType {
			if vType.Id == v.TypeId {
				articleTypeTitle = vType.Title
			}
		}
		row["k"] = k + 1
		row["Id"] = v.Id
		row["ArticleTypeTitle"] = articleTypeTitle
		row["Title"] = v.Title
		row["Image"] = imageUrl + v.Image
		row["CreateAt"] = beego.Date(time.Unix(v.CreateAt, 0), "Y-m-d H:s:i")
		signArticle[k] = row
	}
	//获取自我介绍
	userConfig := self.getUserConfig()
	//获取底部配置
	bottomConfig := self.getBottomConfig()
	//获取广告配置
	adConfig := self.getAdConfig()
	//获取网站顶部配置
	topWebTitle := self.getTopWebTitle()
	//组合返回信息
	baseInfo := make(map[string]interface{})
	baseInfo["articleType"] = articleType
	baseInfo["articleLabel"] = articleLabel
	baseInfo["friendship"] = friendship
	baseInfo["signArticle"] = signArticle
	baseInfo["userConfig"] = userConfig
	baseInfo["bottomConfig"] = bottomConfig
	baseInfo["adConfig"] = adConfig
	baseInfo["webTitle"] = topWebTitle
	//返回
	return baseInfo

}

/**
获取自我配置
 */
func (self *HomeBaseController) getUserConfig() (map[string]interface{}) {
	//获取自我介绍配置
	userTitle := models.GetVariable("userTitle").Value //自我介绍-标题
	userImage := models.GetVariable("userImage").Value //自我介绍-头像
	userDesc := models.GetVariable("userDesc").Value   //自我介绍-描述
	//初始化数据
	userConfig := make(map[string]interface{})
	//赋值数据
	userConfig["userTitle"] = userTitle
	userConfig["userImage"] = userImage
	userConfig["userDesc"] = userDesc
	//返回
	return userConfig
}

/**
获取底部配置
 */
func (self *HomeBaseController) getBottomConfig() (map[string]interface{}) {
	//获取底部配置信息
	webTitle := models.GetVariable("webTitle").Value       //网站名称
	recordTitle := models.GetVariable("recordTitle").Value //备案号
	recordUrl := models.GetVariable("recordUrl").Value     //备案地址
	//初始化数据
	bottomConfig := make(map[string]interface{})
	//赋值数据
	bottomConfig["webTitle"] = webTitle
	bottomConfig["recordTitle"] = recordTitle
	bottomConfig["recordUrl"] = recordUrl
	//返回
	return bottomConfig
}

/**
获取广告配置
 */
func (self *HomeBaseController) getAdConfig() (map[string]interface{}) {
	adTitle := models.GetVariable("adTitle").Value                    //广告-标题
	adImage := models.GetVariable("adImage").Value                    //广告-头像
	adStatus, _ := strconv.Atoi(models.GetVariable("adStatus").Value) //广告-状态【字符串转int】
	adUrl := models.GetVariable("adUrl").Value                        //广告-跳转地址
	//初始化数据
	adConfig := make(map[string]interface{})
	//赋值数据
	adConfig["adTitle"] = adTitle
	adConfig["adImage"] = adImage
	adConfig["adStatus"] = adStatus
	adConfig["adUrl"] = adUrl
	//返回
	return adConfig
}

/**
获取网站名称
 */
func (self *HomeBaseController) getTopWebTitle() string {
	webTopTitle := models.GetVariable("webTopTitle").Value
	return webTopTitle
}
