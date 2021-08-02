package main

import (
	"fmt"
	"sync"
)

// worker_pool 概念，下面 的写法 就是  工作池 模式
// 目标 1个gorotuine 生产 100 个数据放入ch01，2个goroutine 消费 ch01数据 平方 操作后，放入 到 ch02 中，完成后，才再被读出来

var wg sync.WaitGroup
var ch01,ch02 chan int

func f01(ch01 chan int){
    defer wg.Done()
    defer close(ch01)
    for i := 1; i < 101; i++ {
        ch01 <- i
    }
}
func f02(ch01,ch02 chan int){
    defer wg.Done()
    for {
        x,ok:= <-ch01
        if !ok {
            break
        }
        x *=x
        ch02<-x
        // fmt.Println("x平方值：",x)        
    }
}

func main(){
    ch01 := make(chan int ,100)
    ch02 := make(chan int ,100)

    wg.Add(3)
    go f01(ch01)
    go f02(ch01,ch02)
    go f02(ch01,ch02)
    wg.Wait()
    close(ch02)   
    
    var i int = 0
    for{    // channel 这种 每取走一个 长度会变 的 类型，range 处理不了
        x,ok:=<-ch02
        if !ok{
            break
        }
        i++
        fmt.Println("目标结果：",x , i)
    }
    fmt.Println("main 还活着")

}