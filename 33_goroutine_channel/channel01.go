package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 如果用 共享内存 或 队列（因协程竞争 产生 所用 互斥锁 衍生问题），造成性能问题
// channel 先入先出 （水管） 但是 设定缓冲长度 就不能改，一旦关闭 就不能再放入，但还可以取走

var wg sync.WaitGroup
var ar01 []int  //声明一个 切片（引用类型 必须 初始化 才能用）
var ch01 chan int    //声明变量类型 channel 是 引用类型，必须初始化才能用
var ch02 chan *string
func main() {
    // channel01() //不设置 水管 缓存 的用法
    channel02() //设置 100 的容量，一边换一边取

}


func channel02(){
    ch02 = make(chan *string ,100)   //channel 使用 make 初始化 返回的是一个 内存地址 100单位缓冲区
    fmt.Println(ch02)   
    defer close(ch02)   //主程序 退出后 是会自动关闭 的，主动关闭 是 因为有些 场景，主程序 会一直 开着，不用就老实关上
    for i := 0; i < 50; i++ {
        str := strconv.Itoa(i)
        ch02 <- &str

        wg.Add(1)
        go func(){
            wg.Done()
            x := <-ch02
            fmt.Printf("从水管中取出地址：%v,值：%v \n", x, *x)
        }()

    }
    wg.Wait()
    fmt.Println("main 还活着")
}
 
func channel01(){
    ch01 = make(chan int)   //初始化 不定义长度，赋值的时候会卡住，但是 如果 后台 有一个携程接收，就 没问题
    wg.Add(1)
    go func(){
        defer wg.Done()
        x:=<-ch01   //从 指定的 水管 取出
        fmt.Println("匿名协程接收到：",x)
    }()

    ch01<-10    //放入 指定 的水管    

    wg.Wait()
    fmt.Println("main 还活着")
}