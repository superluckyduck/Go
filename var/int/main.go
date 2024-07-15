//整数
package main
import (
	"fmt"
	"unsafe"
)
func main(){
	var i int = 1
	fmt.Println(i)
	
	var j int8 = 127
	fmt.Println(j)

	var k uint16 =255
	fmt.Println(k)

	var a int = 8900
	fmt.Println(a)
	var b uint = 1
	var c byte = 255
	fmt.Println(b,c)

	var n1 = 100
	fmt.Printf(" n1的类型%T\n n1占用的字节数是%d\n",n1,unsafe.Sizeof(n1))
	var age byte = 90
	fmt.Println(age)
}