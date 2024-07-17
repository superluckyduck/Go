package main

import (
	"fmt"
	"strings"
)
func makeSuffix(suffix string) func (string) string{
	return func (name string) string {
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return name
	}
}
func main(){
	f2 :=makeSuffix(".jpg")
	fmt.Println("文件处理后=",f2("winter"))
}