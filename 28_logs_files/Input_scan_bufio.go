package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
    // inputScan()
    inputBufio()
}

func inputScan() {
	var s string //用于收集来自 scran屏幕的输入
	fmt.Println("please input string:")
    fmt.Scanln(&s)  //遇到空格会丢失
    fmt.Println("获取到的输入内容：",s)
}
func inputBufio(){
    var s string
    fmt.Println("please input string:")
    reader := bufio.NewReader(os.Stdin) //标准终端输入 Stdout 标准终端输出
    s,err := reader.ReadString('\n')
    if err!=nil{
        fmt.Println("获取异常,err",err)
    }
    fmt.Println("获取到的输入内容：",s)

}