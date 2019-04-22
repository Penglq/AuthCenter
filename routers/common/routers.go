package common

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func (r *Router) SetMiddlewares(routerMiddleware ...gin.HandlerFunc) *Router {
	for _,middleware := range routerMiddleware {
		r.Use(middleware)
	}
	return r
}

func (r *Router) SetRouter(setRouters ...func( *Router)) *Router {
	for _,setRouter := range setRouters {
		setRouter(r)
	}
	return r
}
