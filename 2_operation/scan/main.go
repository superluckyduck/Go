
//Scanln输入
package main
import "fmt"
func main(){
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)
	fmt.Printf("名字是%v \n年龄是%v \n薪水是%v \n是否通过考试%v \n",name,age,sal,isPass)
//Scanf输入
	var name1 string
	var age1 byte
	var sal1 float32
	var isPass1 bool
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name1,&age1,&sal1,&isPass1)
	fmt.Printf("名字是%v \n年龄是%v \n薪水是%v \n是否通过考试%v \n",name1,age1,sal1,isPass1)
}
