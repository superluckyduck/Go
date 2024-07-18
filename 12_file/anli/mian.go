//拷贝图片/影片/MP3文件案例
package main
import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func CopyFile(dstFileName string,srcFileName string)(writen int64,err error){
	srcFile,err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%vn", err)
	}
	defer srcFile.Close()
	//通过 srcfle ,获取到 Reader
	reader := bufio.NewReader(srcFile)
	//打开 dstFileName
	dstFile, err :=os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%vn", err)
		return
	}
		//通过 dstFile，获取到 Writer
		writer := bufio.NewWriter(dstFile)
		defer dstFile.Close()
	return io.Copy(writer, reader)
}

func main(){
	//将 d:/flower.ipg 文件拷贝到 e:/abc.ipg
	//调用 CopyFile 完成文件拷贝
	srcFile :="D:/Go/src/project/go/file/aaa.png"
	dstFile :="D:/Go/src/project/go/file/bbb.png"
	_ , err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Print("拷贝完成\n")
	}else{
		fmt.Print("拷贝错误 err=%v\n",err)
	}
}