package v1

import (
	"github.com/Penglq/AuthCenter/endpoints"
	"github.com/Penglq/AuthCenter/endpoints/admin"
	"github.com/gin-gonic/gin"
)

func Admin(r *gin.RouterGroup)  {
	r.GET("/test", endpoints.BaseMiddleware(admin.Test))
}
