package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type users struct {
	Id       int32  `gorm:"not null;PRIMARY_KEY;AUTO_INCREMENT;UNIQUE_INDEX"`
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
}

type posts struct {
	Title   string `gorm:"type:varchar(255);not null"`
	Content sql.NullString
	UserId  int32 `gorm:"not null"`
	gorm.Model
}

type comments struct {
	Content string
	UserId  int32 `gorm:"not null"`
	PostId  int32 `gorm:"not null"`
}

var database *gorm.DB

var dsn string = "root:root@tcp(127.0.0.1:3306)/qa?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	//database.AutoMigrate(&users{}, &posts{}, &comments{})
	database = db
	//uslist := []posts{{Title: "书籍01", UserId: 1}, {Title: "书籍01", UserId: 2}, {Title: "书籍01", UserId: 3}}
	//uslist := []comments{{Content: "评论01", PostId: 1, UserId: 1}, {Content: "评论02", PostId: 2, UserId: 2}, {Content: "评论03", PostId: 3, UserId: 3}}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	api := r.Group("/api")
	{
		api.GET("/users", getUsers) // 实际路径：/api/users
		api.POST("/login", login)   // 实际路径：/api/login
	}

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

func getUsers(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	c.String(http.StatusOK, "getUsers %s", name)
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("userpassword")

	c.JSON(http.StatusOK, gin.H{"status": "200", "name": username, "psd": password})
}

func Register(c *gin.Context) {
	var user users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := database.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser users
	if err := database.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.Id,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	_, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 剩下的逻辑...
}
