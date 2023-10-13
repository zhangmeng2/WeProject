package main

import "fmt"

func main() {
	var array [4][2]int

	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	array = [4][2]int{1: {20, 31}, 3: {40, 41}}

	array = [4][2]int{1: {0: 20}, 3: {1: 4}}

	var array1 [2][2]int
	var array2 [2][2]int

	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40

	array1 = array2

	for i, v := range array {

		fmt.Println(i, v)
	}

	for i, v := range array1 {

		fmt.Println(i, v)
	}

	//数组很大的时候，如果需要在方法之间传递为了降低内存的开销，可以使用接受数组地址的方式
}
