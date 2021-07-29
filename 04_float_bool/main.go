package main

import (
	"fmt"
)

// 浮点数 go 默认 小数 都是 float64

func main() {
    // math.MaxFloat32
    f1 := 3.14156535
    fmt.Printf("%v,%T\n",f1,f1)

    f2 := float32(3.14156535) //强制转 float32 小数点后7位
    fmt.Printf("%v,%T\n",f2,f2)

    f3 := float64(f2)
    fmt.Printf("%v,%T\n",f3,f3)

    b1 := true
    var b2 bool //默认是 false
    fmt.Printf("%v , %T\n",b1,b1)
    fmt.Printf("%v , %T\n",b2,b2)
} 