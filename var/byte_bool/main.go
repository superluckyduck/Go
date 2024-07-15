//字符和bool
package main
import (
	"fmt"
	"unsafe"
)
func main(){
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println(c1,c2)
	fmt.Printf("c1=%c c2=%c\n",c1,c2)

	//var c3 byte = '北' overflow \溢出
	var c3 int = '北'
	fmt.Printf("c3=%c c3对应的码值=%d",c3,c3)

	var c4 int = 22269
	fmt.Printf("%c\n",c4)
	var n1 = 10 + 'a'
	fmt.Printf("%c\n",n1)

	var b = false
	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(b))
}