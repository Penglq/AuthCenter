package v1

import (
	"github.com/Penglq/AuthCenter/controllers/v1/admin"
	"github.com/Penglq/AuthCenter/routers/common"
)

func Admin(r *common.Router)  {
	g := r.Group(`/api`)
	// 登录
	g.POST("/login", admin.Admin)
}
