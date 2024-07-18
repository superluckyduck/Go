/*
//统计3个班，每班五个同学，求平均分-for
package main
import "fmt"
func main(){
	var class int = 3
	var stu int = 5
	total := 0.0
	for i :=1;i<=class;i++{
		sum := 0.0
		for j :=1;j<=stu;j++{
			fmt.Printf("请输入第%d班，第%d号同学的成绩",i,j)
			score := 0.0
			fmt.Scanln(&score)
			sum += score
		}
		ave := sum/5
		fmt.Printf("第%v班的总分是%.2f,平均分是%.2f\n",i,sum,ave)
		total += sum
	}
	fmt.Printf("总分%v 平均分%v",total,total/(5*3))
}

//打印金字塔-for
package main
import "fmt"
func main(){
	//完整
	level :=10
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


//打印九九乘法表-for
package main
import "fmt"
func main(){
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ",j,i,i*j)
		}
		fmt.Println()
	}
}

//100以内的数求和，当和大于20时的当前数-break
package main
import "fmt"
func main(){
	sum :=0
	for i :=0;i<=100;i++{
		sum +=i
		if sum>20{
			fmt.Println(sum,i)
			break
		}
	}
}

//打印1--100之内的奇数-continue
package main
import "fmt"
func main(){
	for i := 0; i <= 100; i++ {
		if i%2==0{
			continue
		}
		fmt.Print(i," ")
	}
}

//输入个数不确定的整数，并判断读入的正数和负数的个数，输入0时结束程序 break-continue
package main
import "fmt"
func main(){
	z := 0
	f := 0
	for {
		fmt.Println("请输入一个整数")
		num := 0
		fmt.Scanln(&num)
		if num > 0{
			z++
			continue
		}
		if num < 0{
			f++
			continue
		}else{
			break
		}
	}
	fmt.Printf("正数%v 负数%v",z,f)
}

//某人有 100.000 元,每经过一次路口，需要交费,规则如下当现金>50000 时,每次交 5%
//当现金<=50000 时,每次交 1000
//编程计算该人可以经过多少次路口,使用 for break 方式完成
package main
import "fmt"
func main(){
	count := 100000
	sum := 0
	for{
		if count<=1000{
			break
		}
		if count >50000{
			count -= count/20
		}else{
			count -= 1000
		}
		sum++
		fmt.Println(count)
	}
	fmt.Println(sum)
}
*/

// 输入账号密码-break，continue
package main

import "fmt"

func main() {
	var name string
	var password int
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名、密码用空格隔开")
		fmt.Scanln(&name)
		fmt.Scanln(&password)
		//fmt.Scanf("%s %d ",&name,&password)有问题
		if name == "张无忌" && password == 888 {
			fmt.Println("欢迎")
			break
		} else {
			n := 3 - i
			if n > 0 {
				fmt.Printf("请重新输入，还有%d次机会\n", n)
				continue
			} else {
				fmt.Println("机会用完")
				break
			}
		}
	}

}
