package main
import "fmt"
//创建结构体，封装成类
type Sorter struct {
	nums []int
}
//下面是类方法
//值传递
	//冒泡排序
func BubbleNum(s Sorter) []int{
	newNums := make([]int, len(s.nums))
    copy(newNums, s.nums)
	n := len(s.nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if newNums[j]>newNums[j+1] {
				temp :=newNums[j]
				newNums[j] = newNums[j+1]
				newNums[j+1] = temp
			}
		}
	}
	return newNums
}
	//选择排序
func SelectNum(s Sorter) []int{
	newNums := make([]int, len(s.nums))
    copy(newNums, s.nums)
	n :=len(s.nums)
	for i := 0; i < n; i++ {
		index :=i
		for j := i; j < n; j++ {
			if newNums[j]<newNums[index] {
				index = j
			}
		}
		temp :=newNums[i]
		newNums[i] = newNums[index]
		newNums[index] = temp
	}
	return newNums
}
//指针参数，引用传递
	//冒泡排序
func (s *Sorter) BubbleNums() {
	n := len(s.nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if s.nums[j] > s.nums[j+1] {
				temp := s.nums[j]
				s.nums[j] = s.nums[j+1]
				s.nums[j+1] = temp
			}
		}
	}
}
	//选择排序
func (s *Sorter) SelectNums() {
	n := len(s.nums)
	for i := 0; i < n; i++ {
		index := i
		for j := i; j < n; j++ {
			if s.nums[j] < s.nums[index] {
				index = j
			}
		}
		temp := s.nums[i]
		s.nums[i] = s.nums[index]
		s.nums[index] = temp
	}
}

func main(){
	nums := []int{1, 4, 5, 2, 3}
	sorter := Sorter{nums: nums}
	newNums := BubbleNum(sorter)
	fmt.Println(newNums)
	newNums = SelectNum(sorter)
	fmt.Println(newNums)
	sorter.BubbleNums()
	fmt.Println(sorter.nums)
	sorter = Sorter{nums: nums}
	sorter.SelectNums()
	fmt.Println(sorter.nums)
	
}