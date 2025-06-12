package lesson02

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name   string
	Post   []Posts `gorm:"foreignKey:UserID;references:ID"`
	Number uint    `gorm:"default:0;not null"`
	Status uint    `gorm:"default:1;not null;size:2"`
	gorm.Model
}

type Posts struct {
	Title   string
	UserID  uint
	Comment []Comment `gorm:"foreignKey:PostId;references:ID"`
	gorm.Model
}

type Comment struct {
	PostId  uint
	Content string
	gorm.Model
}

var DB *gorm.DB

func (P Posts) AfterCreate(db *gorm.DB) (err error) {

	err = db.Model(&User{}).Where("id=?", P.UserID).Update("number", gorm.Expr("number+1")).Error
	fmt.Println(err)
	return err
}

func (P Posts) AfterDelete(db *gorm.DB) (err error) {

	fmt.Println(P, P.UserID)
	var us User
	db.Model(&User{}).Where("id=", P.UserID).First(&us)
	qu := db.Model(&User{}).Where("id=?", P.UserID)

	if us.Number-1 > 0 {
		err = qu.Update("number", gorm.Expr("number-1")).Error
	} else {

		err = qu.Updates(map[string]interface{}{"status": 0, "number": us.Number - 1}).Error
	}

	fmt.Println(err, "删除")
	return err
}

func Run() {
	dsn := "root:root@tcp(127.0.0.1:3306)/qa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(&Comment{}, &Posts{}, &User{})
	var po Posts
	po.ID = 4
	db.Debug().Delete(&po)
	/*ps := []Posts{
		{Title: "老吴的书", UserID: 1},
		{Title: "老吴的书", UserID: 1},
	}
	db.Create(&ps)

	users := []User{
		{Name: "sdkdk"},
		{Name: "老adfas刘"},
	}
	db.Create(&users)*/
	/*	ps := []Posts{
			{Title: "老吴的书", Author: 1},
			{Title: "老刘的书", Author: 2},
		}

		db.Create(&users)
		db.Create(&ps)

		comt := []Comment{
			{Content: "评论内容01", PostId: 1},
			{Content: "评论内容01", PostId: 1},
			{Content: "评论内容01", PostId: 1},
			{Content: "评论内容01", PostId: 1},
			{Content: "评论内容02", PostId: 2},
			{Content: "评论内容02", PostId: 2},
			{Content: "评论内容02", PostId: 2},
			{Content: "评论内容02", PostId: 2},
		}
		db.Create(&comt)*/

	//db.Create(&User{Name: "法外狂徒张三"})
	//var us User
	//var ps []Posts
	//	db.Model(&User{}).Preload("posts").Find(&us)
	//db.Model(&User{}).Association("posts").Find(&us)
	//db.Preload("posts").Find(&us)
	//获取评论最多文章
	/*var com Comment
	db.Model(&Comment{}).Select("*,COUNT(*) count").Group("post_id").Order("count desc").First(&com)
	//err = db.Select(&bs, "select * from book where price>50")
	type result struct {
		Name string
		gorm.Model
		Title string
	}

	var res result

	db.Raw("SELECT a.*,b.title FROM users a LEFT JOIN posts b on a.id=b.user_id where b.id=?", com.PostId).Scan(&res)

	fmt.Println(res)*/

}
