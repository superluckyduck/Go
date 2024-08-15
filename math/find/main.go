package main
import (
	"fmt"
)
//二分查找
func binarySearch(arr []int,target int)int{
	i,j := 0,len(arr)-1
	for  i<=j{
		m := i+(j-i)/2
		if arr[m]<target {
			i = m+1
		}else if arr[m] >target{
			j = m-1
		}else{
			return m
		}
	}
	return -1
}
// 二分查找插入点
func binarySearchInsertion(nums []int, target int) int {
    i, j := 0, len(nums)-1
    for i <= j {
        m := i + (j-i)/2
        if nums[m] < target {
            i = m + 1
        } else if nums[m] > target {
            j = m - 1
        } else{
			 // 首个小于 target 的元素在区间 [i, m-1] 中
            j = m - 1
        }
    }
    // 返回插入点 i
    return i
}
/* 二分查找最左一个 target */
func binarySearchLeftEdge(nums []int, target int) int {
    // 等价于查找 target 的插入点
    i := binarySearchInsertion(nums, target)
    // 未找到 target ，返回 -1
    if i == len(nums) || nums[i] != target {
        return -1
    }
    // 找到 target ，返回索引 i
    return i
}
/* 二分查找最右一个 target */
func binarySearchRightEdge(nums []int, target int) int {
    // 转化为查找最左一个 target + 1
    i := binarySearchInsertion(nums, target+1)
    // j 指向最右一个 target ，i 指向首个大于 target 的元素
    j := i - 1
    // 未找到 target ，返回 -1
    if j == -1 || nums[j] != target {
        return -1
    }
    // 找到 target ，返回索引 j
    return j
}
//哈希优化
/* 方法一：暴力枚举 */
func twoSumBruteForce(nums []int, target int) []int {
    size := len(nums)
    // 两层循环，时间复杂度为 O(n^2)
    for i := 0; i < size-1; i++ {
        for j := i + 1; j < size; j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
/* 方法二：辅助哈希表 */
func twoSumHashTable(nums []int, target int) []int {
    // 辅助哈希表，空间复杂度为 O(n)
    hashTable := map[int]int{}
    // 单层循环，时间复杂度为 O(n)
    for idx, val := range nums {
        if preIdx, ok := hashTable[target-val]; ok {
            return []int{preIdx, idx}
        }
        hashTable[val] = idx
    }
    return nil
}
func main(){
	nums := []int{1,2,4,6,7,8}
	a := binarySearch(nums,7)
	fmt.Println(a)
	c := binarySearchLeftEdge(nums,5)
	fmt.Println(c)
}