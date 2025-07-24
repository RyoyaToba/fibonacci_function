package main

// フィボナッチ関数
func get_fibonacci(n int) (ans int) {

	// n = 0であれば0を返す
	if n == 0 {
		return 0
	}

	// n = 1であれば1を返す
	if n == 1 {
		return 1
	}

	// 再帰処理として実装を考える
	return get_fibonacci(n-1) + get_fibonacci(n-2)
}

// エントリーポイント(コマンドでの実行想定はないが、一応)
func main() {

}
