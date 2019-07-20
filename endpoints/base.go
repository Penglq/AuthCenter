package endpoints

import (
	"github.com/Penglq/AuthCenter/conf"
	"github.com/gin-gonic/gin"
)

type endpoint func(*BaseContext)

type BaseContext struct {
	*gin.Context
}

func (b *BaseContext) JSONSucc(data interface{})  {
	b.JSON(200, conf.ResMap.GetSuccData(data))
}

func (b *BaseContext) JSONErr(code string)  {
	b.JSON(200, conf.ResMap.GetErrData(code))
}

func BaseMiddleware(e endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e(&BaseContext{ctx})
	}
}
