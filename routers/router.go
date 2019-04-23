package routers

import (
	"github.com/Penglq/AuthCenter/routers/common"
	"github.com/Penglq/AuthCenter/routers/middlewares"
	"github.com/Penglq/AuthCenter/routers/v1"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func NewRouter() common.Router {
	r := common.Router{gin.New()}

	e, err := casbin.NewSyncedEnforcerSafe(`conf/rbac_domain_model.conf`, `conf/rbac_domain_policy.csv`)
	if err != nil {
		// todo 错误
		panic(err)
	}

	r.SetMiddlewares(
		middlewares.NewApiLog(),
		middlewares.NewJwt(),
		middlewares.NewCasbin(e),
	).SetRouter(
		v1.Admin,
	)
	return r
}
