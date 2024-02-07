package main

import (
	"fmt"
	"sync"
)

// StrucMutx композиция мапы с мьютексом
type StrucMutx struct {
	mu sync.Mutex
	v  map[int]int
}

// Set - метод для безопасной записи в map
func (sm *StrucMutx) Set(key int, value int) {
	sm.mu.Lock()
	sm.v[key] = value
	sm.mu.Unlock()
}

// Get - метод для безопасного чтения из map
func (sm *StrucMutx) Get(key int) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.v[key]
	return val, ok
}

func main() {
	// 1 вариант - спользование StrucMutx
	sm := &StrucMutx{v: make(map[int]int)}
	sm.Set(1, 2)
	fmt.Println(sm.Get(1))

	// 2 вариант - Использование sync.Map
	var syncMap sync.Map
	syncMap.Store(1, 2)
	value, ok := syncMap.Load(1)
	fmt.Println(value, ok)
}
