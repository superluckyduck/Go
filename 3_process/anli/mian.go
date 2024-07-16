/*
package main
import "fmt"
func main(){
	var name string
	var password int
	for i :=1;i<=3;i++{
		fmt.Println("请输入用户名、密码用空格隔开")
		fmt.Scanln(&name)
		fmt.Scanln(&password)
		//fmt.Scanf("%s %d ",&name,&password)有问题
		if name=="张无忌" && password ==888{
			fmt.Println("欢迎")
			break
		}else{
			n := 3-i
			if n>0 {
				fmt.Printf("请重新输入，还有%d次机会\n",n)
				continue
			}else{
				fmt.Println("机会用完")
				break
			}
		}
	}

}


//统计3个班，每班五个同学，求平均分(未完成)
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

//打印金字塔
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
*/

//打印九九乘法表
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
