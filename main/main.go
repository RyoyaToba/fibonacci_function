package main

import (
	"fmt"
)

// フィボナッチ関数
func get_fibonacci(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// エントリーポイント(コマンドでの実行想定はないが、一応)
func main() {
	ans := get_fibonacci(50)
	fmt.Println(ans)
}
