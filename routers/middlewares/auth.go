package middlewares

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func NewCasbin(e *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.EnforceSafe()
	}
}

type checkAuth struct {
	Subject string
	Object  string
	Domain  string
}
