package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var x int = 0

func add(){
    defer wg.Done()
    
    var once sync.Once  //确保 作用域内 某个方法  只 操作 1次
    for i := 0; i < 10; i++ {
        x += i
        fmt.Println("x = ", x)
        if x > 5{
            once.Do( func(){
                fmt.Println("once")
                resteX()
            })
        }
    }

}
func resteX(){
    x = 0
}
func main() {
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go add()
    }
    wg.Wait()
    fmt.Println("main still alive!")
}