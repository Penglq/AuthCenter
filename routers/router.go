package routers

import (
	"github.com/Penglq/AuthCenter/routers/common"
	"github.com/Penglq/AuthCenter/routers/middlewares"
	"github.com/Penglq/AuthCenter/routers/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() common.Router {
	r := common.Router{gin.New()}

	r.SetMiddlewares(
		middlewares.NewApiLog(),
		// middlewares.NewAuth(),
	).SetRouter(
		v1.Admin,
	)
	return r
}
