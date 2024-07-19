package main
import (
	"fmt"
	"time"
)
func main(){
	now := time.Now()
	fmt.Printf("now=%v type=%T",now,now)
}