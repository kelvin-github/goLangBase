package main

import "fmt"

// 匿名函数 大多数用在 函数 内部，当然 也可以 赋值 方式 来用在 各种作用域
// 闭包（术语），其实就是 调用了 一个 函数 内部 的匿名函数，同时 也 可以使用这个函数 内部 的变量

// 场景，假定另一个 同事写了个 函数 是 f1()
func f1(f func()){
    fmt.Println("this is Func1")
}

// 你写的是一个 f2()
func f2(x,y int){
    fmt.Println("this is Func2",x+y)
}

// 然后，他不改代码，你要通过 他 的 f1(),才能 做到 某些事。怎么 把你 的f2 通过 传参方式 执行到 f1 上呢 ？
// 封装 一个 f func() 不带参数，不返回 的 方法，以匿名函数方式 将你的 f2 写进去
func f3( f func(x,y int), x int, y int )( func() ) {
    tmp:= func(){
        fmt.Println("this is Func3")
        f(x,y)
    }
    // tmp()
    return tmp
}

func main() {
    x:=100
    y:=200
    ret := f3(f2,x,y)
    // ret()
    f1( ret )

}