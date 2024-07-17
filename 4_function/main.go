package main
import "fmt"
//递归
//斐波那契数
func fbn(n int) int{
	if(n == 1||n == 2){
		return 1
	}else{
		return fbn(n-1)+fbn(n-2)
	}
}
//求2*f(n-1)+1函数
func f(n int)int{
	if n==1{
		return 3
	}else{
		return 2*f(n-1)+1
	}
}
//猴子吃桃
func peach(n int)int{
	if n>10||n<1{
		fmt.Println("输入天数不对")
		return 0
	}else if n==10{
		return 1
	}else{
		return (peach(n+1)+1)*2
	}
}
//init函数
func init(){
	fmt.Println("init()...")
}
func main(){
	res :=fbn(3)
	fmt.Println("res=",res)
	fmt.Println("res=",fbn(4))
	fmt.Println("res=",fbn(6))
	fmt.Println(f(1))
	fmt.Println(f(5))
	fmt.Println("第一天桃子数量是",peach(1))
}