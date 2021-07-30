package main

import "fmt"

// 有进 有出
func sum(x, y int) int {
	return x + y
}

// 多个 参数 进入
func f01(x string,y ...int) (string) {
    fmt.Println("param01",x)
    fmt.Println("param02",y)
    var x01 string 
    x01 = "param01:" + x
    return x01
}

// 多个进，多个出
func f02(x ,y int)(sum, sub int){
    sum = x + y
    sub = x - y
    return 
}

func main() {
    x := 10
    y := -20
    sum := sum(x, y)
	fmt.Println("普通：",sum)

    s := "hi kelvin"
    result := f01(s,1,2,3,5)
    fmt.Println("多进单出：", result)

    param01,param02 := f02(x,y)
    fmt.Println("多进多出：",param01,param02)

    
    // 匿名函数 01   
    sum01 := func (x,y int) (int) {
        return x + y
    }
    fmt.Println("匿名函数01 ：", sum01(x,y) )   //这里传入参数

    // 自执行 函数 
    sum02 := func (x,y int) int {
        // fmt.Println(x,y) 
        return x+y       
    }(x,y)  //这里传入参数
    fmt.Println( "自执行 函数: ",sum02 )

}