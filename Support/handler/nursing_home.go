package handler

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "your-project/model"
    "your-project/service"
)

func GetNursingHomes(c *gin.Context) {
    homes, err := service.GetNursingHomes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, homes)
}

func GetNursingHomeDetail(c *gin.Context) {
    id := c.Param("id")
    home, err := service.GetNursingHomeByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "养老院不存在"})
        return
    }
    c.JSON(http.StatusOK, home)
}

func CreateNursingHome(c *gin.Context) {
    var home model.NursingHome
    if err := c.ShouldBindJSON(&home); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "参数无效"})
        return
    }
    
    if err := service.CreateNursingHome(&home); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, home)
}

func UpdateNursingHome(c *gin.Context) {
    id := c.Param("id")
    var home model.NursingHome
    
    if err := c.ShouldBindJSON(&home); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "参数无效"})
        return
    }
    
    if err := service.UpdateNursingHome(id, &home); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, home)
}

func DeleteNursingHome(c *gin.Context) {
    id := c.Param("id")
    
    if err := service.DeleteNursingHome(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 其他处理函数... 