package middlewares

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func NewAuth(e *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.EnforceSafe()
	}
}
