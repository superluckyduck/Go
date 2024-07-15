/*
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
*/

//100以内的数求和，当和大于20时的当前数
package main
import "fmt"
func main(){
	sum :=0
	for i :=0;i<=100;i++{
		sum +=i
		if sum>20{
			fmt.Println(sum,i)
			break
		}
	}
}
