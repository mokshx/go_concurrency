package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestDataRaceConditions(t *testing.T) {
	// race condition - when 2 threads try to write to or read same variable
	//
	var state int32
	var mu sync.RWMutex
	// go x
	// go y
	fmt.Println(state)
	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			mu.Unlock()
		}(i)
	}
}
