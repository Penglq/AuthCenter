package middlewares

import (
	"github.com/Penglq/AuthCenter/libs"
	"github.com/gin-gonic/gin"
)

func NewAuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt := libs.Jwt{
			TokenString: ctx.GetHeader("auth-token"),
		}
		err := jwt.ParseToken()
		if err != nil {
			//ctx.JSON()
		}

		ctx.Set(libs.Subject, jwt.Subject)
		ctx.Set(libs.SubjectPath, jwt.SubPath)
		ctx.Set(libs.Domain, jwt.Domain)
	}
}
