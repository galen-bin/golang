package main

import "task03/lesson02"

func main() {
	//dsn := "root:root@tcp(127.0.0.1:3306)/qa?charset=utf8mb4&parseTime=True&loc=Local"
	/*db, err := gorm.Open(mysql.Open(dsn))
	lesson01.Dns = db

	if err != nil {
		panic(err)
	}*/

	//lesson01.Add()
	//lesson01.Get()
	//lesson01.Update()
	//lesson01.Delete()
	//result := lesson01.Transaction(1, 2, 10)
	//fmt.Println(result)
	//lesson01.Run()
	lesson02.Run()

}
