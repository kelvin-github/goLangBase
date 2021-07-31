package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client
func main(){
    socket,err:=net.DialUDP("udp",nil,&net.UDPAddr{
        IP:net.IPv4(127,0,0,1),
        Port: 40000,
    })
    if err!=nil{
        fmt.Println("udp connect fail,err",err)
        return
    }
    defer socket.Close()

    // 收发数据
    reader := bufio.NewReader(os.Stdin)
    var reply [1024]byte
    for{
        // 发数据
        msg,_:=reader.ReadString('\n')
        socket.Write( []byte(msg) )
        

        // 收数据
        n,_,err :=socket.ReadFromUDP( reply[:] )
        if err!=nil{
            fmt.Println("reply msg fail,err",err)
            return
        }
        fmt.Println("get reply:",string( reply[:n]) )
    }
    
}