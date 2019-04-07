package routers

import (
	"github.com/Penglq/AuthCenter/routers/middlewares"
	"github.com/Penglq/AuthCenter/routers/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() Router {
	r := Router{gin.New()}

	r.setMiddlewares().setRouter()
	return r
}

type Router struct {
	*gin.Engine
}

func (r *Router) setMiddlewares() *Router {
	r.Use(middlewares.NewApiLog())
	r.Use(middlewares.NewAuth())
	return r
}

func (r *Router) setRouter() *Router {
	v1.Admin(r)
	return r
}
