package main

import "fmt"

func main() {
	// 変数の宣言
	var a = "initial"
	fmt.Println(a)

	// 一度に複数宣言可能
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// 明示的に初期化していない変数の値は0
	var e int
	fmt.Println(e)

	// := は宣言と初期化を行うための省略記法
	// var f string = "apple" と同義
	f := "apple"
	fmt.Println(f)
}
