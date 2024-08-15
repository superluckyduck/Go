package main

import (
	"container/list"
	"fmt"
)
func main(){
	/* 初始化队列 */
	// 在 Go 中，将 list 作为队列来使用
	queue := list.New()

	/* 元素入队 */
	queue.PushBack(1)
	queue.PushBack(3)
	queue.PushBack(2)
	queue.PushBack(5)
	queue.PushBack(4)

	/* 访问队首元素 */
	peek := queue.Front()
	fmt.Println(peek.Value)

	/* 元素出队 */
	pop := queue.Front()
	fmt.Println(pop.Value)
	queue.Remove(pop)
	fmt.Println(queue)
	/* 获取队列的长度 */
	size := queue.Len()
	fmt.Println(size)

	/* 判断队列是否为空 */
	isEmpty := queue.Len() == 0
	fmt.Println(isEmpty)
}