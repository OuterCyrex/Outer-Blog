package version1

import (
	"gin-vue-blog/utils"
	"gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, header, _ := c.Request.FormFile("file")
	fileSize := header.Size
	url, code := utils.UploadFile(file, fileSize)
	if code != errmsg.SUCCESS {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"url":  "",
		})
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"url":  url,
	})
}
