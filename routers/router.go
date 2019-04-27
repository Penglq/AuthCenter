package routers

import (
	"github.com/Penglq/AuthCenter/routers/middlewares"
	"github.com/Penglq/AuthCenter/routers/v1"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	e, err := casbin.NewSyncedEnforcerSafe(`conf/rbac_domain_model.conf`, `conf/rbac_domain_policy.csv`)
	if err != nil {
		// todo 错误
		panic(err)
	}
	r.Use(
		middlewares.NewApiLog(),
		middlewares.NewJwt(),
		middlewares.NewCasbin(e),
	)

	v1.Admin(r)

	return r
}
