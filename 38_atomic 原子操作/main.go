package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作 ，针对 并发时 ，修改 共享变量（不用 锁 的另一种方法）

var (
    x int64
    wg sync.WaitGroup
)

func dogoroutine(){
    defer wg.Done()
    // x++
    atomic.AddInt64(&x,1)   //传 变量 的地址 过去
    fmt.Println( atomic.LoadInt64(&x) )
}

func main(){
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go dogoroutine()
    }
    wg.Wait()
    fmt.Println("X =", x)
}