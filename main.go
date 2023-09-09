package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// 定义数据模型
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// 自动迁移数据库表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	newUser := User{Name: "John", Email: "john@example.com"}
	db.Create(&newUser)

	// 查询数据
	var user User
	db.First(&user, 1) // 查询ID为1的用户

	// 更新数据
	db.Model(&user).Update("Name", "New Name")

	// 删除数据
	//db.Delete(&user)

	// 查询所有用户
	var users []User
	db.Find(&users)

	// 打印查询结果
	fmt.Println("All Users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}
}
