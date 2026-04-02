package main

import (
	"fmt"
	"sync"
)

// 线程安全的泛型映射
type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

// 创建新的 SafeMap
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// 设置键值对
func (m *SafeMap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[key] = value
}

// 获取值
func (m *SafeMap[K, V]) Get(key K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	value, exists := m.data[key]
	return value, exists
}

// 删除键
func (m *SafeMap[K, V]) Delete(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, key)
}

// 获取所有键
func (m *SafeMap[K, V]) Keys() []K {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	keys := make([]K, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// 使用示例
func main() {
	// 创建字符串到整数的映射
	scores := NewSafeMap[string, int]()
	scores.Set("Alice", 95)
	scores.Set("Bob", 87)

	if score, exists := scores.Get("Alice"); exists {
		fmt.Printf("Alice's score: %d\n", score) // 输出: Alice's score: 95
	}

	fmt.Println("Keys:", scores.Keys()) // 输出: Keys: [Alice Bob]
	
	// 类型推断
	// 类型推断示例
	fmt.Println(Max([]int{1, 2, 3}))      // 编译器推断 T 为 int
	fmt.Println(Max([]float64{1.1, 2.2})) // 编译器推断 T 为 float64

	// 显式指定类型（有时需要）
	var result int = Max[int]([]int{1, 2, 3})
	fmt.Println(result)
}
