package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Person struct {
		Name    string
		Age     int
		Address struct {
			Street      string
			City        string
			Coordinates struct {
				Latitude  float64
				Longitude float64
			}
		}
	}
	Server struct {
		Hostname string
		Port     int
	}
	Logging struct {
		Level string
		File  string
	}
}

func main() {
	// 打开TOML配置文件
	file, err := os.Open("config1.toml")
	if err != nil {
		fmt.Println("无法打开配置文件:", err)
		return
	}
	defer file.Close()

	// 创建一个Config结构体实例来存储解析后的配置数据
	var config Config

	// 使用toml.Decode来解析TOML文件并将数据填充到config结构体中
	if _, err := toml.DecodeReader(file, &config); err != nil {
		fmt.Println("解析TOML文件时发生错误:", err)
		return
	}

	// 打印解析后的配置数据
	fmt.Printf("Person Name: %s\n", config.Person.Name)
	fmt.Printf("Person Age: %d\n", config.Person.Age)
	fmt.Printf("Street: %s, City: %s\n", config.Person.Address.Street, config.Person.Address.City)
	fmt.Printf("Latitude: %f, Longitude: %f\n", config.Person.Address.Coordinates.Latitude, config.Person.Address.Coordinates.Longitude)
	fmt.Printf("Server Hostname: %s, Port: %d\n", config.Server.Hostname, config.Server.Port)
	fmt.Printf("Logging Level: %s, File: %s\n", config.Logging.Level, config.Logging.File)
}
