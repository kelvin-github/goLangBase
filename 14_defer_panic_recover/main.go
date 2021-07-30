package main

import "fmt"

// defer 理解为栈，先入后出,将标签 后 的语句，延迟到 函数 即将 结束(return) 前 倒序执行
func main() {
    defer01()


}

func defer01(){
    defer fmt.Println("001")    //6
    defer fmt.Println("002")    //5
    defer fmt.Println("003")    //4

    fmt.Println("Start")    //1
    fmt.Println("01")   //2
    fmt.Println("End")  //3
}