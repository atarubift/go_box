package main

import "fmt"

func main() {

	// 基本
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("t is odd")
	}

	// elseはなくてもいい
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 条件の前に文を書いてよい
	// この分で宣言された変数は、いずれの分岐からも見える
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
