package router

import (
    "github.com/gin-gonic/gin"
    "your-project/handler"
)

func InitRouter(r *gin.Engine) {
    // 养老院相关路由
    api := r.Group("/api")
    {
        // 养老院列表
        api.GET("/nursing-homes", handler.GetNursingHomes)
        // 养老院详情
        api.GET("/nursing-homes/:id", handler.GetNursingHomeDetail)
        // 添加养老院
        api.POST("/nursing-homes", handler.CreateNursingHome)
        // 更新养老院信息
        api.PUT("/nursing-homes/:id", handler.UpdateNursingHome)
        // 删除养老院
        api.DELETE("/nursing-homes/:id", handler.DeleteNursingHome)
    }
} 