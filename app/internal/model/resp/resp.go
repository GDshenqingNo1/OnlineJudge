package resp

import "github.com/gin-gonic/gin"

func ResponseFail(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  message,
		"ok":   false,
	})
}

func ResponseSuccess(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  message,
		"ok":   true,
	})
}
