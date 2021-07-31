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
// 接收者表示 调用方法 的 具体类型 变量，多用类型首字母小写 表示
func (d dog) wang() {
    fmt.Printf("%s:汪汪\n",d.name)
}

func main(){
    d1 := newDog("dog01")
    d2 := newDog("dog02")
    d1.wang()
    d2.wang()
}