# 基于beego开发的个人博客网站

# 1、安装
修改conf目录下app.conf文件数据库配置，然后导入doc目录下所有的sql文件

# 2、使用
直接在项目根目录下执行beego run

# 3、开发前准备

- 由于使用了mysql、腾讯云COS，防止出错，请在开发前安装下面4个扩展包

- go get -u github.com/gin-gonic/gin

- go get -u github.com/tencentyun/cos-go-sdk-v5

- go get -u github.com/jinzhu/gorm

- go get -u github.com/go-sql-driver/mysql
