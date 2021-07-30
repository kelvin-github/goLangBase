package main

import "fmt"

// 自定义 类型 和 类型别名 （意义：例如自己定义一些 特定 代码，错误类型之类，给自己一眼识别 干嘛用的）
type myInt int  //自定义类型，编译后，还是 这个类型
type youInt = int   //类型别名（编译后，就回到 原来 的类型 int）
func main() {
    var m  myInt = 100
    var n youInt = 200
    fmt.Printf("%T, %v\n",m,m)
    fmt.Printf("%T, %v\n",n,n)

    var c rune = '中'   //utf8 字符编码 int32
    fmt.Printf("%T, %v\n",c,c)

}   