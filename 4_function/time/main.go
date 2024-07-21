package main
import (
	"fmt"
	"time"
)
func main(){
	now := time.Now()
	fmt.Printf("now=%v type=%T\n",now,now)
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",now.Month())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())
	fmt.Printf("当前时间 %d-%d-%d %d:%d:%d\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	date :=fmt.Sprintf("当前时间 %d-%d-%d %d:%d:%d\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(date)
}