package main
import "fmt"
func main(){
	//判断interface
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type){
		case nil:
			fmt.Printf("%T\n",i)
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float64")
		case func(int) float64:
			fmt.Println("func(int)")
		case bool,string:
			fmt.Println("bool或者string")
		default:
			fmt.Println("未知")
	}
	
	b := "nil"
	switch b {
		case "nil":
			fmt.Printf("%v\n",b)
			fallthrough //只能穿透一层
		case "int":
			fmt.Printf("int")
		case "float64":
			fmt.Printf("float64")
		case "func(int) float64":
			fmt.Printf("func(int)")
		case "bool","string":
			fmt.Printf("bool或者string")
		default:
			fmt.Println("未知")
	}
}
