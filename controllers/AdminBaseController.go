package controllers

//引入模块
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"math/rand"
	"time"
	"fmt"
	"io"
	"crypto/md5"
	"strconv"
	"personForum/models"
)

/**后台基础配置**/
type AdminBaseController struct {
	beego.Controller
}

/**函数初始化**/
func (self *AdminBaseController) Prepare() {
	//获取当前路由
	getRoute := self.Ctx.Request.RequestURI
	//获取是否路由白名单
	routeStatus := getRouteStatus(getRoute)
	//判断路由处理
	if routeStatus != true {
		ok := IsLogin(self.Ctx)
		if !ok {
			self.Ctx.Redirect(302, "/adminLogin")
		}
	}
}

/**
获取路由白名单
 */
func getRouteStatus(route string) bool {
	//设置路由白名单
	var whiteRoute = [3]string{"/adminLogin", "/adminDoLogin", "/adminLoginOut"}
	//判断路由处理
	for _, v := range whiteRoute {
		if v == route {
			return true
		}
	}
	//收尾
	return false
}

// 是否POST提交
func (self *AdminBaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//ajax返回
func (self *AdminBaseController) ajaxMsg(msg string, status int) {
	out := make(map[string]interface{})
	out["status"] = status
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

/**
产生随机数
 */
func (self *AdminBaseController) randCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return randCode
}

/**
获取左侧导航栏点击状态
 */
func (self *AdminBaseController) getLeftActive(number int) (map[string]interface{}) {
	//定义初始返回结构
	leftActive := make(map[string]interface{})
	home := ""
	blog := ""
	classify := ""
	label := ""
	friend := ""
	articleComment := ""
	setting := ""
	admin := ""
	statistic := ""
	if number == 1 {
		home = "active"
	} else if number == 2 {
		blog = "active"
	} else if number == 3 {
		classify = "active"
	} else if number == 4 {
		label = "active"
	} else if number == 5 {
		friend = "active"
	} else if number == 6 {
		articleComment = "active"
	} else if number == 7 {
		setting = "active"
	} else if number == 8 {
		admin = "active"
	} else if number == 9 {
		statistic = "active"
	}
	//获取登陆用户信息
	adminId, adminNickname := self.getLoginAdminInfo()
	//获取底部配置
	bottomConfig := self.GetBottomConfig()
	//赋值进结构体
	leftActive["HomeActive"] = home
	leftActive["BlogActive"] = blog
	leftActive["ClassifyActive"] = classify
	leftActive["LabelActive"] = label
	leftActive["FriendActive"] = friend
	leftActive["ArticleCommentActive"] = articleComment
	leftActive["SettingActive"] = setting
	leftActive["AdminActive"] = admin
	leftActive["StatisticActive"] = statistic
	//赋值登陆用户信息
	leftActive["adminId"] = adminId
	leftActive["adminNickname"] = adminNickname
	//赋值底部配置
	leftActive["bottomConfig"] = bottomConfig
	//返回
	return leftActive
}

/**
获取底部配置
 */
func (self *AdminBaseController) GetBottomConfig() (map[string]interface{}) {
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
获取登陆用户信息
 */
func (self *AdminBaseController) getLoginAdminInfo() (int, string) {
	//获取当前登陆者cookie
	tokenId := self.Ctx.GetCookie("tokenId")
	adminId, _ := strconv.Atoi(tokenId)
	//获取用户信息
	adminInfo := models.GetAdminUser(adminId)
	//获取用户昵称
	adminNickname := ""
	if adminInfo.Id > 0 {
		adminNickname = adminInfo.Nickname
	}
	//返回用户信息
	return adminId, adminNickname
}

/**
密码加密算法
 */
func (self *AdminBaseController) encryption(password, username string) string {
	conversion := username + password
	encryption := md5.New()
	io.WriteString(encryption, conversion)                //将str写入到w中
	md5password := fmt.Sprintf("%x", encryption.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5password
}

/**
随机生成token
 */
func generatingToken() string {
	//随机生成token
	cruTime := time.Now().Unix()
	md5String := md5.New()
	fmt.Println("strconv.FormatInt(crutime, 10)-->", strconv.FormatInt(cruTime, 10))
	io.WriteString(md5String, strconv.FormatInt(cruTime, 10))
	token := fmt.Sprintf("%x", md5String.Sum(nil))
	return token
}

/**
获取当前访问Ip
 */
func (self *AdminBaseController) GetIpInfo() string {
	/*ip := strings.Split(self.Ctx.Request.RemoteAddr, ":")[0]
	return ip*/
	ip := self.Ctx.Input.IP()
	return ip
}

///////////////////////////////////////        用户登陆权限【控制是否登陆】      ///////////////////////////////////////////////

/**
获取用户登陆状态【处理用户身份机制】
 */
func IsLogin(ctx *context.Context) (bool) {
	//获取用户tokenId、token
	tokenId := ctx.GetCookie("tokenId")
	token := ctx.GetCookie("token")
	//判断用户是否存在令牌
	if tokenId == "" {
		return false
	}
	if token == "" {
		return false
	}
	//获取用户信息
	adminId, _ := strconv.Atoi(tokenId)
	admin := models.GetAdminUser(adminId)
	if admin.Id <= 0 {
		return false
	}
	//判断令牌是否正确(防止多地用户登陆)
	if token != admin.Token {
		return false
	}
	//判断用户令牌时间是否过期
	timeNow := time.Now().Unix()
	if timeNow > admin.TokenOverAt {
		return false
	}
	//用户在线状态合法并返回正确标识
	return true
}

/**
判断用户是否登陆
 */
/*var HasPermission = func(ctx *context.Context) {
	ok := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/adminLogin")
	} else {
	}
}
*/