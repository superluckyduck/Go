
package main
import (
	"fmt"
	"project/go/var/model"
)
func main() {
	// 指针
	i := 10
	fmt.Println("i address", &i)
	var ptr *int = &i
	fmt.Printf("%v\n", ptr)
	fmt.Printf("address%v\n", &ptr)
	fmt.Printf("ptr指向的值%v", *ptr)
	//标识符
	fmt.Println(model.HeroName)
}
