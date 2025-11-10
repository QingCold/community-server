package handler

import (
	"net/http"

	"github.com/QingCold/community-server/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateComplaint(c *gin.Context) {
	var req model.Complaint
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: 调用 service 层保存数据
	c.JSON(http.StatusOK, gin.H{"message": "投诉已提交", "data": req})
}

func GetComplaint(c *gin.Context) {
	id := c.Param("id")
	// TODO: 查询数据库
	c.JSON(http.StatusOK, gin.H{"id": id, "content": "示例投诉"})
}
