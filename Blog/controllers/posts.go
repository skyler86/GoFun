package controllers

import (
	"Blog/services"
	"Blog/types"
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreatePost
func CreatePost(c *gin.Context) {
	title, ok := c.GetPostForm("title"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	desc := c.PostForm("desc")
	content, ok := c.GetPostForm("content"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	author, ok := c.GetPostForm("author"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	tags := c.PostForm("tags")
	category := c.PostForm("category")

	newCategory, err := strconv.Atoi(category); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	ok = services.CreatePostService(title, desc, content, author, tags, newCategory); if !ok {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

//GetPostsList
func GetPostsList(c *gin.Context) {
	posts := services.GetPostsListService()

	utils.Success(c, posts)
}

//UpdatePostById
func UpdatePostById(c *gin.Context) {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	title, ok := c.GetPostForm("title"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	desc := c.PostForm("desc")
	content, ok := c.GetPostForm("content"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	author, ok := c.GetPostForm("author"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	tags := c.PostForm("tags")

	_, flag, _ := services.UpdatePostByIdService(newId, title, desc, content, author, tags)
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

//DeletePostById
func DeletePostById(c *gin.Context) {
	id, ok := c.GetPostForm("id"); if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	flag, _ := services.DeletePostByIdService(newId)
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
