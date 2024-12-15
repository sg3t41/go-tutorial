package main

import "fmt"

func main() {
	// 長さ4の配列
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// 0番目と1番目
	a := names[0:2]
	// 1番目と2番目
	b := names[1:3]
	fmt.Println(a, b)

	// スライスは配列の参照(厳密には違う)なのでスライスの変更は配列に影響する
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// スライスのキャパを越えた
	test()
}

func test() {
	var array [30]int = [30]int{1, 2, 3}
	fmt.Println("配列: ", array)
	fmt.Println(&array[0])
	fmt.Println("配列のキャパ: ", cap(array))

	var sliceA []int = array[:] //cap 3
	fmt.Println("スライス: ", sliceA)
	fmt.Println(&sliceA[0])
	fmt.Println("スライスのキャパ：", cap(sliceA))

	fmt.Println("************")
	fmt.Println("スライスのキャパを超えない処理")
	sliceA[0] = 1000
	fmt.Println("スライス: ", sliceA)
	fmt.Println(&sliceA[0])
	fmt.Println("スライスのキャパ：", cap(sliceA))
	fmt.Println("配列: ", array)
	fmt.Println(&array[0])
	fmt.Println("配列のキャパ: ", cap(array))

	fmt.Println("************")
	fmt.Println("スライスのキャパを超えた編集をした場合")
	sliceA = append(sliceA, 1)
	fmt.Println("スライス: ", sliceA)
	fmt.Println(&sliceA[0])
	fmt.Println("スライスのキャパ：", cap(sliceA))
	fmt.Println("配列: ", array)
	fmt.Println(&array[0])
	fmt.Println("配列のキャパ: ", cap(array))

	fmt.Println("************")
	fmt.Println("スライスのキャパを1000にして同様の処理をした場合")
	sliceB := make([]int, 10, 10)
	sliceB = append(sliceB, array[:]...)

	sliceB = append(sliceB, 1)
	fmt.Println("スライス: ", sliceB)
	fmt.Println(&sliceB[0])
	fmt.Println("スライスのキャパ：", cap(sliceB))
	fmt.Println("配列: ", array)
	fmt.Println(&array[0])
	fmt.Println("配列のキャパ: ", cap(array))

	//基本的にスライスに対して処理を行うときは新しいスライスを作って返すのが無難そう
}
