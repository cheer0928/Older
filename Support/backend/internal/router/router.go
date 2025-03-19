package router

import (
	"github.com/gin-gonic/gin"
	"nursing-home/backend/internal/handler"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/nursing-homes", handler.GetNursingHomes)
		api.GET("/nursing-homes/:id", handler.GetNursingHomeDetail)
		api.POST("/nursing-homes", handler.CreateNursingHome)
		api.PUT("/nursing-homes/:id", handler.UpdateNursingHome)
		api.DELETE("/nursing-homes/:id", handler.DeleteNursingHome)
	}
} 