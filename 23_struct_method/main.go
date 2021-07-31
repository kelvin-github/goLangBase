package main

import "fmt"

// 定义结构体
type dog struct {
    name string
}
// 构造函数
func newDog(name string) dog {
    return dog{
        name:name,
    }
}

// 方法是作用于 特定类型的 函数 （为这个结构体 定义 一个绑定 的 方法，表行为 ）
func (d dog) wang() {
    fmt.Println("汪汪")
}

func main(){
    d1 := newDog("dog01")
    d1.wang()
}