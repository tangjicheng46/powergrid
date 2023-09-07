package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // 导入 SQLite 驱动程序
)

func main() {
	// 打开或创建 SQLite 数据库文件
	db, err := sql.Open("sqlite3", "mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	insertData := `
	INSERT INTO users (username, email) VALUES (?, ?);
	`
	_, err = db.Exec(insertData, "user1", "user1@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	query := `
	SELECT id, username, email FROM users;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var id int
		var username string
		var email string
		err := rows.Scan(&id, &username, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", id, username, email)
	}

	// 检查是否有错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
