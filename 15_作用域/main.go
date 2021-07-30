package main

import "fmt"

// 全局、函数、语句块 作用域 3种作用域
var a = 100

func main() {
    var a = 10
    f01(a)
}
func f01(x int) (y int){
    // 函数中查找变量的顺序
    // 1.先在 函数 内部查找(函数 内部 的变量 不能在外部使用)
    // 2.在 外部 查找
    var a = 20 + x
	fmt.Println("f01",a)
    return a
}