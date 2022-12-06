package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//func EnVarInfo(URL, USERNAME, PASSWORD string) (string, string, string) {
//	//for _, v := range os.Environ() {
//	//	//输出系统所有环境变量的值
//	//	fmt.Println(v)
//	//}
//	envIP := os.Getenv(URL)
//	envUserName := os.Getenv(USERNAME)
//	envPassWord := os.Getenv(PASSWORD)
//	if envIP == "" {
//		fmt.Println("missing address")
//	} else {
//		fmt.Println("os获取环境变量上的URL是:", envIP)
//	}
//	if envUserName == "" {
//		fmt.Println("missing UserName")
//	} else {
//		fmt.Println("os获取环境变量上的用户名是:", envUserName)
//	}
//	if envPassWord == "" {
//		fmt.Println("missing PassWord")
//	} else {
//		fmt.Println("os获取环境变量上的用户密码是:", envPassWord)
//	}
//	return envIP, envUserName, envPassWord
//}

func EnVarInfo(URL, USERNAME, PASSWORD string) (string, string, string) {
	//表示 先预加载匹配的环境变量
	viper.AutomaticEnv()
	//读取已经加载到default中的环境变量
	envIP := viper.GetString(URL)
	envUserName := viper.GetString(USERNAME)
	envPassWord := viper.GetString(PASSWORD)
	if envIP == "" {
		fmt.Println("missing address")
	} else {
		fmt.Println("viper获取环境变量上的URL是:", envIP)
	}
	if envUserName == "" {
		fmt.Println("missing address")
	} else {
		fmt.Println("viper获取环境变量上的用户名是:", envUserName)
	}
	if envPassWord == "" {
		fmt.Println("missing address")
	} else {
		fmt.Println("viper获取环境变量上的用户密码是:", envPassWord)
	}
	return envIP, envUserName, envPassWord

}

func YamlInfo() *viper.Viper {
	//创建viper实例
	config := viper.New()

	//指定配置文件的信息
	config.AddConfigPath("./config/") // 添加配置文件的路径
	config.SetConfigName("config")    // 设置配置文件的名称
	config.SetConfigType("yaml")      // 设置配置文件的格式

	//判断配置文件的信息是否有误
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("找不到配置文件..")
		} else {
			log.Println("配置文件出错..")
		}
		log.Fatal(err)
	}

	//if err := config.ReadInConfig(); err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		fmt.Println("找不到配置文件..")
	//	} else {
	//		fmt.Println("配置文件出错..")
	//	}
	//}else {
	//	fmt.Println("配置文件已找到..")
	//}

	return config
}
