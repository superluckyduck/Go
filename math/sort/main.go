package main

import (
	"fmt"
	"math"
	"sort"
)
// 创建结构体，封装成类
type Sorter struct {
	nums []int
}

// 下面是类方法
// 值传递
// 冒泡排序
func BubbleSort(s Sorter) []int {
	newNums := make([]int, len(s.nums))
	copy(newNums, s.nums)
	n := len(s.nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if newNums[j] > newNums[j+1] {
				newNums[j], newNums[j+1] = newNums[j+1], newNums[j]
			}
		}
	}
	return newNums
}

//选择排序
func SelectSort(s Sorter) []int {
	newNums := make([]int, len(s.nums))
	copy(newNums, s.nums)
	n := len(s.nums)
	for i := 0; i < n; i++ {
		index := i
		for j := i; j < n; j++ {
			if newNums[j] < newNums[index] {
				index = j
			}
		}
		newNums[index], newNums[i] = newNums[i], newNums[index]
	}
	return newNums
}

//插入排序
func InsertSort(s Sorter) []int {
	newNums := make([]int, len(s.nums))
	copy(newNums, s.nums)
	n := len(s.nums)
	for i := 0; i < n; i++ {
		index := i - 1
		num := newNums[i]
		for {
			if index >= 0 && newNums[index] > num {
				newNums[index+1] = newNums[index]
				index--
			} else {
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

//归并排序
func MergeSort(s Sorter) []int {
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	le := Sorter{nums: left}
	right := arr[middle:]
	ri := Sorter{nums: right}
	return Merge(MergeSort(le), MergeSort(ri))
}

func Merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

//堆排序
func HeapSort(s Sorter) []int {
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	arrLen := len(arr)
	//从下往上，建立大顶堆，每个节点都比子节点大，如果不建立，后面的排序不成立（不理解可以举例说明）
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
	for i := arrLen - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}
//因为是从上往下比，所以要先建立大顶堆，才能确保第一个是最大的
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
//桶排序

func BucketSort(s Sorter, bucketSize int) []int {
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	if len(arr) == 0 {
		return arr
	}
	// 找到数组中的最大和最小值
	minValue, maxValue := arr[0], arr[0]
	for _, num := range arr {
		if num < minValue {
			minValue = num
		}
		if num > maxValue {
			maxValue = num
		}
	}
	// 创建桶
	bucketCount := int((maxValue-minValue)/int(bucketSize)) + 1
	buckets := make([][]int, bucketCount)
	// 将数组元素放入对应的桶
	for _, num := range arr {
		index := int((num - minValue) / int(bucketSize))
		buckets[index] = append(buckets[index], num)
	}
	// 对每个桶进行排序并合并
	sortedArr := []int{}
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sort.Ints(bucket) // 使用 Go 的内置排序
			sortedArr = append(sortedArr, bucket...)
		}
	}

	return sortedArr
}
//计数排序
func CountingSort(s Sorter) []int {
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	if len(arr) == 0 {
		return arr
	}
	// 找到数组中的最大值和最小值
	maxValue := arr[0]
	minValue := arr[0]
	for _, num := range arr {
		if num > maxValue {
			maxValue = num
		}
		if num < minValue {
			minValue = num
		}
	}
	// 创建计数数组
	rangeOfElements := maxValue - minValue + 1
	count := make([]int, rangeOfElements)
	// 计数每个元素出现的次数
	for _, num := range arr {
		count[num-minValue]++
	}
	// 修改计数数组，使其存储每个元素的最终位置
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	// 创建输出数组
	output := make([]int, len(arr))
	// 构建输出数组
	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]-minValue]-1] = arr[i]
		count[arr[i]-minValue]--
	}
	return output
}
//基数排序
func RadixSort(s Sorter) []int {
	arr := make([]int, len(s.nums))
	copy(arr, s.nums)
	maxVal := arr[0]
	for _, num := range arr {
		if num > maxVal {
			maxVal = num
		}
	}
	for pos := 0; maxVal/int(math.Pow10(pos)) > 0; pos++ {
		arr = counterSort(arr, pos)
	}
	return arr
}
func counterSort(arr []int, pos int) []int {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)
	for i := 0; i < n; i++ {
		digit := (arr[i] / int(math.Pow10(pos))) % 10
		count[digit]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / int(math.Pow10(pos))) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}
	return output
}




// 指针参数，引用传递
// 冒泡排序
func (s *Sorter) BubbleNums() {
	n := len(s.nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if s.nums[j] > s.nums[j+1] {
				s.nums[j+1], s.nums[j] = s.nums[j], s.nums[j+1]
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
		s.nums[index], s.nums[i] = s.nums[i], s.nums[index]
	}
}

//插入排序
func (s Sorter) InsertNums() {
	n := len(s.nums)
	for i := 0; i < n; i++ {
		index := i - 1
		num := s.nums[i]
		for {
			if index >= 0 && s.nums[index] > num {
				s.nums[index+1] = s.nums[index]
				index--
			} else {
				break
			}
		}
		s.nums[index+1] = num
	}
}

//快速排序
func (s Sorter) QuickNums() {
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

	//归并排序
func (s Sorter)MergeNums() []int {
	length := len(s.nums)
	if length < 2 {
		return s.nums
	}
	middle := length / 2
	left := s.nums[0:middle]
	le := Sorter{nums: left}
	right := s.nums[middle:]
	ri := Sorter{nums: right}
	return Merge(le.MergeNums(), ri.MergeNums())
}

// func Merge(left []int, right []int) []int {
// 	var result []int
// 	for len(left) != 0 && len(right) != 0 {
// 		if left[0] <= right[0] {
// 			result = append(result, left[0])
// 			left = left[1:]
// 		} else {
// 			result = append(result, right[0])
// 			right = right[1:]
// 		}
// 	}
// 	for len(left) != 0 {
// 		result = append(result, left[0])
// 		left = left[1:]
// 	}
// 	for len(right) != 0 {
// 		result = append(result, right[0])
// 		right = right[1:]
// 	}
// 	return result
// }


func main() {
	nums := []int{1, 4, 5, 2, 3}
	sorter := Sorter{nums: nums}
	fmt.Println("值传递")
	newNums := BubbleSort(sorter)
	fmt.Println("冒泡", newNums)
	newNums = SelectSort(sorter)
	fmt.Println("选择", newNums)
	newNums = InsertSort(sorter)
	fmt.Println("插入", newNums)
	newNums = QuickSort(sorter)
	fmt.Println("快速", newNums)
	newNums = MergeSort(sorter)
	fmt.Println("归并", newNums)
	newNums = HeapSort(sorter)
	fmt.Println("堆排序", newNums)
	newNums = BucketSort(sorter,10)
	fmt.Println("桶排序", newNums)
	newNums = CountingSort(sorter)
	fmt.Println("计数", newNums)
	newNums = RadixSort(sorter)
	fmt.Println("基数", newNums)

	fmt.Println("引用传递")
	fmt.Println(sorter)
	sorter.BubbleNums()
	fmt.Println("冒泡", sorter.nums)

	nums = []int{1, 4, 5, 2, 3}
	sorter = Sorter{nums: nums}
	fmt.Println(sorter)
	sorter.SelectNums()
	fmt.Println("选择", sorter.nums)

	nums = []int{1, 4, 5, 2, 3}
	sorter = Sorter{nums: nums}
	fmt.Println(sorter)
	sorter.InsertNums()
	fmt.Println("插入", sorter.nums)

	nums = []int{1, 4, 5, 2, 3}
	sorter = Sorter{nums: nums}
	fmt.Println(sorter)
	sorter.QuickNums()
	fmt.Println("快速", sorter.nums)

	nums = []int{1, 4, 5, 2, 3}
	sorter = Sorter{nums: nums}
	fmt.Println(sorter)
	sorter.MergeNums()
	fmt.Println("快速", sorter.nums)
}
