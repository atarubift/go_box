package main

import "fmt"

func main() {

	// 条件1つ(基本)
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 初期化、条件、ループ間処理を記述(古典的)
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 条件を書かないforループはbreakeまたはreturnで抜けるまで繰り返す
	for {
		fmt.Println("loop")
		break
	}

	// continue ループ内の次の繰り返しに進む
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
