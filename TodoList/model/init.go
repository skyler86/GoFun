package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

// 配置数据库连接
func Database(connstring string) {
	fmt.Println("connstring",connstring)
	db, err := gorm.Open("mysql",connstring)
	if err != nil {
		fmt.Println(err)
		panic("Mysql数据库连接错误")
	}
	fmt.Print("数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {	// 如果这个Mode是发行版的话就不用输出这个日志
		db.LogMode(false)
	}

	// 设置数据库的参数
	db.SingularTable(true)	// 创建数据库表名时可以默认不加s
	db.DB().SetMaxIdleConns(20)	// 设置连接池
	db.DB().SetMaxOpenConns(100)	// 最大连接数
	db.DB().SetConnMaxLifetime(time.Second*30)	// 连接的时间
	// 将db参数赋值到全局DB变量中
	DB = db
	migration()			//数据库连接时对它进行迁移
}