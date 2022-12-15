package logic

import (
	"TodoList/model"
	"TodoList/serializer"
	"time"
)

// 创建服务
// 定义创建备忘录的结构体
type CreatTaskService struct {
	Title string	`json:"title" form:"title"`		// 标题
	Content string `json:"content" form:"content"`	// 内容
	Status int	`json:"status" form:"status"`		// 状态；0是未做，1是已做
}

// 定义读取备忘录的结构体
type ShowTaskService struct {		// 因为是Get请求所以结构体是空的

}

// 定义展示所有备忘录的结构体
type ListTaskService struct {
	// 做分页的功能
	PageNum int `json:"page_num" form:"page_num"` 		// 页数码
	PageSize int `json:"page_size" form:"page_size"`	// 页大小

}


// 新增一条备忘录的方法
func (logic *CreatTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200			// 初始化定义code为200
	model.DB.First(&user,id)
	task := model.Task{		// 将接过来的参数传到Task模型里
		User:user,
		Uid:user.ID,
		// Title是前端传过来的
		Title:logic.Title,
		Status: 0,
		// Content前端传过来的
		Content: logic.Content,
		StartTime: time.Now().Unix(),	// 开始时间戳
		EndTime: 0,
	}
	err := model.DB.Create(&task).Error	// 错误处理
	if err != nil {
		code = 500		// 发生错误，创建不成功
		return serializer.Response{
			Status: code,
			Msg: "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg: "创建备忘录成功",
	}
}

// 展示一条备忘录的方法
func (logic *ShowTaskService) Show(tid string) serializer.Response {	// tid是备忘录的ID;uid是u用户的id(可以不写)
	var task model.Task
	code := 200
	err := model.DB.First(&task,tid).Error
	if err!=nil {
		code = 500
		return serializer.Response{
			Status:code,
			Msg: "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data: serializer.BuildTask(task),	// 将当前这个任务进行序列化并返回
	}
}

// 列表返回所有备忘录的方法
func (logic *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task		// 切片类型的变量
	count := 0		// 计算出所有备忘录
	if 	logic.PageSize == 0 {		// 先判断分页,如果传过来的是0，则判断它是15页
		logic.PageSize=15
	}

	// 首先用外键将user预加载出来，找到具体的user然后在进行一遍聚类函数
	model.DB.Model(&model.Task{}).Preload("User").Where("uid = ?", uid).Count(&count).	// 首先用外键将user预加载出来，找到具体的user然后在进行一遍聚类函数
		Limit(logic.PageSize).Offset((logic.PageNum - 1)*logic.PageSize).Find(&tasks)			// 最后对它进行分页操作，把这个用户所有备忘录都给找到
	return serializer.BuildListResponse(serializer.BuildTasks(tasks),uint(count))
}