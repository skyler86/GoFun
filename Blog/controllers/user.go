package controllers

import (
	"fmt"
	"Blog/services"
	"Blog/types"
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateUser
func CreateUser(c *gin.Context) {
	username, ok := c.GetPostForm("username")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	newPassword := utils.EncryMd5(password)

	phone := c.PostForm("phone")
	email := c.PostForm("email")
	state, _ := strconv.Atoi(c.PostForm("state"))

	ok = services.CreateUserService(username, newPassword, phone, email, state)
	if !ok {
		utils.Error(c, int(types.ApiCode.CREATEUSERFAILED), types.ApiCode.GetMessage(types.ApiCode.CREATEUSERFAILED))
		return
	}
	utils.Success(c, nil)

}
//FetchUsers
func GetUsers(c *gin.Context)  {
	userList := services.GetUsersService()
	fmt.Println(userList)

	utils.Success(c, userList)
}

//DeleteUserById
func DeleteUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	fmt.Println("-------" + id + "---------")

	newId, err := strconv.Atoi(id)
	fmt.Printf("%T", newId)
	if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	err, ok = services.DeleteUserByIdService(newId)
	if !ok {
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

//GetUserById
func GetUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	user := services.GetUserByIdService(newId)

	utils.Success(c, user)
}

//UpdateUserById
func UpdateUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	username, ok := c.GetPostForm("username"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newPassword := utils.EncryMd5(password)

	phone := c.PostForm("phone")
	email := c.PostForm("email")

	_, ok, err = services.UpdateUserByIdService(newId, username, newPassword, phone, email); if !ok {
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

//DisableUserById
func DisableUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	_, err = services.DisableUserByIdService(newId); if err != nil {
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	utils.Success(c, nil)
}

//EnableUserById
func EnableUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	_, err = services.EnableUserByIdService(newId); if err != nil {
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	utils.Success(c, nil)
}
