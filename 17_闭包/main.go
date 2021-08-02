package main

import "fmt"

// 匿名函数 大多数用在 函数 内部，当然 也可以 赋值 方式 来用在 各种作用域
// 闭包（术语），其实就是 调用了一个 匿名函数，同时 也可以使用 这个匿名 外部 的变量
// 底层原理：
// 1.函数可以当返回值用
// 2.函数的变量，先从函数内部找，找不到，往上一层找。。。

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
    // tmp()    //tmp 这个时候 就是 一个函数，可以（）执行
    return tmp
}

func f4( x int )( f func(int)int ){    
    temp:= func(y int) int {
        return x + y
    }
    return temp
}

func f5(base int)( func(int)int, func(int)int ){
    add := func(i int) int{
        base += i
        return base
    }
    sub := func(i int) int{
        base -= i
        return base
    }
    return add,sub
}

func main() {
    x:=100
    y:=200
    ret := f3(f2,x,y)   //ret 这个时候 也成了一个函数
    // ret()
    f1( ret )

    ret04 := f4(100)
    fmt.Println( "闭包调用：", ret04(200) )

    // 闭包（公共值）变量 会随使用 而变化
    add,sub :=f5(10)
    fmt.Println("add",add(1), "sub",sub(2)) //10+1=11；11-2=9
    fmt.Println("add",add(1), "sub",sub(2)) //9+1=10；10-2=8

}