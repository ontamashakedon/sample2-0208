package main

import (
	"fmt"
)

func main() {

	//リテラル

	x := 1_2_3	//123
	fmt.Printf("%v\n",x)
	x = 0b1100	//2進数12
	x = 0o123	//8進数83
	x = 0xBEEF	//16進数48879

	y := 1.23e-3
	y = 0x1.fp3	//16進数浮動小数点 15.5
				//(1 + 15/16)*2^3  ,f=15,
	fmt.Printf("%v\n",y)

	z := 123i	//虚数
	fmt.Println(z)

	//unicode


	a := 'A'
	b := '\101'	//8bit8進数
	b = '\x41'	//8bit16進数
	b = '\u672c'	//16bit 16進数
	b = '\U0001F604'//32bit 16進数

	fmt.Println(a,b)

	
}