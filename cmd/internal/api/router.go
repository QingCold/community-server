package api

import (
	"github.com/QingCold/community-server/cmd/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/complaint", handler.CreateComplaint)
		v1.GET("/complaint/:id", handler.GetComplaint)
	}
	return r
}
