package comt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

type comments struct {
	Content string `form:"content" json:"content" uri:"content" xml:"content" binding:"required"`
	UserId  int64
	PostId  int32 `form:"pid" json:"pid" uri:"pid" xml:"pid" binding:"required"`
}

func Create_comt(c *gin.Context) {
	var coms comments
	if err := c.ShouldBindJSON(&coms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := c.Keys["userID"]
	coms.UserId = uid.(int64)

	err := DB.Create(&coms).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok"})
	return

}

func Read_comt(c *gin.Context) {
	var coms []comments
	pid := c.DefaultQuery("pid", "")
	uid := c.Keys["userID"]

	query := DB.Model(&comments{})
	if pid != "" {
		query.Where("post_id", pid)
	} else {
		query.Where("user_id", uid)
	}

	query.Find(&coms)

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": coms})
	return

}
