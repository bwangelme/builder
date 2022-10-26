package router

import (
	"builder/builder/views/recordview"
	"builder/builder/views/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(middlewares ...gin.HandlerFunc) http.Handler {
	r := gin.New()
	r.Use(middlewares...)
	r.Use(middleware.HandlerHttpErr)

	recordview.InitRouter(r)
	return r
}
