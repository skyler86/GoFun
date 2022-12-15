package controller

import (
	"TodoList/logic"
	"TodoList/pkg/utils"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// 定义接口
func CreateTask(c *gin.Context)  {	// 将上下文传过来
	var createTask logic.CreatTaskService
	claim,_ := utils.ParseToken(c.GetHeader("Authorization"))		// 验证身份
	if err := c.ShouldBind(&createTask);err == nil {		// 进行绑定
		res := createTask.Create(claim.Id)
		c.JSON(200,res)
	}else {		// 如果有错误就会打印错误，并返回日志
		logging.Error(err)
		c.JSON(400,ErrorResponse(err))
	}
}

func ShowTask(c *gin.Context)  {
	var showTask logic.ShowTaskService
	// claim,_ := utils.ParseToken(c.GetHeader("Authorization"))		// 验证身份
	if err := c.ShouldBind(&showTask);err == nil {		// 进行绑定
		res := showTask.Show(c.Param("id"))		// c.Param(id)就是备忘录的ID，也就是接收前端传过来的ID;claim.Id就是从请求头Authorization这里拿到的用户ID(可以不写)
		c.JSON(200,res)
	}else {		// 如果有错误就会打印错误，并返回日志
		logging.Error(err)
		c.JSON(400,ErrorResponse(err))
	}
}

func ListTask(c *gin.Context)  {
	var listTask logic.ListTaskService
	claim,_ := utils.ParseToken(c.GetHeader("Authorization"))		// 验证身份
	if err := c.ShouldBind(&listTask);err == nil {		// 进行绑定
		res := listTask.List(claim.Id)		// 从Authorization里面解析出当前访问的是哪个用户
		c.JSON(200,res)
	}else {		// 如果有错误就会打印错误，并返回日志
		logging.Error(err)
		c.JSON(400,ErrorResponse(err))
	}
}