package main
import (
	"fmt"
	"math/rand"
)
//随机访问元素
func randomAccess(nums []int) int{
	randomIndex := rand.Intn(len(nums))
	randomNum := nums[randomIndex]
	return randomNum
}
//在数组的索引 index 处插入元素 num
func insert(nums []int,num int,index int){
	for i := len(nums)-1;i>index;i--{
		nums[i]=nums[i-1]
	}
	nums[index]=num
}
//删除索引 index 处的元素
func remove(nums []int,index int){
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}
//遍历数组
func traverse(nums []int) {
    count := 0
    // 通过索引遍历数组
    for i := 0; i < len(nums); i++ {
        count += nums[i]
    }
    count = 0
    // 直接遍历数组元素
    for _, num := range nums {
        count += num
    }
    // 同时遍历数据索引和元素
    for i, num := range nums {
        count += nums[i]
        count += num
    }
}
//在数组中查找指定元素
func find(nums []int, target int) (index int) {
    index = -1
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            index = i
            break
        }
    }
    return
}
//扩展数组长度
func extend(nums []int, enlarge int) []int {
    // 初始化一个扩展长度后的数组
    res := make([]int, len(nums)+enlarge)
    // 将原数组中的所有元素复制到新数组
    for i, num := range nums {
        res[i] = num
    }
    // 返回扩展后的新数组
    return res
}
func main(){
	var arr [5]int
	fmt.Println(arr)
	nums := []int{1,2,3,4,5}
	fmt.Println(nums)
	fmt.Println(randomAccess(nums))
	insert(arr[:],2,3)
	fmt.Println(arr)
	remove(nums,1)
	fmt.Println(nums)
	traverse(nums)
	fmt.Println(find(nums,3))
	fmt.Println(extend(nums,3))
}