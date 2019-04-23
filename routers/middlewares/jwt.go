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
		_, err := jwts.ParseToken()
		if err != nil {
			// todo 错误
		}
		ctx.Set("casbin", checkAuth{
			Subject: jwts.Subject,
			Object:  jwts.Object,
			Domain:  jwts.Domain,
		})
	}
}
