package admin

import (
	"github.com/Penglq/AuthCenter/libs"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context)  {
	jwt := libs.NewJwt()
	jwt.GenerateToken()
	ctx.JSON(0, gin.H{
		"jwt": "",
	})
}
