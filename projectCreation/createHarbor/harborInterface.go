package createHarbor

import (
	"context"
	"fmt"
	"github.com/mittwald/goharbor-client/v4/apiv1"
	"projectcreation/config"

	"log"
	"strings"
)

type harborConfig struct {
	url        string
	userName   string
	userPasswd string
}

func Logon(projectName string) {
	var hc harborConfig
	hc.url, hc.userName, hc.userPasswd = config.EnVarInfo("HARBOR_API_URL", "HARBOR_USERNAME", "HARBOR_PASSWORD")

	hc.url = config.YamlInfo().GetString("harbor.URL")
	hc.userName = config.YamlInfo().GetString("harbor.USERNAME")
	hc.userPasswd = config.YamlInfo().GetString("harbor.PASSWORD")

	fmt.Println("输出harbor的url: ", hc.url)
	fmt.Println("输出harbor的用户名: ", hc.userName)
	fmt.Println("输出harbor的用户密码: ", hc.userPasswd)

	harborClient, error := apiv1.NewRESTClientForHost(hc.url, hc.userName, hc.userPasswd)
	if error != nil {
		log.Println(error)
	}

	////将harbo仓库里项目的名字全部打印出来:
	//value, error := harborClient.ListProjects(context.TODO(), "")
	//for _, v := range value {
	//	println(v.Name)
	//}

	CreateProject(projectName, harborClient)
}

func CreateProject(projectName string, harborClient *apiv1.RESTClient) {
	//fmt.Printf("projectName=%v \n", projectName)
	var countLimit int
	var storageLimit int

	result1, err := harborClient.NewProject(context.TODO(), projectName, countLimit, storageLimit)
	result2, err := harborClient.NewProject(context.TODO(), CreateAgain(projectName), countLimit, storageLimit)

	if err != nil {
		log.Println("项目创建失败:", err.Error())
	} else {
		fmt.Println("项目创建成功:", result1)
		fmt.Println("项目创建成功:", result2)
	}
}

func CreateAgain(pn string) string {
	npn := []string{"autosync", pn}
	newProjectName := strings.Join(npn, "-")
	return newProjectName
}
