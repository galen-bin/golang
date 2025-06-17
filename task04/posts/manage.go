package postsmanage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type posts struct {
	Title   string `gorm:"type:varchar(255);not null"`
	Content string
	UserId  int64 `gorm:"not null"`
	gorm.Model
}

type users struct {
	Id       int64  `gorm:"not null;PRIMARY_KEY;AUTO_INCREMENT;UNIQUE_INDEX"`
	Username string `gorm:"type:varchar(100);not null;UNIQUE"`
	Password string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);"`
}

type create_post struct {
	Title   string `form:"title" json:"title" uri:"title" xml:"title" binding:"required"`
	Content string `form:"content" json:"content" uri:"content" xml:"content" binding:"required"`
	UserId  int64
	Id      int `form:"id" json:"id" uri:"id" xml:"id"`
}

var DB *gorm.DB

func Create_posts(c *gin.Context) {
	uid := c.Keys["userID"]

	var pos create_post
	var create_pos posts

	if err := c.ShouldBindJSON(&pos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	create_pos.Title = pos.Title
	create_pos.Content = pos.Content
	create_pos.UserId = uid.(int64)

	err := DB.Create(&create_pos).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "error": "Failed to create user"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "error": "ok"})

}

func Getpost(c *gin.Context) {
	id := c.DefaultQuery("Id", "")
	var post posts
	DB.Where("id", id).First(&post)
	c.JSON(http.StatusOK, gin.H{"code": 0, "error": "ok", "data": &post})

}

func Save_post(c *gin.Context) {
	var pos create_post
	var posinfo posts

	if err := c.ShouldBindJSON(&pos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "error": err.Error()})
		return
	}

	DB.Where("id", pos.Id).First(&posinfo)
	uid := c.Keys["userID"]
	if posinfo.UserId != uid {
		c.JSON(http.StatusForbidden, gin.H{"code": -1, "error": "修改失敗無權限"})
		return
	}

	posinfo.Title = pos.Title
	posinfo.Content = pos.Content
	err := DB.Save(&posinfo).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "error": "修改失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "error": "ok", "data": "修改成功"})
}

func Del_post(c *gin.Context) {

	var posinfo posts

	id := c.DefaultQuery("id", "")

	DB.Where("id", id).First(&posinfo)
	uid := c.Keys["userID"]
	if posinfo.UserId != uid {
		c.JSON(http.StatusForbidden, gin.H{"code": -1, "error": "修改失敗無權限"})
		return
	}

	err := DB.Delete(&posinfo).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "error": "修改失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "error": "ok", "data": "删除成功"})
}
