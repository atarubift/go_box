package main

import (
	"fmt"
	"math"
)

// const は定数
const s string = "constant"

func main() {
	fmt.Println(s)

	// var文が書けるところであればどこでも定義可能
	const n = 500000000

	// 任意精度で算術を実行
	const d = 3e20 / n
	fmt.Println(d)

	// 数値の定数は明示的にキャストするなどして
	// 型を決めるまで型をもたない
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
