package main

/*

//O_CREATE创建并O_WRONLY输入
import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	path := "D:/Go/src/project/go/file/abc.txt"
	file,err := os.OpenFile(path,os.O_WRONLY | os.O_CREATE,0666)
	if err !=nil{
		fmt.Println("open file err = ",err)
		return
	}
	defer file.Close()
	str := "hello Gardon\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

//os.O_TRUNC替换
import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	path := "D:/Go/src/project/go/file/abc.txt"
	file,err := os.OpenFile(path,os.O_WRONLY | os.O_TRUNC,0666)
	if err !=nil{
		fmt.Println("open file err = ",err)
		return
	}
	defer file.Close()
	str := "你好 Gardon\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

//O_APPEND追加
import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	path := "D:/Go/src/project/go/file/abc.txt"
	file,err := os.OpenFile(path,os.O_WRONLY | os.O_APPEND,0666)
	if err !=nil{
		fmt.Println("open file err = ",err)
		return
	}
	defer file.Close()
	str := "你好 hello Gardon\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
*/
//读写 O_RDWR
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main(){
	path := "D:/Go/src/project/go/file/abc.txt"
	file,err := os.OpenFile(path,os.O_RDWR | os.O_APPEND,0666)
	if err !=nil{
		fmt.Println("open file err = ",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		str,err :=reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
	}
	str := "你好  Gardon\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	//拷贝文件
	path2 := "D:/Go/src/project/go/file/bcd.txt"
	data,err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read file err = ",err)
		return
	}
	err = os.WriteFile(path2,data,0666)
	if err != nil {
		fmt.Println("read file err = ",err)
	}
}