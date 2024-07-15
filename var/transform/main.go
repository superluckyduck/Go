package main
import (
	"fmt"
	"strconv"
)
func main(){
//基本数据类型转换
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n",i,n1,n2,n3)
	fmt.Printf("i=%T n1=%T n2=%T n3=%T\n",i,n1,n2,n3)
//基本数据类型转string，使用fmt.Sprintf
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var c byte = 'h'
	var str string
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("type=%T str=%q\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("type=%T str=%q\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("type=%T str=%q\n",str,str)

	str = fmt.Sprintf("%c",c)
	fmt.Printf("type=%T str=%q\n",str,str)

	str = strconv.FormatInt(int64(num1),10)
	fmt.Printf("type=%T str=%q\n",str,str)
	//'f'格式 10:表示小数位保留10位 64:表示这个小数是float64
	str = strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("type=%T str=%q\n",str,str)
	str = strconv.FormatBool(b)
	fmt.Printf("type=%T str=%q\n",str,str)
//string转基本数据类型
	var str1 string = "true"
	var b1 bool
	b1,_ = strconv.ParseBool(str1)
	fmt.Printf("type=%T b=%v\n",b1,b1)

	str2 := "123456"
	var n4 int64
	var n5 int
	n4,_ = strconv.ParseInt(str2,10,64)
	n5 = int(n4)
	fmt.Printf("%T %v\n",n4,n4)
	fmt.Printf("%T %v\n",n5,n5)

	str3 := "123.123"
	var f float64
	f,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("%T %v\n",f,f)
}