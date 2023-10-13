package main

import (
	"fmt"
)

func main() {

	/*--------创建map start ------*/
	//“创建一个映射，键的类型是string，值的类型是int”
	//dict1 := make(map[string]int)

	// 创建一个映射，键和值的类型都是string
	// 使用两个键值对初始化映射
	//dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	/*--------创建map end ------*/

	/*--------创建map start ------*/
	clours := map[string]string{}
	clours["Red"] = "£da1337"

	value, exists := clours["Red"]

	if exists {
		fmt.Println(value)
	}

	value2 := clours["Red"]

	if value2 != "" {
		fmt.Println(value2)
	}

	/*--------创建map end ------*/

	/*--------map 迭代 start ------*/
	fmt.Println("--------clours2--------")
	clours2 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	removeColor(clours2, "Coral")
	for key, value := range clours2 {
		fmt.Println(key, value)
	}
	/*--------map 迭代 end ------*/
}

// removeColor将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
