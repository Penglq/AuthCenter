package middlewares

import (
	"github.com/Penglq/AuthCenter/conf"
	"github.com/Penglq/AuthCenter/libs"
	"github.com/gin-gonic/gin"
)

func NewJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(conf.XAuthToken)

		jwts := libs.NewJwt()
		jwts.TokenString = token
		err := jwts.ParseToken()
		if err != nil {
			// todo 错误
			return
		}
		ctx.Set(CasbinSub, jwts.Subject)
		ctx.Next()
	}
}
