package main

import (
	"fmt"
	"math/cmplx"
)

var (
	typeBool   bool   = false
	typeString string = "hello world"

	typeInt    int    = 1<<63 - 1
	typeUint   uint   = 1<<64 - 1
	typeInt8   int8   = 1<<7 - 1
	typeUint8  uint8  = 1<<8 - 1
	typeInt16  int16  = 1<<15 - 1
	typeUint16 uint16 = 1<<16 - 1
	typeInt32  int32  = 1<<31 - 1
	typeUint32 uint32 = 1<<32 - 1
	typeInt64  int64  = 1<<63 - 1
	typeUint64 uint64 = 1<<64 - 1

	typeFloat32 float32 = 3.4028235e38
	typeFloat64 float64 = 5e-324

	// uint8 のエイリアス
	typeByte byte = 1<<8 - 1

	// int32 のエイリアス
	// unicodeのコードポイントを表すためのもの
	typeRune rune = 1<<31 - 1

	// 虚数
	typeComplex64  complex64  = complex(3.0, 4.0)
	typeComplex128 complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// すべての変数の値を標準出力に表示
	fmt.Println("typeBool:", typeBool)
	fmt.Println("typeString:", typeString)
	fmt.Println("typeInt:", typeInt)
	fmt.Println("typeUint:", typeUint)
	fmt.Println("typeInt8:", typeInt8)
	fmt.Println("typeUint8:", typeUint8)
	fmt.Println("typeInt16:", typeInt16)
	fmt.Println("typeUint16:", typeUint16)
	fmt.Println("typeInt32:", typeInt32)
	fmt.Println("typeUint32:", typeUint32)
	fmt.Println("typeInt64:", typeInt64)
	fmt.Println("typeUint64:", typeUint64)
	fmt.Println("typeFloat32:", typeFloat32)
	fmt.Println("typeFloat64:", typeFloat64)
	fmt.Println("typeByte:", typeByte)
	fmt.Println("typeRune:", typeRune)
	fmt.Println("typeComplex64:", typeComplex64)
	fmt.Println("typeComplex128:", typeComplex128)
}

