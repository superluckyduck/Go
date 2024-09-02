package main
import (
	"fmt"
	"sort"
)
type Sort struct{
	nums []int
}
func (s *Sort)Bubble(){
	n := len(s.nums)
	for i :=0;i<n;i++{
		for j :=0;j<n-i-1;j++{
			if s.nums[j]>s.nums[j+1] {
				s.nums[j],s.nums[j+1] = s.nums[j+1],s.nums[j]
			}
		}
	}
}
func (s *Sort)Quick(){
	n := len(s.nums)
	Quicksort(s.nums,0,n-1)
}
func Quicksort(arr []int,left int,right int){
	if left<right {
		index := Quickly(arr,left,right)
		Quicksort(arr,left,index)
		Quicksort(arr,index+1,right)
	}
}
func Quickly(arr []int,left int,right int)int{
	fixed := left
	index := left+1
	for i :=left;i<=right;i++{
		if arr[i]<arr[fixed] {
			arr[i],arr[index]= arr[index],arr[i]
			index++
		}
	}
	arr[index-1],arr[fixed] = arr[fixed],arr[index-1]
	return index-1
}
func (s Sort)Insert(){
	n := len(s.nums)
	for i:=0;i<n;i++{
		num := s.nums[i]
		j := i-1
		for j>=0 && s.nums[j]>num {
			s.nums[j+1] = s.nums[j]
			j--
		}
		s.nums[j+1] = num
	}
}
func (s Sort)Merge()[]int{
	arr := s.nums
	if len(arr)<2 {
		return arr
	}
	length := len(arr)
	middle := length / 2
	left := arr[0:middle]
	le := Sort{nums: left}
	right := arr[middle:]
	ri := Sort{nums: right}
	return MergeSort(le.Merge(), ri.Merge())

}
func MergeSort(left []int,right []int) []int{
	var res []int
	for len(left)>0 && len(right)>0{
		if left[0]<right[0] {
			res = append(res, left[0])
			left= left[1:]
		}else{
			res = append(res, right[0])
			right= right[1:]
		}
	}
	for len(left)>0{
		res = append(res, left[0])
		left= left[1:]
	}
	for len(right)>0{
		res = append(res, right[0])
		right= right[1:]
	}
	return res
}
func (s Sort)HeapSort() []int {
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	arrLen := len(arr)
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
	for i := arrLen - 1; i >= 0; i-- {
		heapify(arr, 0, arrLen)
		arr[i], arr[0] = arr[0], arr[i]
		arrLen -= 1
	}
	return arr
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest, arrLen)
	}
}
func (s Sort)Bucket(size int){
	maxVal := s.nums[0]
	minVal := s.nums[0]
	for i:=0;i<len(s.nums);i++{
		if s.nums[i] < minVal {
			minVal = s.nums[i]
		}
		if s.nums[i] > maxVal {
			maxVal = s.nums[i]
		}
	}
	count := (maxVal-minVal)/size+1
	buckets := make([][]int, count)
	for _, num := range s.nums {
		index := int((num - minVal) / int(size))
		buckets[index] = append(buckets[index], num)
	}
	sortedArr := []int{}
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sort.Ints(bucket)
			sortedArr = append(sortedArr, bucket...)
		}
	}
	for i:=0;i<len(s.nums);i++{
		s.nums[i] = sortedArr[i]
	}
}

func main(){
	arr := []int{1,6,5,3,2,4,7,9}
	s := Sort{nums: arr}
	fmt.Println(s.nums)
	// s.Bubble()
	// s.Quick()
	// s.Insert()
	// new :=s.Merge()
	// new := s.HeapSort()
	s.Bucket(2)
	fmt.Println(s.nums)

}