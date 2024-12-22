package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// fmtパッケージは、内部的にString() か Error() を呼び出すためeのままだと無限ループになる
	// String() と Error() どちらも実装されている場合はErrorが優先される
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// 平方根の計算を行います
// 値xに対して、z^2が最もxに近いzを求めます
func sqrt(x float64) (float64, error) {
	if x < 0 {
		// キャスト
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2 // 初期値を適切に設定
	for {
		// 平均をとることでより正確に近づける
		newZ := z - (z*z-x)/(2*z)
		if z == newZ { // 収束したらループを抜ける
			break
		}
		z = newZ
	}
	return z, nil
}

func main() {
	r, err := sqrt(-25)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
