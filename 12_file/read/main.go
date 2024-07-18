package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main(){
	file,err := os.Open("D:/Go/src/project/go/file/test.txt")
	if err !=nil{
		fmt.Println("open file err = ",err)
	}
	fmt.Printf("file=%v\n",file)
	fmt.Println("第一次文件读取……")
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		str,err :=reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
	}
	
	fmt.Println("文件读取结束……")
	fmt.Println("第二次文件读取……ioutil适用于文件小")
	file2 := "D:/Go/src/project/go/file/test.txt"
	content,err :=os.ReadFile(file2)
	if err != nil {
		fmt.Println("read file err = ",err)
	}
	fmt.Printf("%v",string(content))
}
