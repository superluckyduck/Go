package main
import "fmt"
func main(){
	//使用方式
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
	j :=1
	for j <=10{
		fmt.Println("hello")
		j++
	}
	k :=1
	for{
		if k<=10{
			fmt.Println("ok",k)
		}else{
			break
		}
		k++
	}
	//遍历
	//传统方式
	var str string = "hello,world!北京"
	str2 := []rune(str)//把str转成[]rune
	for n :=0;n < len(str2);n++{
		fmt.Printf("%c\n",str[n])
	}
	fmt.Println()
	//for-range
	str = "abc~ok"
	for index,val :=range str{
		fmt.Printf("%d,%c\n",index,val)
	}
	//while语句的实现
	var a int = 1
	for{
		if a>10{
			break
		}
		fmt.Println("hello,world",a)
		a++
	}
	//do-while语句的实现
	var b int = 1
	for{
		fmt.Println("hello,world",b)
		b++
		if b>10{
			break
		}
	}
	//break,continue
	lable1:
	for c := 0; c < 4; c++ {
		if c==2{
			continue
		}
		fmt.Println(c)
		for j:=0;j<10;j++{
			if j ==2{
				break lable1
			}
			fmt.Println("j=",j)
		}
	}
	//goto
	d :=30
	fmt.Println("ok1")
	if d > 20{
		goto lable2
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	lable2:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}