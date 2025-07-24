package main

import "fmt"

// フィボナッチ関数
func get_fibonacci(n int) (ans int) {

	// 値を保管しておくためのキャッシュを用意
	cash := make(map[int]int)

	var helper func(int) int

	helper = func(k int) int {
		// キャッシュに値が存在すればその値を返す
		if v, ok := cash[k]; ok {
			return v
		}

		if k == 0 {
			return 0
		}

		if k == 1 {
			return 1
		}
		// キャッシュに値を保存する
		cash[k] = helper(k-1) + helper(k-2)
		return cash[k]
	}

	// 再帰処理として実装を考える
	return helper(n)
}

// エントリーポイント(コマンドでの実行想定はないが、一応)
func main() {
	ans := get_fibonacci(100)
	fmt.Println(ans)
}
