package main

import "fmt"

// 声明变量
var name01 string
var age01 int
var isOk01 bool

// 批量声明
var (
    name02 string
    age02 int
    isOk02 bool
)

func main(){
    name01 = "hello kelvin01" 
    age01 = 10
    isOk01 =true
    
    name02 = "hello kelvin02"
    age02 = 11
    isOk02 = false
    fmt.Printf("value= %s,type= %T \n",name01,name01)
    fmt.Printf("value= %v,type= %T \n",age01,age01)
    fmt.Printf("value= %v,type= %T \n",isOk01,isOk01)

    fmt.Printf("value= %s,type= %T \n",name02,name02)
    fmt.Printf("value= %v,type= %T \n",age02,age02)
    fmt.Printf("value= %v,type= %T \n",isOk02,isOk02)

    var s1 string = "kelvin"
    var s2 = 20
    fmt.Print("hi,",s1,s2)

    // 短变量声明，只能在函数内使用，自动识别类型
    S3:="heihei"
    fmt.Println("hello ",S3)

    // 匿名变量，不用 不消耗内存
    x,_ := foo()
    fmt.Println("x= ",x)
    _,y := foo()
    fmt.Println("y= ",y)

}

// 匿名变量，不用 不消耗内存
func foo()(int,string){
    return 10,"string"
}