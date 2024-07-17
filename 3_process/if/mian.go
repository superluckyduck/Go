package main
import "fmt"
func main(){
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	if score == 100{
		fmt.Println("优秀")
	}else if score > 80 && score <= 99{
		fmt.Println("良好")
	}else if score >= 60 && score <=80{
		fmt.Println("中等")
	}else{
		fmt.Println("不及格")
	}
}