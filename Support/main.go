package main

import (
    "log"
    "net/http"
    
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    // 初始化路由
    initRouter(r)
    
    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("服务器启动失败: %v", err)
    }
} 