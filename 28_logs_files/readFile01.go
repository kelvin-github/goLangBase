package main

import (
	"fmt"
	"os"
)


func main() {
	// fileObj, err := os.OpenFile("./logs.txt",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)  //如果文件不存在，就创建，区别 Open
	fileObj, err := os.Open("./logs.txt")
	if err != nil {
		fmt.Println("open file failed,err:",err)        
        return
	}
    defer fileObj.Close()
    
    // var temp = make( []byte,128 )   //指定 读的长度
    var temp [128]byte //这是个数组，字节类型 128字节（每次读取128个字节）
    var context string
    for{
        // 读取文件
        n,_ := fileObj.Read(temp[:])
        if n==0{    //设定一个跳出 的条件
            goto x01            
        }
        context += string(temp[:n])
        fmt.Println( string(temp[:n]) )
    }
    
    x01:
    fmt.Println(context)
}