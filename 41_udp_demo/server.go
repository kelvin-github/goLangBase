package main

import (
	"fmt"
	"net"
	"strings"
)

// udp server
// 适用于直播
func main(){
    conn,err:=net.ListenUDP("udp",&net.UDPAddr{
        IP:net.IPv4(127,0,0,1),
        Port:40000,
    })
    if err!=nil{
        fmt.Println("Listen udp failed ,err",err)
        return
    }
    defer conn.Close()

    // 不需要建立连接，自己收发数据
    var data [1024]byte
    for{
        // 接收数据
        n,addr,err:=conn.ReadFromUDP(data[:])   //通过 连接类，读取数据的时候，同事也获取到来源（下面就返回来源）
        if err!=nil{
            fmt.Println("read from UDP fail,err",err)
            return
        }
        fmt.Println( string(data[:n]) )

        // 发送数据
        reply := strings.ToUpper(string(data[:n]))        
        conn.WriteToUDP([]byte(reply),addr)

    }
}