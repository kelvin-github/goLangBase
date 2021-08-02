package main

import "fmt"

// 多路复用，就是 那个 通道 能运行，就走哪里(都可以的情况下，随机走)都不能走，就走 default
func main(){
    f01()

}

func f01(){
    ch01:=make(chan int, 10)
    for i := 0; i < 10; i++ {

        select{
        case x:=<-ch01 :
            fmt.Println("获取成功:", x)

        case ch01<-i:
            fmt.Println("输入成功:", i) 
        
        default:
            break
        }
        

    }
}