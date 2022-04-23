package main

import "fmt"

func main() {
	// 文字列は + を使って連結可能
	fmt.Println("go" + "lang")

	// 整数と実数
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// 真偽値。ブール演算子使用可能
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
