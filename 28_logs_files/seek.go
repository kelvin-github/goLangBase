package main

import (
	"fmt"
	"os"
)

// 分段读取，使用光标偏移
func main() {
    var str string ="插入的内容"
    seekRead(str)
}

func seekRead(str string){
    fileObj, err := os.OpenFile("./logs.txt",os.O_CREATE,0644)
    fileObj.Seek(3,0)   //光标位置 从0偏移到3
	if err != nil {
		fmt.Println("open file failed,err:",err)        
        return
	}
    defer fileObj.Close()

    fileObj.WriteString(str)    //写入（复杂一点的写法，就是 读取，记下 n ，拼接写入临时文集，返回源文件 偏移到 n 继续 读取。。。）

    
}