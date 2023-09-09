package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// 打开或创建数据库文件
	db, err := sql.Open("sqlite3", "my-database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS my_table (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	insertStatement := "INSERT INTO my_table (name, age) VALUES (?, ?)"
	_, err = db.Exec(insertStatement, "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(insertStatement, "Bob", 25)
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name, age FROM my_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// 修改数据
	updateStatement := "UPDATE my_table SET age = ? WHERE name = ?"
	_, err = db.Exec(updateStatement, 28, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// 删除数据
	deleteStatement := "DELETE FROM my_table WHERE name = ?"
	_, err = db.Exec(deleteStatement, "Bob")
	if err != nil {
		log.Fatal(err)
	}

	// 再次查询数据以查看修改和删除的效果
	rows, err = db.Query("SELECT id, name, age FROM my_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Updated query results:")
	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
