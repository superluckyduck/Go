//变量
package main
import "fmt"
//全局变量声明
var n9,n10,n11 = 9,10,11
var(
	n12 = 12
	n13 = 13
	n14 =14
)
func main() {
	//变量声明的三种方式
	var a int
	a = 1
	fmt.Println(a)

	var b = 2
	fmt.Println(b)

	c :=3
	fmt.Println(c)
//多变量声明
	var n1,n2,n3 int
	fmt.Println(n1,n2,n3)

	var n4,name,n5 = 4,"name",5
	fmt.Println(n4,name,n5)

	n6,id,n7 := 6,"id",7
	fmt.Println(n6,id,n7)

	fmt.Println(n9,n10,n11)
	fmt.Println(n12,n13,n14)
//变量不断变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println(i)
}
