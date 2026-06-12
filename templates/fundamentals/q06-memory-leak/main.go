package main

import (
	"fmt"
	"runtime"
)

func main() {
	var fns []func()
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024*1024) // 1MB each
	}

	for i, v := range data {
		_ = v
		fns = append(fns, func() {
			fmt.Println(i)
		})
	}

	_ = fns
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %d MB\n", m.Alloc/1024/1024)
}
