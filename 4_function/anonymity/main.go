package main
import "fmt"
//全局匿名函数
var(
	Fun1 = func (n1 int,n2 int) int {
		return n1*n2
	}
)
func main(){
	//匿名函数
	res1 := func (n1 int,n2 int) int {
		return n1+n2
	}(10,20)
	fmt.Println("res1=",res1)
	//匿名函数赋给变量，通过变量调用函数
	a := func (n1 int,n2 int) int {
		return n1-n2
	}
	res2 := a(10,30)
	fmt.Println("res2=",res2)
	res3 := Fun1(4,9)
	fmt.Println("res3=",res3)
}