package main

import "fmt"

/**
“记住，如果在[]运算符里指定了一个值，那么创建的就是数组而不是切片。只有不指定值的时候，才会创建切片”
*/

func main() {
	//创建一个长度和容量都是5的切片
	//slice := make([]string, 5)

	//创建一个长度为3，容量为5的切片
	//slice2 := make([]int, 3, 5)

	//字符串切片其长度和容量都是5
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	//创建长度和容量都是100个元素的切片
	//slice4 := []string{99:""}

	//创建3个元素的数组
	//array := [3]int{10,20,30}
	//创建长度和容量都是3的切片
	//slice5 := []int{10,20,30}

	for i, v := range slice3 {
		fmt.Println(i, v)
	}

	//创建nil 整型切片
	//var slice []int

	/*-----------空切片 start--------*/
	//slice6 := make([]int,0)
	//slice7 := [] int{}
	/*-----------空切片 end--------*/

	/*----------切片赋值 同数组 start------*/
	slice3[1] = "Block"
	for i, v := range slice3 {
		fmt.Println(i, v)
	}
	/*----------切片赋值 同数组 end------*/

	/*----------创建新切片 长度为2，容量为4 start
	长度，容量计算公示：
	“对底层数组容量是k的切片slice[i:j]来说
	长度: j - i     容量为 3-1=2
	容量: k - i”    容量为 5-1=4
	--------*/

	slice4 := slice3[1:3]
	for i, v := range slice4 {
		fmt.Println(i, v)
	}
	/*----------创建新切片 end--------*/

	/*----------切片容量增加 start------*/
	slice5 := append(slice4, "Pink")
	for i, v := range slice5 {
		fmt.Println(i, v)
	}

	fmt.Println("--------slice3------")
	for i, v := range slice3 {
		fmt.Println(i, v)
	}
	/*----------切片容量增加 end------*/

	/*----------切片容量限制 第三个参数限制容量 start------*/
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice := source[2:3:4]
	fmt.Println("--------slice------")
	for i, v := range slice {
		fmt.Println(i, v)
	}
	/*----------切片容量限制 end------*/

	/*----------两个slice start------------*/
	// 创建两个切片，并分别用两个整数进行初始化
	s1 := []int{1, 2}
	s2 := []int{3, 4}

	// 将两个切片追加在一起，并显示结果
	fmt.Printf("%v\n", append(s1, s2...))
	/*----------两个slice end---------*/

}
