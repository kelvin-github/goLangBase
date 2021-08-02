package main

import (
	"fmt"
	"sync"
)

// worker_pool 概念，下面 的写法 就是  工作池 模式
// 单向通道，一般只用在 函数中，只读 或 只写，只是 基于 操作安全 的考虑

// 只写 通道
func f01(ch01 chan<- int){
    defer close(ch01)
    defer wg.Done()
    for i := 1; i <= 10; i++ {
        ch01 <- i
    }
}
// 只读  只写 通道
func f02(ch01 <-chan int, ch02 chan<- int){
    defer close(ch02)
    defer wg.Done()
    for{
        x ,ok:= <-ch01
        if !ok{
            break
        }
        ch02 <- x*2
    }
}

// 只读 通道
func f03(ch02 <-chan int){
    defer wg.Done()
    for{
        x,ok:= <-ch02
        if !ok{
            break
        }
        fmt.Println("结果：",x)
    }
}

var wg sync.WaitGroup
func main(){
    var ch01,ch02 chan int  //声明 变量
    ch01 = make(chan int,50)
    ch02 = make(chan int,100)
    
    wg.Add(3)
    go f01(ch01)
    go f02(ch01,ch02)
    go f03(ch02)
    wg.Wait()
}