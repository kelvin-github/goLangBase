package main

import (
	"fmt"
	"os"
)

// 打开文件 再写入
func main() {
	fileObj, err := os.OpenFile("./logs.txt", os.O_CREATE|os.O_APPEND,  0644)  //没有就创建，追加
    if err != nil {
		fmt.Println("open file failed,err:",err)        
        return
	}
    defer fileObj.Close()

    // 写入内容
    str := "happy birday ！\n"
    fileObj.Write( []byte(str) )    //写法1

    fileObj.WriteString(str)    //写法2
}