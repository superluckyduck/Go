package main
import (
	"fmt"
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

func main(){
	arr := []int{1,6,3,2,4}
	s := Sort{nums: arr}
	fmt.Println(s.nums)
	// s.Bubble()
	// s.Quick()
	// s.Insert()
	new :=s.Merge()
	fmt.Println(new)

}