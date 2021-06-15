package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

/**数据库结构**/

/**
文章分类表
 */
type Classify struct {
	Id       int
	Title    string
	Describe string
	Sort     int
	Status   int
	CreateAt int64
	UpdateAt int64
}

/**
文章标签表
 */
type ArticleLabel struct {
	Id       int
	Title    string
	Describe string
	Sort     int
	Status   int
	CreateAt int64
	UpdateAt int64
}

/**
文章表
 */
type Article struct {
	Id       int
	TypeId   int
	UserId   int
	LabelId  int
	Title    string
	Image    string
	Describe string
	Content  string
	Sort     int
	Sign     int
	Url      string
	Status   int
	Thumbs   int
	Hiss     int
	CreateAt int64
	UpdateAt int64
}

/**
文章-评论表
 */
type ArticleComment struct {
	Id                  int
	ArticleId           int
	CommentId           int
	CommentUserNickname string
	UserNickname        string
	UserEmail           string
	Content             string
	Status              int
	Sign                int
	AdminStatus         int
	CreateAt            int64
	UpdateAt            int64
}

/**
管理员表
 */
type AdminUser struct {
	Id          int
	Nickname    string
	Username    string
	Password    string
	Token       string
	TokenOverAt int64
	Status      int
	CreateAt    int64
	UpdateAt    int64
}

/**
变量表
 */
type Variable struct {
	Id       int
	Name     string
	Desc     string
	Value    string
	CreateAt int64
	UpdateAt int64
}

/**
友情链接表
 */
type Friendship struct {
	Id       int
	Title    string
	Url      string
	Describe string
	Sort     int
	Status   int
	CreateAt int64
	UpdateAt int64
}

/**
统计-文章访问记录
 */
type StatisticArticle struct {
	Id         int
	ArticleId  int
	Ip         string
	Address    string
	CreateDate string
	CreateAt   int64
	UpdateAt   int64
}

/**
统计-管理员访问记录
 */
type StatisticAdmin struct {
	Id         int
	AdminId    int
	Ip         string
	Address    string
	CreateDate string
	CreateAt   int64
	UpdateAt   int64
}

/**
注册数据库
 */
func ConnectionDataBase() {
	//注册模型
	orm.RegisterModel(new(Classify), new(ArticleLabel), new(Article), new(AdminUser), new(Variable), new(Friendship), new(StatisticArticle),
		new(StatisticAdmin), new(ArticleComment))
	//获取数据库配置
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPass := beego.AppConfig.String("mysqlPass")
	mysqlDb := beego.AppConfig.String("mysqlDb")
	mysqlUrls := beego.AppConfig.String("mysqlUrls")
	//初始化Orm
	orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPass+"@tcp("+mysqlUrls+":3306)/"+mysqlDb+"?charset=utf8", 30, 30)
	orm.RunSyncdb("default", false, true)
}
