package middlewares

import (
	"github.com/Penglq/AuthCenter/conf"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)



func NewCasbin(e *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sub string
		if sub = ctx.GetString(CasbinSub); sub == "" {
			ctx.Abort()
			ctx.JSON(200, conf.ResMap.GetErrData(conf.AUTH_NODE))
			return
		}

		if res, err := e.EnforceSafe(ctx.GetString(CasbinSub), ctx.GetHeader(XClientDomain), ctx.GetHeader(XClientObject), ctx.GetHeader(XClientACTION)); err != nil {
			ctx.Abort()
			ctx.JSON(200, conf.ResMap.GetErrData(conf.SYSTERM_NONE))
		} else if !res {
			ctx.Abort()
			ctx.JSON(200, conf.ResMap.GetErrData(conf.AUTH_NODE))
		}
		ctx.Next()
	}
}

type checkAuth struct {
	Subject string
	Object  string
	Domain  string
	Method  string
}
