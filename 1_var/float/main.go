//小数
package main
import "fmt"
func main(){
	var price float32 = 89.12
	fmt.Println(price)

	var n1 float32 = -0.00089
	var n2 float64 = -780234.09
	fmt.Println(n1,n2)

	var n3 float32 = -123.0000901
	var n4 float64 = -123.0000901
	fmt.Println(n3,n4)

	var n5 = 1.1
	fmt.Printf("n5的数据类型是%T\n",n5)

	n6 := 5.12
	n7 := .123
	fmt.Println(n6,n7)

	n8 := 5.1234e2
	n9 := 5.1234E2
	n10 := 5.1234e-2
	fmt.Println(n8,n9,n10)
}