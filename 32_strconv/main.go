package main

import (
	"fmt"
	"strconv"
)

// 字符串 转换 包
func main(){
    var i01 int = 100
    i02 := strconv.Itoa(i01)
    fmt.Printf("Int 2 String:%v, Type:%T\n", i02, i02)

    i03,_ := strconv.Atoi(i02)
    fmt.Printf("String 2 Int:%v, Type:%T\n", i03, i03)

}