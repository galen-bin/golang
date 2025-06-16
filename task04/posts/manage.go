package postsmanage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type posts struct {
	Title   string `gorm:"type:varchar(255);not null"`
	Content string
	UserId  int32 `gorm:"not null"`
	gorm.Model
}

type users struct {
	Id       int32  `gorm:"not null;PRIMARY_KEY;AUTO_INCREMENT;UNIQUE_INDEX"`
	Username string `gorm:"type:varchar(100);not null;UNIQUE"`
	Password string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);"`
}

type create_post struct {
	Title   string `form:"title" json:"title" uri:"title" xml:"title" binding:"required"`
	Content string `form:"content" json:"content" uri:"content" xml:"content" binding:"required"`
	UserId  int32
}

var DB *gorm.DB

func Create_posts(c *gin.Context) {
	//uid := c.Keys["userID"]
	var pos create_post
	var create_pos posts

	if err := c.ShouldBindJSON(&pos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	create_pos.Title = pos.Title
	create_pos.Content = pos.Content
	create_pos.UserId = c.Keys["userID"].(int32)
	fmt.Println(create_pos)
	fmt.Println(3344)
	create_pos = posts{
		Title:   pos.Title,
		Content: pos.Content,
		UserId:  c.Keys["userID"].(int32),
	}
	fmt.Println(create_pos)
	fmt.Println(1123)
	err := DB.Create(&create_pos).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "error": "Failed to create user"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "error": "ok"})

}
