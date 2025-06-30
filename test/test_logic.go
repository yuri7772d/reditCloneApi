package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMemUsage(tag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc = %v KB, TotalAlloc = %v KB, Sys = %v KB, NumGC = %v\n",
		tag, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

func main() {

	var count int64 = 0

	printMemUsage("ก่อนสร้าง goroutine")

	start := time.Now()
	duration := time.Second

	for {
		if time.Since(start) > duration {
			break
		}

		go func() {
			time.Sleep(10000)
		}()
		count++
	}

	printMemUsage("หลังสร้าง goroutine")

	fmt.Printf("สร้าง goroutine ได้ %d ตัวใน 1 วินาที\n", count)
}
