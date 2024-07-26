package main
import (
	"fmt"
	"math"
)
type Sort struct {
	nums []int
}
//基数排序
func RadixNums(s Sort) []int{
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	maxVal := 0
	for _,num := range arr{
		if num > maxVal {
			maxVal = num
		}
	}
	for pos := 0; maxVal/int(math.Pow10(pos)) > 0; pos++ {
		arr = counterNums(arr, pos)
	}
	return arr
}
func counterNums(arr []int,pos int) []int{
	n := len(arr)
	out := make([]int,n)
	count := make([]int,10)
	for i := 0; i < n; i++ {
		digit := (arr[i] / int(math.Pow10(pos))) % 10
		count[digit]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / int(math.Pow10(pos))) % 10
		out[count[digit]-1] = arr[i]
		count[digit]--
	}
	return out
}

//归并排序
func MergeSort(s Sort)[]int{
	arr := make([]int,len(s.nums))
	copy(arr,s.nums)
	n :=len(arr)
	if n < 2{
		return arr
	}
	mid := n/2
	left := arr[0:mid]
	leftarr :=Sort{nums: left}
	right :=arr[mid:n]
	rightarr :=Sort{nums: right}
	return merge(MergeSort(leftarr),MergeSort(rightarr))
}
func merge(left []int,right[]int) []int{
	var out []int
	for len(left) !=0 && len(right) !=0{
		if left[0] <= right[0] {
			out = append(out,left[0])
			left = left[1:]
		}else{
			out = append(out,right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		out = append(out, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		out = append(out, right[0])
		right = right[1:]
	}
	return out
}
//选择排序
func Select(s Sort) []int{
	arr := make([]int,len(s.nums))
	copy(arr,s.nums)
	n := len(arr)
	for i := 0; i < n; i++ {
		index :=i
		for j := i; j < n; j++ {
			if arr[j]<arr[index]{
				index = j
			}
		}
		arr[index],arr[i] = arr[i],arr[index]
	}
	return arr
}

func main(){
	nums := []int{1,5,2,3,4}
	s := Sort{nums: nums}
	fmt.Println(s)
	arr := RadixNums(s)
	fmt.Println(arr)
	arr = MergeSort(s)
	fmt.Println(arr)
	arr = Select(s)
	fmt.Println(arr)
}