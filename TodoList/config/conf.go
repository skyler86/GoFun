package config

import (
	"TodoList/model"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode string
	HttpPort string
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
)

// 进行mysql数据库配置的读写
func Init() {
	// 读取配置文件位置
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Print("配置文件读取错误，请检查文件路径")
	}

	// 加载配置文件信息
	LoadServer(file)
	LoadMysql(file)
	// 拼接Mysql连接路径
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path)
}

// 加载服务器配置
func LoadServer(file *ini.File){
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

// 加载数据库mysql配置
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}