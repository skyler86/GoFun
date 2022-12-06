package controllers

import (
	"Blog/services"
	"Blog/types"
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateComment
func CreateComment(c *gin.Context) {
	post_id, ok := c.GetPostForm("post_id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_post_id, err := strconv.Atoi(post_id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	comment_content, ok := c.GetPostForm("comment_content"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	email, ok := c.GetPostForm("email"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	ok = services.CreateCommentService(convert_post_id, comment_content, email); if !ok {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

//GetCommentList
func GetCommentList(c *gin.Context) {
	comments := services.GetCommentListService()

	utils.Success(c, comments)
}

//DelCommentById
func DelCommentById(c *gin.Context) {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_id, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	flag, _ := services.DelCommentByIdService(convert_id)
	switch flag {
	case 0:
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	case 1:
		utils.Success(c, nil)
	case 2:
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
}