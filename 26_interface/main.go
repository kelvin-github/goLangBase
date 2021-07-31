package main

import "fmt"

// 值类型(int float bool string [x]数组 struct )赋值后  修改 副本 不影响 源 深cp
// 引用类型 （map []slice chan interface）赋值后 修改副本 会引起 源变化，浅cp

// interface 是一种 特殊的类型，它规定了变量有哪些方法
// 接口实例

type cat struct{
    name string
}
type dog struct{
    name string
}
func (c cat) speak(){
    fmt.Println("miaomiao...")
}
func (d dog) speak(){
    fmt.Println("wanwanwan...")
}

// 定义一种 接口 类型
type speaker interface{
    speak() //这个类型中 有这种方法
}
// 为接口 类型 定义一个 接收者 方法
func bit( i speaker ){
    fmt.Println(i)
    i.speak()
}

func main() {
    var c1 = cat{   //声明 一个变量 为 cat 结构体 的类型
        name:"cat01",
    }
    var d1 = dog{   //声明 一个变量 为 dog 结构体 的类型
        name:"dog01",
    }
    bit(c1)
    bit(d1)

    var i01 speaker //声明 一个变量 为 speaker 的接口类型
    i01 = c1    //能接收 任何 结构体 类型，并获得了，speaker 类型 中的 方法
    i01.speak()

}