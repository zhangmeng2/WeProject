package main

import "fmt"

func main22() {

	var arr [5]int

	var arr2 = [5]int{10, 20, 30, 40, 50}

	//“容量由初始化值的数量决定”
	var arr3 = [...]int{10, 20, 39}

	//初始化位置为 1，2的值
	var arr4 = [5]int{1: 20, 2: 40}

	//指针数组
	arr5 := [5]*int{0: new(int), 1: new(int)}
	*arr5[0] = 10
	*arr5[1] = 20

	//指针赋值
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array1 = array2

	fmt.Println("---------1")
	for _, i := range arr {
		fmt.Println(i)
	}

	fmt.Println("---------2")
	for _, j := range arr2 {
		fmt.Println(j)
	}

	fmt.Println("---------3")
	for _, t := range arr3 {
		fmt.Println(t)
	}

	fmt.Println("---------4")
	for _, t := range arr4 {
		fmt.Println(t)
	}

	fmt.Println("---------5")
	for _, t := range array1 {
		fmt.Println(t)
	}

	fmt.Println("---------6")
	for _, t := range array2 {
		fmt.Println(t)
	}
}
