package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var m01 = sync.Map{}    //声明并 初始化 了一个 假 map 类型（区别于真的map，只能 用 包 的方法来操作） 

func makeMap(n int){
    defer wg.Done()
    key := strconv.Itoa(n)
        
    m01.Store(key,n)    //操作 map key 写入 （写入时 map 本身 也 支持 动态扩容）
    value,_:=m01.Load(key)  //操作 map key 读取

    fmt.Printf("m01 keyname:%v  value:%v \n",key,value)
}

func main() {
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go makeMap(i)
    }
    wg.Wait()
    fmt.Println("m01:", m01)    //由这里可见，与map 不一样
}