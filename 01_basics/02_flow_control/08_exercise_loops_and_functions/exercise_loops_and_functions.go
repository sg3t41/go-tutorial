package main

import "fmt"

// 平方根の計算を行います
// 値xに対して、z^2が最もxに近いzを求めます
func sqrt(x float64) float64 {
	z := x / 2 // 初期値を適切に設定
	for {
		// 平均をとることでより正確に近づける
		newZ := z - (z*z-x)/(2*z)
		if z == newZ { // 収束したらループを抜ける
			break
		}
		z = newZ
	}
	return z
}

func main() {
	r := sqrt(25)
	fmt.Println(r)
}

