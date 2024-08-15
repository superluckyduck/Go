package main
import (
	"fmt"
	"math"
	"sort"
)
type Sort struct {
	nums []int
}
//基数排序
func (s Sort)RadixNums() []int{
	maxVal := 0
	for _,num := range s.nums{
		if num > maxVal {
			maxVal = num
		}
	}
	for pos := 0; maxVal/int(math.Pow10(pos)) > 0; pos++ {
		s.nums = s.counterNums(pos)
	}
	return s.nums
}
func (s Sort)counterNums(pos int) []int{
	n := len(s.nums)
	out := make([]int,n)
	count := make([]int,10)
	for i := 0; i < n; i++ {
		digit := (s.nums[i] / int(math.Pow10(pos))) % 10
		count[digit]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		digit := (s.nums[i] / int(math.Pow10(pos))) % 10
		out[count[digit]-1] = s.nums[i]
		count[digit]--
	}
	return out
}

//归并排序
func (s Sort)MergeSort()[]int{
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
	return merge(leftarr.MergeSort(),rightarr.MergeSort())
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
func (s Sort)Select() []int{
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
//桶排序
func (s Sort)Bucket(b int)[]int{
	arr := make([]int,len(s.nums))
	copy(arr,s.nums)
	n := len(arr)
	max := arr[0]
	min := arr[0]
	for i := 0; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	bucket := int(max-min)/int(b)+1
	buckets := make([][]int,bucket)
	for _,num := range arr {
		index := int((num - min) / int(b))
		buckets[index] = append(buckets[index], num)
	}
	sorter := []int{}
	for _,val := range buckets{
		if len(val) > 0 {
			sort.Ints(val)
			sorter = append(sorter, val...)
		}
	}
	return sorter
}
//计数排序
func (s Sort)Count()[]int{
	arr := make([]int,len(s.nums))
	copy(arr,s.nums)
	maxVal := arr[0]
	n := len(arr)
	for i := 0; i < n; i++ {
		if maxVal < arr[i] {
			maxVal = arr[i]
		}
	}
	Len := maxVal+1
	sorter := make([]int, Len)
	for i := 0; i < n; i++ {
		j := arr[i]
		sorter[j] += 1
	}
	out := make([]int,n)
	index := 0
	for i := 0; i < Len; i++ {
		for sorter[i]>0 {
			out[index] = i
			index++
			sorter[i]--
		}
	}
	return out
}
func (s Sort)Heap()[]int{
	arr := make([]int,len(s.nums))
	copy(arr,s.nums)
	arrLen := len(arr)
	for i := arrLen/2-1; i >-1; i-- {
		heapify(arr,arrLen,i)
	}
	for i := arrLen-1; i > 0; i-- {
		arr[0],arr[i] = arr[i],arr[0]
		heapify(arr,i,0)
	}
	return arr
}
func heapify(arr []int,arrLen int,i int) []int {
	for true{
		left := 2*i +1
		right := 2*i +2
		max := i
		if left<arrLen && arr[left]>arr[max] {
			max = left
		}
		if right<arrLen && arr[right]>arr[max] {
			max = right
		}
		if max == i {
			break
		}
		arr[max],arr[i] = arr[i],arr[max]
		i = max
	}
	return arr
}

func main(){
	nums := []int{1,6,2,3,4}
	s := Sort{nums: nums}
	fmt.Println(s)
	arr := s.RadixNums()
	fmt.Println(arr)
	arr = s.MergeSort()
	fmt.Println(arr)
	arr = s.Select()
	fmt.Println(arr)
	arr = s.Bucket(2)
	fmt.Println(arr)
	arr = s.Count()
	fmt.Println(arr)
	arr = s.Heap()
	fmt.Println(arr)
}