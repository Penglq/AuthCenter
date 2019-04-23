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
<<<<<<< HEAD

	e, err := casbin.NewSyncedEnforcerSafe(`conf/rbac_domain_model.conf`, `conf/rbac_domain_policy.csv`)
	if err != nil {
		// todo 错误
		panic(err)
	}

	r.SetMiddlewares(
		middlewares.NewApiLog(),
		middlewares.NewJwt(),
		middlewares.NewCasbin(e),
=======
	r.SetMiddlewares(
		middlewares.NewApiLog(),
		middlewares.NewAuthToken(),
		// middlewares.NewAuth(),
>>>>>>> 48dce5218152edb99f46b30b21d2bf416a220677
	).SetRouter(
		v1.Admin,
	)
	return r
}
