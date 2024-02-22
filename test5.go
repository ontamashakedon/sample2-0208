package main

import (
	"fmt"
)

func main(){
	// int 符号付　unit 符号なし→unit8 = byte
	// 浮動小数点　float64
	// 文字列 string "~"   " hello, \"world\""  "


	// var 変数名　データ型

	var m int
	fmt.Println(m)
	m = 3
	fmt.Println(m)

	var x = 3
	fmt.Printf("\n%v %T\n",x,x)

	y := 3
	fmt.Printf("%T\n",y)

	a,b,c := 1,2,"hello"
	fmt.Println(a,b,c)

	var (
		d = 1
		e = "hello"
	)
	fmt.Println(d,e)
}