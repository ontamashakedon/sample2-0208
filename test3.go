package main

import (
	"fmt"
)

func main() {
	x := "Hello"
	fmt.Printf("%s\n",x)		//c言語と同じ
	pi := 3.14159265
	fmt.Printf("Pi = %f\n",pi)
	fmt.Printf("/%7.2f/,%T\n",pi,pi)	//幅7文字、小数点以下2桁
	fmt.Printf("\n")
	fmt.Printf("Pi = %v\n",pi)	//規定の型
	fmt.Println(x,pi)			//出力の間に空白、終端に改行
	fmt.Println(x,"\n",pi)	
}