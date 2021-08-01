package main

import (
	"fmt"
	"io/ioutil"
)

// 打开文件 再写入
func main() {	
    str := "happy birday ！\n"
    writer01(str)
    writer02(str)
}
// 读取 并 追加
func writer02(str string){    
    ret,err := ioutil.ReadFile("./logs.txt")
    if err!=nil{
        fmt.Println("write file fail,err",err)
        return
    }
    temp:=string(ret)+string(str)
    writer01(temp)
}
// 创建|清空 -> 写入
func writer01(str string){
    err := ioutil.WriteFile("./logs.txt", []byte(str), 0666 )
    if err!=nil{
        fmt.Println("write file fail,err",err)
        return
    }  
}