package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)


func main() {
	fileObj, err := os.OpenFile("./logs.txt",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)  //如果文件不存在，就创建，区别 Open
	if err != nil {
		fmt.Println("open file failed,err:",err)
        
        return
	}
    defer fileObj.Close()
    
    var context string
    var buf [256]byte
    reader := bufio.NewReader(fileObj)
    for{
        // 读取文件
        n,err := reader.Read( buf[:] )   //以行为单位读取的
        if err==io.EOF{    //设定一个跳出 的条件
            goto x01            
        }
        context += string(buf[:n])
        fmt.Println( string(buf[:n]) )
    }
    
    x01:
    fmt.Println(context)    
}