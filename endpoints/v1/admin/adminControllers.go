package admin

import (
	"github.com/gin-gonic/gin"
)

func Admin(ctx *gin.Context)  {
	ctx.JSON(0, gin.H{})
}
