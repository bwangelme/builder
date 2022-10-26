package recordview

import "github.com/gin-gonic/gin"

func InitRouter(r gin.IRouter) {
	r.GET("/records", ListRecord)
	r.POST("/record", CreateRecord)
}
