/*
package main
import "fmt"
func printPyramid1(level int){
	//完整
	for i := 1; i <= level; i++ {
		//fmt.Printf("*")
		for k := level-i; k > 0; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func printPyramid2(level int){
	//空心
	for i := 1; i <= level; i++ {
		//fmt.Printf("*")
		for k := level-i; k > 0; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j==1||j==2*i-1||i==level{
				fmt.Print("*")
				//fmt.Print(" ")
			}else{
				//fmt.Print(" ")
				fmt.Print(" ")
			}
			
		}
		fmt.Println()
	}
}
func main(){
	level := 0
	fmt.Println("请输入完整金字塔层数")
	fmt.Scanln(&level)
	printPyramid1(level)
	fmt.Println("请输入空心金字塔层数")
	fmt.Scanln(&level)
	printPyramid2(level)
}
*/
//九九乘法表层数
package main
import "fmt"
func jjcfb(num int){
	for i := 1; i < num+1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ",j,i,i*j)
		}
		fmt.Println()
	}
}
func main(){
	num := 0
	fmt.Println("请输入九九乘法表层数")
	fmt.Scanln(&num)
	jjcfb(num)
}
