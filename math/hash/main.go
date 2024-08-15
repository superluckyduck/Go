package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 键值对 */
type pair struct {
    key int
    val string
}

/* 基于数组实现的哈希表 */
type arrayHashMap struct {
    buckets []*pair
}

/* 初始化哈希表 */
func newArrayHashMap() *arrayHashMap {
    // 初始化数组，包含 100 个桶
    buckets := make([]*pair, 100)
    return &arrayHashMap{buckets: buckets}
}

/* 哈希函数 */
func (a *arrayHashMap) hashFunc(key int) int {
    index := key % 100
    return index
}

/* 查询操作 */
func (a *arrayHashMap) get(key int) string {
    index := a.hashFunc(key)
    pair := a.buckets[index]
    if pair == nil {
        return "Not Found"
    }
    return pair.val
}

/* 添加操作 */
func (a *arrayHashMap) put(key int, val string) {
    pair := &pair{key: key, val: val}
    index := a.hashFunc(key)
    a.buckets[index] = pair
}

/* 删除操作 */
func (a *arrayHashMap) remove(key int) {
    index := a.hashFunc(key)
    // 置为 nil ，代表删除
    a.buckets[index] = nil
}

/* 获取所有键对 */
func (a *arrayHashMap) pairSet() []*pair {
    var pairs []*pair
    for _, pair := range a.buckets {
        if pair != nil {
            pairs = append(pairs, pair)
        }
    }
    return pairs
}

/* 获取所有键 */
func (a *arrayHashMap) keySet() []int {
    var keys []int
    for _, pair := range a.buckets {
        if pair != nil {
            keys = append(keys, pair.key)
        }
    }
    return keys
}

/* 获取所有值 */
func (a *arrayHashMap) valueSet() []string {
    var values []string
    for _, pair := range a.buckets {
        if pair != nil {
            values = append(values, pair.val)
        }
    }
    return values
}

/* 打印哈希表 */
func (a *arrayHashMap) print() {
    for _, pair := range a.buckets {
        if pair != nil {
            fmt.Println(pair.key, "->", pair.val)
        }
    }
}

/* 链式地址哈希表 */
type hashMapChaining struct {
    size        int      // 键值对数量
    capacity    int      // 哈希表容量
    loadThres   float64  // 触发扩容的负载因子阈值
    extendRatio int      // 扩容倍数
    buckets     [][]pair // 桶数组
}

/* 构造方法 */
func newHashMapChaining() *hashMapChaining {
    buckets := make([][]pair, 4)
    for i := 0; i < 4; i++ {
        buckets[i] = make([]pair, 0)
    }
    return &hashMapChaining{
        size:        0,
        capacity:    4,
        loadThres:   2.0 / 3.0,
        extendRatio: 2,
        buckets:     buckets,
    }
}

/* 哈希函数 */
func (m *hashMapChaining) hashFunc(key int) int {
    return key % m.capacity
}

/* 负载因子 */
func (m *hashMapChaining) loadFactor() float64 {
    return float64(m.size) / float64(m.capacity)
}

/* 查询操作 */
func (m *hashMapChaining) get(key int) string {
    idx := m.hashFunc(key)
    bucket := m.buckets[idx]
    // 遍历桶，若找到 key ，则返回对应 val
    for _, p := range bucket {
        if p.key == key {
            return p.val
        }
    }
    // 若未找到 key ，则返回空字符串
    return ""
}

/* 添加操作 */
func (m *hashMapChaining) put(key int, val string) {
    // 当负载因子超过阈值时，执行扩容
    if m.loadFactor() > m.loadThres {
        m.extend()
    }
    idx := m.hashFunc(key)
    // 遍历桶，若遇到指定 key ，则更新对应 val 并返回
    for i := range m.buckets[idx] {
        if m.buckets[idx][i].key == key {
            m.buckets[idx][i].val = val
            return
        }
    }
    // 若无该 key ，则将键值对添加至尾部
    p := pair{
        key: key,
        val: val,
    }
    m.buckets[idx] = append(m.buckets[idx], p)
    m.size += 1
}

/* 删除操作 */
func (m *hashMapChaining) remove(key int) {
    idx := m.hashFunc(key)
    // 遍历桶，从中删除键值对
    for i, p := range m.buckets[idx] {
        if p.key == key {
            // 切片删除
            m.buckets[idx] = append(m.buckets[idx][:i], m.buckets[idx][i+1:]...)
            m.size -= 1
            break
        }
    }
}

/* 扩容哈希表 */
func (m *hashMapChaining) extend() {
    // 暂存原哈希表
    tmpBuckets := make([][]pair, len(m.buckets))
    for i := 0; i < len(m.buckets); i++ {
        tmpBuckets[i] = make([]pair, len(m.buckets[i]))
        copy(tmpBuckets[i], m.buckets[i])
    }
    // 初始化扩容后的新哈希表
    m.capacity *= m.extendRatio
    m.buckets = make([][]pair, m.capacity)
    for i := 0; i < m.capacity; i++ {
        m.buckets[i] = make([]pair, 0)
    }
    m.size = 0
    // 将键值对从原哈希表搬运至新哈希表
    for _, bucket := range tmpBuckets {
        for _, p := range bucket {
            m.put(p.key, p.val)
        }
    }
}

/* 打印哈希表 */
func (m *hashMapChaining) print() {
    var builder strings.Builder

    for _, bucket := range m.buckets {
        builder.WriteString("[")
        for _, p := range bucket {
            builder.WriteString(strconv.Itoa(p.key) + " -> " + p.val + " ")
        }
        builder.WriteString("]")
        fmt.Println(builder.String())
        builder.Reset()
    }
}

func main(){
	/* 初始化哈希表 */
	hmap := make(map[int]string)

	/* 添加操作 */
	// 在哈希表中添加键值对 (key, value)
	hmap[12836] = "小哈"
	hmap[15937] = "小啰"
	hmap[16750] = "小算"
	hmap[13276] = "小法"
	hmap[10583] = "小鸭"

	/* 查询操作 */
	// 向哈希表中输入键 key ，得到值 value
	name := hmap[15937]
	fmt.Println(name)

	/* 删除操作 */
	// 在哈希表中删除键值对 (key, value)
	delete(hmap, 10583)

	/* 遍历哈希表 */
	// 遍历键值对 key->value
	for key, value := range hmap {
		fmt.Println(key, "->", value)
	}
	// 单独遍历键 key
	for key := range hmap {
		fmt.Println(key)
	}
	// 单独遍历值 value
	for _, value := range hmap {
		fmt.Println(value)
	}
}