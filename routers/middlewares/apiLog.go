package middlewares

import (
	"github.com/Penglq/QLog"
	"github.com/gin-gonic/gin"
	"time"
)

func NewApiLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func(t time.Time) {
			QLog.GetLogger().Info(
				"tooks/s", time.Since(t).Seconds(),
			)
		}(time.Now())
		ctx.Next()
	}
}
