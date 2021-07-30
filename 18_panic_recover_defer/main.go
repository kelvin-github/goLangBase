package main

import "fmt"

// panic recover 成对出现； 有recover 一定由 defer 先压入栈，埋伏好
func main() {
    fun01()
    fun02()
    fun03()
    fun04()
}

func fun01(){
    fmt.Println("func01")
}
func fun02(){
    defer func(){
        err := recover()  //固定函数，直接获取到错误的信息 
        if err!=nil{
            fmt.Println("释放 。。。错误处理:", err)
        }
        
    }()
    // panic("\r\n错误出现 02 ")   //程序 就到这里 break 了，为了能让他不要 break,可以继续 走下去，执行 recover()  
    fmt.Println("func02")
}
func fun03(){
    fmt.Println("func03")
}
func fun04(){
    fmt.Println("func04")
}