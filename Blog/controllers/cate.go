package controllers


import (
	"Blog/services"
	"Blog/types"
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateCate
func CreateCate(c *gin.Context) {
	name, ok :=c.GetPostForm("name"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	flag, _ := services.CreateCateService(name)
	switch flag {
	case 1:
		utils.Success(c, nil)
	case 0:
		utils.Error(c, int(types.ApiCode.EXISTSNAME), types.ApiCode.GetMessage(types.ApiCode.EXISTSNAME))
	case 2:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	}
}

//GetCateList
func GetCateList(c *gin.Context) {
	cate := services.GetCateListService()
	utils.Success(c, cate)
}

//UpdateCateById
func UpdateCateById(c *gin.Context) {
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
	flag, _ := services.UpdateCateByIdService(newId, name)
	switch flag {
	case 1:
		utils.Success(c, nil)
	case 0:
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
	case 2:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	}
}

//DeleteCateById
func DeleteCateById(c *gin.Context) {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	flag, _ := services.DeleteCateByIdService(newId)
	switch flag {
	case 1:
		utils.Success(c, nil)
	case 2:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	case 0:
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
	}
}

