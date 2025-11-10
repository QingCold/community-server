package handler

import (
	"net/http"
	"strconv"

	"github.com/QingCold/community-server/internal/db"
	"github.com/QingCold/community-server/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateComplaint(c *gin.Context) {
	var req model.Complaint
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 默认状态
	if req.Status == "" {
		req.Status = "未处理"
	}

	// 保存到数据库
	if err := db.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "投诉已提交", "data": req})
}

func GetComplaint(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var complaint model.Complaint
	if err := db.DB.First(&complaint, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "complaint not found"})
		return
	}

	c.JSON(http.StatusOK, complaint)
}
