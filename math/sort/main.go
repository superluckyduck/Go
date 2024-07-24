package main
import "fmt"
//创建结构体，封装成类
type Sorter struct {
	nums []int
}
//下面是类方法
//值传递
	//冒泡排序
func BubbleSort(s Sorter) []int{
	newNums := make([]int, len(s.nums))
    copy(newNums, s.nums)
	n := len(s.nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if newNums[j]>newNums[j+1] {
				newNums[j],newNums[j+1] = newNums[j+1],newNums[j]
			}
		}
	}
	return newNums
}
	//选择排序
func SelectSort(s Sorter) []int{
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
		newNums[index],newNums[i] = newNums[i],newNums[index]
	}
	return newNums
}
	//插入排序
func InsertSort(s Sorter) []int{
	newNums := make([]int, len(s.nums))
    copy(newNums, s.nums)
	n :=len(s.nums)
	for i := 0; i < n; i++ {
		index :=i-1
		num := newNums[i]
		for{
			if index>=0 && newNums[index]>num{
				newNums[index+1] = newNums[index]
				index--
			}else{
				break
			}
		}
		newNums[index+1] = num
	}
	return newNums
}
	//快速排序
func QuickSort(s Sorter) []int {
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

//指针参数，引用传递
	//冒泡排序
func (s *Sorter) BubbleNums() {
	n := len(s.nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if s.nums[j] > s.nums[j+1] {
				s.nums[j+1],s.nums[j]= s.nums[j],s.nums[j+1]
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
		s.nums[index],s.nums[i] = s.nums[i],s.nums[index]
	}
}
	//插入排序
func (s Sorter) InsertNums(){
	n :=len(s.nums)
	for i := 0; i < n; i++ {
		index :=i-1
		num := s.nums[i]
		for{
			if index>=0 && s.nums[index]>num{
				s.nums[index+1] = s.nums[index]
				index--
			}else{
				break
			}
		}
		s.nums[index+1] = num
	}
}
	//快速排序
	func (s Sorter)QuickNums() {
		_quickSort(s.nums, 0, len(s.nums)-1)
	}
	
	// func _quickSort(arr []int, left, right int) []int {
	// 	if left < right {
	// 		partitionIndex := partition(arr, left, right)
	// 		_quickSort(arr, left, partitionIndex-1)
	// 		_quickSort(arr, partitionIndex+1, right)
	// 	}
	// 	return arr
	// }
	
	// func partition(arr []int, left, right int) int {
	// 	pivot := left
	// 	index := pivot + 1
	// 	for i := index; i <= right; i++ {
	// 		if arr[i] < arr[pivot] {
	// 			arr[i], arr[index] = arr[index], arr[i]
	// 			index += 1
	// 		}
	// 	}
	// 	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	// 	return index - 1
	// }





func main(){
	nums := []int{1, 4, 5, 2, 3}
	sorter := Sorter{nums: nums}
	fmt.Println("值传递")
	newNums := BubbleSort(sorter)
	fmt.Println("冒泡",newNums)
	newNums = SelectSort(sorter)
	fmt.Println("选择",newNums)
	newNums = InsertSort(sorter)
	fmt.Println("插入",newNums)
	newNums = QuickSort(sorter)
	fmt.Println("快速",newNums)


	fmt.Println("引用传递")
	fmt.Println(sorter)
	sorter.BubbleNums()
	fmt.Println("冒泡",sorter.nums)

	nums = []int{1, 4, 5, 2, 3}
	sorter = Sorter{nums: nums}
	fmt.Println(sorter)
	sorter.SelectNums()
	fmt.Println("选择",sorter.nums)

	nums = []int{1, 4, 5, 2, 3}
	sorter = Sorter{nums: nums}
	fmt.Println(sorter)
	sorter.InsertNums()
	fmt.Println("插入",sorter.nums)

	nums = []int{1, 4, 5, 2, 3}
	sorter = Sorter{nums: nums}
	fmt.Println(sorter)
	sorter.QuickNums()
	fmt.Println("快速",sorter.nums)
}