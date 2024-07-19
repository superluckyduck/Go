package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main(){
	//统计字符串长度
	str := "hello北京"
	fmt.Println("str len= ",len(str))
	//字符串遍历，处理中文
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n",r[i])
	}
	//字符串转int
	n,err := strconv.Atoi("hello")
	if err != nil{
		fmt.Println("转换错误",err)
	}else{
		fmt.Println("转换的结果是",n)
	}
	//int转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T",str,str)
	//字符串转[]byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n",bytes)
	//[]byte转字符串
	str = string([]byte{97,98,99})
	fmt.Printf("str=%v\n",str)
	//10进制转2，8，16
	str = strconv.FormatInt(123,2)
	fmt.Printf("123对应的二进制是%v\n",str)
	str = strconv.FormatInt(123,16)
	fmt.Printf("123对应的十六进制是%v\n",str)
	//查找子串是否在指定字符串中
	b := strings.Contains("seafood","mary")
	fmt.Printf("b=%v\n",b)
	//统计一个字符串有多少个指定的子串
	num := strings.Count("ceheese","e")
	fmt.Printf("num=%v\n",num)
	//不区分大小写比较
	b = strings.EqualFold("abc","ABc")
	fmt.Printf("b=%v\n",b)
	fmt.Println("结果","abc"=="ABc")
	//返回子串在字符串第一次出现的index值，如果没有返回-1
	index := strings.Index("go go go","go")
	fmt.Printf("index=%v\n",index)
	//返回子串在字符串最后一次出现的index值，如果没有返回-1
	index = strings.LastIndex("go golang","go")
	fmt.Printf("index=%v\n",index)
	//将指定子串替换成另一个子串
	str2 = "go go hello"
	str = strings.Replace(str2,"go","北京",-1)
	//-1代表替换所有，换成几就是替换几个
	//按照指定的某字符为分割标识，将一个字符串拆分为字符串数组
	strArr := strings.Split("hello,world,ok",",")
	for i :=0;i<len(strArr);i++{
		fmt.Printf("str[%v]=%v\n",i,strArr[i])
	}
	fmt.Printf("strArr=%v\n",strArr)
	//将字符串的字母进行大小写替换
	str = "goLang hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n",str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n",str)
	//将字符串两边的空格去掉
	str = strings.TrimSpace(" tn a long gopher  ")
	fmt.Printf("str=%q\n",str)
	//将字符串两边的指定字符去掉
	str = strings.Trim("! hello! ","!")
	fmt.Printf("str=%q\n",str)
	//将字符串左边的指定字符去掉
	str = strings.TrimLeft("! hello! ","!")
	fmt.Printf("str=%q\n",str)
	//将字符串右边的指定字符去掉
	str = strings.TrimRight("! hello! ","!")
	fmt.Printf("str=%q\n",str)
	//判断字符串是否以指定字符串开头或者结束
	b = strings.HasPrefix("ftp://192.168.10.1","hsp")
	fmt.Printf("b=%v\n",b)
	b = strings.HasSuffix("ftp://192.168.10.1 hsp","hsp")
	fmt.Printf("b=%v\n",b)
}