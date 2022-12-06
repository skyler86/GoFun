package logic		// 可以理解成 service 层
import (
	"TodoList/model"
	"TodoList/pkg/utils"
	"TodoList/serializer"
	"github.com/jinzhu/gorm"
)

// 用户注册服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

// 注册服务的逻辑
func (logic *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?",logic.UserName).
		First(&user).Count(&count)		// user_name不能重复
	if count == 1 {		// 如果count等1，说明已经有这个用户了，不需要再注册了
		return serializer.Response{
			Status:400,
			Msg: "当前用户已存在，无需再注册",
		}
	}
	user.UserName = logic.UserName

	// 密码加密
	if err := user.SetPassword(logic.Password);err!=nil{
		return serializer.Response{
			Status: 400,
			Msg:err.Error(),		// 如果密码有错误,则将错误信息传回去
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err!=nil {
		return serializer.Response{
			Status: 500,
			Msg: "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "用户注册成功！",
	}
}

func (logic *UserService) Login() serializer.Response {
	var user model.User
	// 先在数据库里查找这个user，看看有没有这个用户
	if err := model.DB.Where("user_name=?",logic.UserName).First(&user).Error; err!=nil{
		if gorm.IsRecordNotFoundError(err){		// 如果未找到要查询的记录则返回状态码：400
			return serializer.Response{
				Status: 400,
				Msg: "用户不存在，请先登录",
			}
		}
		//如果不是用户不存在，而是其它不可抗拒的因素导致的查询错误
		return  serializer.Response{
			Status: 500,
			Msg: "数据库查询错误!",
		}
	}

	// 当查到用户存在时再验证密码是否正确
	if user.CheckPassword(logic.Password)==false{
		return serializer.Response{
			Status: 400,
			Msg: "密码错误",
		}
	}

	// 当密码验证成功后要发送一个token，用途是为了其它功能需要身份验证时所用来给前端存储的
	// 比如这个功能是创建一个备忘录，这个功能就要使用到token，否则都不知道谁创建的备忘录。
	// 3.发送token
	token,err := utils.GenerateToken(user.ID,logic.UserName,logic.Password)
	if err != nil {
		return serializer.Response{		// 如果有分发错误就返回状态码500
			Status: 500,
			Msg: "Token签发错误",
		}
	}
	return serializer.Response{		// 如果没有错误则返回200
		Status: 200,
		Data: serializer.TokenData{User: serializer.BuildUser(user),Token:token},		// TokenData是指带token的返回值。把token放进data里
		Msg: "登录成功！",
	}
}