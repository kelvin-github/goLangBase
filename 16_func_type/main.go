package main

import "fmt"

// 高阶函数
func main() {
	a := f2 // 函数传递
	res01 := f1(a)   // 函数嵌套
	res02 := f1(f2)   // 函数嵌套
    fmt.Println( res01,res02 )

    // 将函数 作为 返回值 返回
    b := f3(f2)
    fmt.Println("将函数 作为 返回值 返回" ,b(10,20) )
}

func f3( a func(x,y int)int )( func(x,y int)int ){
    x :=10
    y:=20
    res := a(x,y)
    fmt.Println(res)

    b := f2 //将 b 函数 返回去
    return b
}


func f1( a func(x, y int)(res int) ) int {
	x := 100
	y := 200
    res := a(x,y)	
    return res
}

func f2(x,y int) int{
    return x+y
}