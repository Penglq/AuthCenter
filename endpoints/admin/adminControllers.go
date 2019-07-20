package admin

import (
	"github.com/Penglq/AuthCenter/conf"
	"github.com/Penglq/AuthCenter/endpoints"
	"github.com/Penglq/AuthCenter/libs"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(ctx *endpoints.BaseContext) {
	jwt := libs.NewJwt()
	jwt.Subject = "123456"
	jwt.ExpiresAt = time.Now().Unix() + 3600
	token, err := jwt.GenerateToken()
	if err != nil {
		ctx.JSONErr(conf.SYSTERM_ERROR)
		return
	}
	ctx.JSONSucc(gin.H{
		"jwt": token,
	})
}
func Test(ctx *endpoints.BaseContext) {
	ctx.JSONSucc(gin.H{
		"subject": ctx.Get(conf.CasbinSub),
	})
}
