package main

import (
    "log"
    "nursing-home/backend/internal/router"
    "nursing-home/backend/internal/repository"
    
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func main() {
    // 数据库连接
    db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/nursing_home?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
    if err != nil {
        log.Fatalf("数据库连接失败: %v", err)
    }

    // 初始化数据库连接
    repository.SetDB(db)

    // 创建 Gin 实例
    r := gin.Default()

    // CORS 配置
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })

    // 初始化路由
    router.InitRouter(r)

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("服务器启动失败: %v", err)
    }
} 