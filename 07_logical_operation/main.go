package main

import "fmt"

// 运算符


func main() {
    var(
        a = 5
        b = 2
    )
    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Println(a / b)
    fmt.Println(a % b)  //求余数 5/2 余1

    // a++ //go 中 这是 单独语句 不能写成 a = a++

    // 关系运算符(需要 类型 相同 才能 比较)
    fmt.Println( a == b )   //等于
    fmt.Println( a != b )   //不等于
    fmt.Println( a > b )    //大于
    fmt.Println( a >= b )   //大于等于
    fmt.Println( a < b )    //小于
    fmt.Println( a <= b )   //小于等于

    // 逻辑运算符
    age:=22
    if age>18 && age <60 {
        fmt.Println("并")   //与
    }

    if age<23 || age > 60{  //或
        fmt.Println("或")
    }

    isMarried := false
    fmt.Println(!isMarried) //非 取反

    // 位运算符,操作是 针对 二进制 数
    // 5 二进制 101
    // 2 二进制  10

    // & 按位与 ，对齐看 上下 都是1 才能得到1)
    fmt.Printf("二进制：%b, 十进制：%d \n", 5 & 2, 5 & 2)  // 000

    // | 按位或 ，对齐看 上下 有一个为1 得 1 
    fmt.Printf("二进制：%b, 十进制：%d \n", 5 | 2, 5 | 2)  // 111
    
    // ^ 相异或 对齐看 上下 不一样 才能得到1 
    fmt.Printf("二进制：%b, 十进制：%d \n", 5 | 2, 5 | 2)  // 111


    // << 左移符
    fmt.Printf("二进制：%b, 十进制：%d \n", 5 << 1, 5 << 1) //101 ->1010

    // >> 右移符
    fmt.Printf("二进制：%b, 十进制：%d \n", 5 >> 1, 5 >> 1)    //101 -> 10 

    // 赋值运算符

    x:=10
    x+=1    // x = x+1 或 x++
    x-=1    // x = x-1 或 x--
    x*=2    // x = x*2  相乘后赋值
    x/=2    // x = x/2  相除后赋值
    x%=2    // x = x%2  取余后赋值
    
    x<<=1   // x = x<<1 //左移1位
    x>>=1   // x = x>>1 //右移1位

    x&=2    // x = & 1  // 和 2 转成二进制后， 按位 与  （上下都是1，才能 1）
    x|=2    // x = x | 1 // 和 2 转成二进制后， 按位 或 （上下有一个1，就得 1）
    x^=2    // x = x ^ 1 // 和 2 转成二进制后，相异或 （上下 不一样 才能得到1）
    fmt.Println(x)    

}