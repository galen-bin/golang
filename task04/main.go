package main

import (
	postsmanage "blog/posts"
	"fmt"
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
	Username string `gorm:"type:varchar(100);not null;UNIQUE"`
	Password string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);"`
}

type posts struct {
	Title   string `gorm:"type:varchar(255);not null"`
	Content string
	UserId  int32 `gorm:"not null"`
	gorm.Model
}

type comments struct {
	Content string
	UserId  int32 `gorm:"not null"`
	PostId  int32 `gorm:"not null"`
}

// 定义接收数据的结构体
type Logins struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"user" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

var database *gorm.DB

var dsn string = "root:root@tcp(127.0.0.1:3306)/qa?charset=utf8mb4&parseTime=True&loc=Local"

type CustomClaims struct {
	// 可根据需要自行添加字段
	ID                   int64  `json:"user_id"`
	user                 string `json:"user"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

var jwtKey = []byte("12345678")

func main() {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	//database.AutoMigrate(&users{}, &posts{}, &comments{})
	database = db
	postsmanage.DB = db

	//uslist := []posts{{Title: "书籍01", UserId: 1}, {Title: "书籍01", UserId: 2}, {Title: "书籍01", UserId: 3}}
	//uslist := []comments{{Content: "评论01", PostId: 1, UserId: 1}, {Content: "评论02", PostId: 2, UserId: 2}, {Content: "评论03", PostId: 3, UserId: 3}}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	api := r.Group("/api")
	{

		api.GET("/users", getUsers) // 实际路径：/api/users
		api.POST("/reg", Register)
	}

	prv := r.Group("/prv")
	prv.Use(auth)
	{
		prv.POST("/login", Login)                          // 实际路径：/prv/Login
		prv.POST("/tests", tests)                          // 实际路径：/prv/tests
		prv.POST("/create_post", postsmanage.Create_posts) // 创建文章
	}

	//r.Any("/auth", auth)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

func tests(c *gin.Context) {
	uid := c.Keys["userID"]
	c.String(http.StatusOK, "test %v", uid)
}

func getUsers(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	c.String(http.StatusOK, "getUsers %s", name)
}

func Register(c *gin.Context) {
	var user Logins
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pssword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	var reg users
	reg.Password = string(hashedPassword)
	reg.Username = user.User
	if err := database.Create(&reg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user Logins
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser users
	if err := database.Where("username = ?", user.User).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Pssword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	claims := CustomClaims{
		int64(storedUser.Id),
		storedUser.Username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间
			Issuer:    "somebody",                                         // 签发人
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	tokensr, err := tokens.SignedString(jwtKey)
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"code": "-1", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "0", "token": tokensr})

}

func auth(c *gin.Context) {
	// JWT 身份验证中间件
	fmt.Println("JWT 身份验证中间件")
	// 从请求头中获取 JWT
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "-1", "msg": "Missing Authorization header01"})
		c.Abort()
		return
	}

	// 解析并验证 JWT
	token, err := ParseJWT(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "-1", "msg": err.Error(), "token": tokenString})
		c.Abort()
		return
	}

	// 从令牌中提取声明
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "-1", "msg": err.Error(), "token": token.user})
		c.Abort()
		return
	}
	c.Set("userID", token.ID)
	c.Set("userName", token.user)
	c.Next()
}

// 解析并验证给定的 JWT
func ParseJWT(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
