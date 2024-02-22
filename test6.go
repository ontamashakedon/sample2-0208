package main

import(
	"fmt"
)

func main(){
	// const 定数名 = 定数値
	// const 定数名 データ型 = 定数値

	const M = 10	//定数　型は決まっていない

	var i int = M
	var x float64 = M

	fmt.Printf("%v %T,%v %T\n",i,x)


	const n int = 10
	var a int = n
	var b float64 = float64(n)

	fmt.Printf("%T, %T\n",a,b)

	//iota

	const (
		q = iota + 10
		w
		e
		r
	)

	fmt.Println(q,w,e,r)

}