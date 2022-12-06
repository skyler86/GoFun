package controllers

import (
	"Blog/services"
	"Blog/types"
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateTags
func CreateTags(c *gin.Context)  {
	name, ok := c.GetPostForm("name"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	_, ok, res := services.CreateTagsService(name); if !ok {
		utils.Error(c, int(types.ApiCode.EXISTSNAME), types.ApiCode.GetMessage(types.ApiCode.EXISTSNAME))
		return
	}

	if !res {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

//GetTagsList
func GetTagsList(c *gin.Context)  {
	tags := services.GetTagsListService()

	utils.Success(c, tags)
}

//UpdateTagsById
func UpdateTagsById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	name, ok := c.GetPostForm("name"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	_, flag := services.UpdateTagsByIdService(newId, name)
	switch flag {
	case 0:
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
	case 2:
		utils.Error(c, int(types.ApiCode.EXISTSNAME), types.ApiCode.GetMessage(types.ApiCode.EXISTSNAME))
	case 3:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	case 1:
		utils.Success(c, nil)
	}
}

//DeleteTagsById
func DeleteTagsById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	_, flag := services.DeleteTagsByIdService(newId)

	switch flag {
	case 1:
		utils.Success(c, nil)
	case 2:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	case 0:
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
	}
}
