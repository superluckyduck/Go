//随机生成 1-100 的一个数，直到生成了 99 这个数
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)+1
		fmt.Println(n)
		count++
		if (n==99){
			break
		}
	}
	fmt.Println(count)
}