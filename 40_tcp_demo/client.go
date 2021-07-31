package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client
func main(){
    // 1.与server端 建立连接
    conn,err:=net.Dial("tcp","127.0.0.1:20000")
    if err!=nil{
        fmt.Println("Dial failed ,err",err)
        return
    }
    defer conn.Close()  //关闭连接
    
    // 2.收发信息
    inputReader := bufio.NewReader(os.Stdin)
    for{
        fmt.Println("please input :")
        // fmt.Scanln(&msg)
        input,_ := inputReader.ReadString('\n')   //读到换行
        inputInfo := strings.Trim(input,"\r\n")  //
        if strings.ToUpper(inputInfo) == "EXIT"{
            return
        }

        // 发出
        _,err = conn.Write( []byte(inputInfo) )
        if err!=nil{
            return
        }

        // 接收
        buf := [512]byte{}
        n,err := conn.Read(buf[:])
        if err!=nil{
            fmt.Println("recv failed,err",err)
            return
        }
        fmt.Println( string(buf[:n]) )
        
    }
}