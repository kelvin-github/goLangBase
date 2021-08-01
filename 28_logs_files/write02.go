package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件 再写入
func main() {
	fileObj, err := os.OpenFile("./logs.txt", os.O_CREATE|os.O_TRUNC,  0644)  //没有就创建，清空
    if err != nil {
		fmt.Println("open file failed,err:",err)        
        return
	}
    defer fileObj.Close()

    // 写入内容
    str := "happy birday ！\n"
    b := bufio.NewWriter(fileObj)
    b.WriteString(str)  //这里是写入到缓存中
    b.Flush()   //bufio 这里才是真正 写入到文件
}