package main

import "fmt"

// 构造函数
type persion struct{
    name string
    age int
}

// 返回 指针 会节约内存开销，特别是结构体构造函数 用得多，内容 又大 的情况下
func makePersion2(name string, age int) *persion {
    return &persion{
        name:name,
        age:age,
    }
}
func makePersion1(name string, age int) persion {
    return persion{
        name:name,
        age:age,
    }
}


func main() {
    p1 := makePersion1("kelvin01",18)
    p2 := makePersion2("kelvin02",28)
    fmt.Println(p1,p2)
}