package controllers

import (
	"Blog/services"
	"Blog/types"
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateLink
func CreateLink(c *gin.Context)  {
	name, ok := c.GetPostForm("name"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	url, ok := c.GetPostForm("url"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	ok = services.CreateLinkService(name, url); if !ok {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

//GetLinkList
func GetLinkList(c *gin.Context)  {
	links := services.GetLinkListService()

	utils.Success(c, links)
}

//UpdateLinkById
func UpdateLinkById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_id, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	name, ok := c.GetPostForm("name"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	url, ok := c.GetPostForm("url"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	_, flag := services.UpdateLinkByIdService(convert_id, name, url)
	switch flag {
	case 0:
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	case 1:
		utils.Success(c, nil)
		return
	case 2:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
}

//DeleteLinkById
func DeleteLinkById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_id, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	res := services.DeleteLinkByIdService(convert_id)
	switch res {
	case 0:
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	case 1:
		utils.Success(c, nil)
		return
	case 2:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
}
