package main

import (
	"sync"
	"testing"
)

/**
5. 性能比较，vs call
Benchmark 基准测试？ command？
*/
var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	m.Unlock()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}