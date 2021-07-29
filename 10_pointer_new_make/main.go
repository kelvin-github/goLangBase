package main

import "fmt"

// go 不存在 指针 操作
// & 根据变量名 获取指针 地址；
// * 根据变量名（指针地址） 取值
func main() {
    n:=18
    fmt.Println(&n)

    p := &n //p 就是 内存地址 *内存地址 
    fmt.Printf("%T, %v, %v \n",p,p,*p)

    *p = 10
    fmt.Printf("%T, %v, %v \n",p,p,*p)

    //声明一个指针
    var a00 *int //nil 因此不能操作毫无意义
    fmt.Println(a00)

    var a01 = new(int)  // new 申请一个内存地址
    *a01 = 100
    fmt.Printf("%T,%v,%v",*a01,&a01,*a01)

    // new 很少用 一般用于 string int 申请内存（基本类型）
    // make 只用于 slice map chan 申请内存，这三种类型 都是 引用类型（浅cp）
    
}