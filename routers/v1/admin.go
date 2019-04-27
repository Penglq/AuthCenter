package v1

import (
	"github.com/Penglq/AuthCenter/endpoints/v1/admin"
	"github.com/gin-gonic/gin"
)

func Admin(r *gin.Engine)  {
	g := r.Group(`/api`)
	// 登录
	g.POST("/login", admin.Login)
}
