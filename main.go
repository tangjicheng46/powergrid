package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

// ConvertTOMLtoJSON 将TOML文件转换为JSON文件
func ConvertTOMLtoJSON(tomlFileName, jsonFileName string) error {
	// 打开TOML配置文件
	tomlFile, err := os.Open(tomlFileName)
	if err != nil {
		return fmt.Errorf("无法打开TOML配置文件: %v", err)
	}
	defer tomlFile.Close()

	// 创建一个空的interface{}来存储TOML解析后的数据
	var tomlData interface{}

	// 解析TOML文件
	if _, err := toml.DecodeReader(tomlFile, &tomlData); err != nil {
		return fmt.Errorf("解析TOML文件时发生错误: %v", err)
	}

	// 创建JSON文件
	jsonFile, err := os.Create(jsonFileName)
	if err != nil {
		return fmt.Errorf("无法创建JSON文件: %v", err)
	}
	defer jsonFile.Close()

	// 将TOML数据编码为JSON格式并写入JSON文件
	jsonEncoder := json.NewEncoder(jsonFile)
	if err := jsonEncoder.Encode(tomlData); err != nil {
		return fmt.Errorf("写入JSON文件时发生错误: %v", err)
	}

	fmt.Printf("成功将 %s 转换为 %s\n", tomlFileName, jsonFileName)
	return nil
}

// ConvertJSONtoTOML 将JSON文件转换为TOML文件
func ConvertJSONtoTOML(jsonFileName, tomlFileName string) error {
	// 打开JSON文件
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		return fmt.Errorf("无法打开JSON文件: %v", err)
	}
	defer jsonFile.Close()

	// 创建一个 map 用于存储 JSON 解析后的数据
	var jsonData map[string]interface{}

	// 解码 JSON 文件
	jsonDecoder := json.NewDecoder(jsonFile)
	if err := jsonDecoder.Decode(&jsonData); err != nil {
		return fmt.Errorf("解码JSON文件时发生错误: %v", err)
	}

	// 创建 TOML 文件
	tomlFile, err := os.Create(tomlFileName)
	if err != nil {
		return fmt.Errorf("无法创建TOML文件: %v", err)
	}
	defer tomlFile.Close()

	// 将 JSON 数据编码为 TOML 格式并写入 TOML 文件
	if err := toml.NewEncoder(tomlFile).Encode(jsonData); err != nil {
		return fmt.Errorf("写入TOML文件时发生错误: %v", err)
	}

	fmt.Printf("成功将 %s 转换为 %s\n", jsonFileName, tomlFileName)
	return nil
}

func main() {
	jsonFileName := "config1.json"
	tomlFileName := "config2.toml"

	if err := ConvertJSONtoTOML(jsonFileName, tomlFileName); err != nil {
		fmt.Printf("转换出错: %v\n", err)
	}
}

//func main() {
//	tomlFileName := "config1.toml"
//	jsonFileName := "config1.json"
//
//	if err := ConvertTOMLtoJSON(tomlFileName, jsonFileName); err != nil {
//		fmt.Printf("转换出错: %v\n", err)
//	}
//}
