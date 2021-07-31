package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server 端
func main(){
    // 1.本地端口启动服务
    listener, err := net.Listen("tcp","127.0.0.1:20000")
    if err != nil{
        fmt.Println("start server failed ",err)
        return
    }

    // 2.等待 别人 连接过来
    for {
        conn,err := listener.Accept()
        if  err!=nil {
            fmt.Println("accept failed ",err)
            return
        }
        go processConn( conn )  //异步处理
    }
}

func processConn(conn net.Conn){
    // 3. 与客户端通讯
    var temp [128]byte
    inputReader := bufio.NewReader(os.Stdin)
    for{
        n,err:=conn.Read(temp[:])
        if err!=nil{
            fmt.Println("read from conn failed ,err:",err)
            return
        }
        fmt.Println(string(temp[:n]))

        //回复
        fmt.Println("please input :")        
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

    }
    
}