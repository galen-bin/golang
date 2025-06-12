package lesson01

import (
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var Dns *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Users struct {
	gorm.Model
	Name         string
	Age          sql.NullInt16
	Email        string  `gorm:"type:varchar(100);nuique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:unique;not null`
	Num          int     `gorm:AUTO_INCREMENT`
	Address      string  `gorm:index:addr`
	IgnoreMe     int     `gorm:"-"`
}

type User struct {
	gorm.Model
	CreditCard CreditCard `gorm:"foreignkey:CardRefer"`
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}

type Students struct {
	ID    int    `gorm:"PRIMARY_KEY;unique;not null;AUTO_INCREMENT"`
	Name  string `gorm:"type:varchar(50);not null;"`
	Age   uint16 `gorm:"size:2;not null;"`
	Grade string `gorm:"type:varchar(100);not null;"`
}

/*
type Product struct {
	gorm.Model
	Code  string
	Price uint
}*/

type Accounts struct {
	gorm.Model
	Balance float32 `gorm:"default:0;not null;type:decimal(10,2)"`
}

type Transactions struct {
	gorm.Model
	From_account_id int64   `gorm:"default:0;not null"`
	To_account_id   int64   `gorm:"default:0;not null"`
	Amount          float32 `gorm:"type:decimal(10,2)"`
}

func Add() {
	/*studen := []Students{
		{Name: "张三", Age: 18, Grade: "一年级"},
		{Name: "李四", Age: 19, Grade: "二年级"},
		{Name: "王麻子", Age: 20, Grade: "三年级"},
	}*/
	//studen01 := Students{Name: "张仁义", Age: 14, Grade: "一年级"}
	ant := []Accounts{{Balance: 100}, {Balance: 100}}
	Dns.Create(&ant)

}

func Get() {
	var st []Students
	Dns.Where("age>?", 18).Find(&st)

	fmt.Println(st)
}

func Update() {
	//Dns.Model(&Product{}).Update("Price", 200).Where("Code", "100")
	//Dns.Table("users").Where("ID in (?)", []int{1, 2}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	Dns.Model(&Students{}).Where("name=?", "张三").Updates(Students{Grade: "四年级"})
}

func Delete() {
	//var p Product
	//Dns.Delete(&p)
	//Dns.Where("age = ?", 20).Delete(&User{})
	//Dns.Where("code = ?", 100).Delete(&Product{})
	Dns.Where("age<?", 15).Delete(&Students{})
}

func Create_table() {

	Dns.AutoMigrate(&Accounts{}, &Transactions{})
}

func Transaction(now_id int64, to_id int64, money float32) error {
	var now, tos Accounts
	result := Dns.Transaction(func(tx *gorm.DB) error {

		// 加锁查询
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Take(&now, now_id).
			Error; err != nil {

			return err
		}

		if now.Balance-money < 0 {

			return errors.New("金额不足")

		}

		err := tx.Model(&now).Where("id", now_id).Update("balance", gorm.Expr("balance - ?", money)).Error
		if err != nil {

			return err
		}

		error02 := tx.Model(&tos).Where("id", to_id).Update("balance", gorm.Expr("balance + ?", money)).Error

		if error02 != nil {

			return error02
		}

		fmt.Println("end")
		return nil
	})
	fmt.Println(result)
	return result

	tx := Dns.Begin()

	// 加锁查询
	if err := tx.Set("gorm:query_option", "FOR UPDATE").
		Take(&now, now_id).
		Error; err != nil {
		fmt.Sprintf("账户: %d 加锁失败.", now_id)
		return err
	}

	err := tx.Model(&now).Where("id", now_id).Update("balance", gorm.Expr("balance - ?", money)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	error02 := tx.Model(&tos).Where("id", to_id).Update("balance", gorm.Expr("balance + ?", money)).Error

	if error02 != nil {
		tx.Rollback()
		return error02
	}
	error03 := tx.Create(&Transactions{From_account_id: now_id, To_account_id: to_id, Amount: money}).Error
	if error03 != nil {
		tx.Rollback()
		return error02
	}

	if now.Balance-money < 0 {
		tx.Rollback()
		fmt.Println(now.Balance-money < 0)
		return errors.New("金额不足")

	}
	fmt.Println("end")
	return tx.Commit().Error

}
