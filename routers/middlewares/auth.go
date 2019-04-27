package middlewares

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

const (
	CasbinKey     = "casbinKey"
	CasbinSub     = "casbinSub"
	XClientDomain = "X-CLIENT-DOMAIN"
	XClientObject = "X-CLIENT-OBJECT"
	XClientACTION = "X-CLIENT-ACTION"
)

func NewCasbin(e *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sub string
		if sub = ctx.GetString(CasbinSub); sub == "" {
			//ctx.JSON()
			return
		}

		if res, err := e.EnforceSafe(ctx.GetString(CasbinSub), ctx.GetHeader(XClientDomain), ctx.GetHeader(XClientObject), ctx.GetHeader(XClientACTION)); err != nil {
			return
		} else if !res {
			return
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
