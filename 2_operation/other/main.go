package main

import "fmt"

func main() {
	//&返回变量地址，*指针变量
	a := 100
	fmt.Println("a的地址", &a)

	var ptr *int = &a
	fmt.Println("ptr指向的值", *ptr)
	//理解三元运算符，go没有
	var n int
	var i int = 10
	var j int = 12
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println("n =", n)

	//二进制输出
	fmt.Printf("%b \n", i)
	//八进制转十进制输出
	var k int = 011
	fmt.Println(k)
	//十六进制转十进制输出
	var l int = 0x11
	fmt.Println(l)

	//位运算
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)
	fmt.Println(-2 ^ 2)
	b := 1 >> 2
	c := 1 << 2
	fmt.Println(b, c)
}
