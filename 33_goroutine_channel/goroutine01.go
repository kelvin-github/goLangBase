package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 并发：一个线程，轮流执行任务 1个核执行2个协程 go 函数
// 并行：多个线程，同时执行任务 2个核各做各的协程
// 用户态：自定义程序 的线程；
// goroutine 管理模型 GMP （G协程 M设备 P管理者) go一个协程 2kb+- java一个OS线程2MB
var wg sync.WaitGroup

func makeRand(){    
    defer wg.Done() //goroutine 结束就 wg-=1
    // 随机数
    rand.Seed( time.Now().UnixNano() )  //以纳秒 来 参数 一个随机种子(不是一定要)
    r1 := rand.Int()
    r2 := rand.Intn(100)   //0 <= x <100
    fmt.Println("随机数：", r1,r2)
}

// main 其实本身就是 一个线程
func main(){
    runtime.GOMAXPROCS(2)   //让cpu 跑几个核？ 4个 不写默认，有几个跑几个
    timeStart:= time.Now()  //记录开始时间
    
    for i := 0; i < 100; i++ {
        wg.Add(1)   //启动1个 goroutine 登记 wg+=1
        go makeRand()

        wg.Add(1)   //启动1个 goroutine 登记 wg+=1
        go func(i int){
            defer wg.Done() //做完就 wg-=1
            fmt.Println("匿名函数：",i)
        }(i)    //传入
    }
    wg.Wait()   //等待 所有 goroutine wg==0 结束
    
    timeEnd := time.Now()   //记录结束时间
    timeDs := timeEnd.Sub(timeStart)    //计算时间差值
    fmt.Println("用时间：",timeDs.Seconds(),"秒")
}

