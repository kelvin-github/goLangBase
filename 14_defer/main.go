package main

import "fmt"

// defer 理解为栈，先入后出,将标签 后 的语句，延迟到 函数 即将 结束(return) 前 倒序执行
// defer 最常用于 close 的操作，一开始怕忘记，先写好，关闭数据库连接，关闭 文件，关闭连接之类的，待到函数快要结束时执行

// 注意：return 本质是2步执行
// 1.先获取 ret 的赋值
// 2.才是 return
// defer 是 在 赋值后，返回前，所以，不会影响到 返回的结果
func main() {
    // defer01()
    
    fmt.Println( "defer01:", defer01() )
    fmt.Println( "defer02:", defer02() )
    fmt.Println( "defer03:", defer03() )
    fmt.Println( "defer04:", defer04() )
    fmt.Println( "defer05:", defer05() )

    // 证明 压入栈 的时候，变量是什么值，就按当时的值来计算
    a:=1
    b:=2
    defer calc("1",a, calc("10",a,b))   //calc("1",1,3)->calc("1",1,3) "1" 1 3 4
    a =0 
    defer calc("2",a, calc("20",a,b))   //calc("2",0,2)-> cal("2",0,2) "2" 0 2 2
    b =1

}

func calc(i string,a,b int) int{
    ret := a + b
    fmt.Println(i, a , b , ret)
    return ret
}

func defer05()(x int){ //6
    defer func(x *int){
        (*x)++
    }(&x)   //这里把x 的地址传过去了，相当于就是直接修改了内存地址
    return 5    // x = 5
}

func defer04()(x int){  //return x = 5，这个 x 的作用域 是 return的 ??
    defer func(x int){  
        x+=2    //这个时候压入栈 的x 由于没有初始化，int 默认值是0；0+2 =2
        fmt.Println("def04",x)  //x=0+2 = 2
    }(x)    //x = 0 自己的函数（这个时候 x 的著用于 只在 这个匿名函数中，外部是不能 调用的，因此 外部的x 不是 这个x）
    return 5    //ret=5, x= 0+2  
}

func defer0301()(x int){    //return x = x = 5+2
    x=5 //作用域 是 整个函数
    defer func(){
        x+=2    //这个时候压入栈的x=5 +2 = 7
        fmt.Println("def0301",x)
    }()
    return x    //ret=5, x=5+2
}

func defer03()(y int){  //retrun y = x = 5
    x:=5 //作用域是这个函数，
    defer func(){
        x+=2    //这个时候压入栈的x=5+2=7
        fmt.Println("def03",x)  //这个x 确实 是 作用域内的x ，但是竟然 没来得及 赋值 给y 
    }()
    return x    //ret=5, x=ret+2,
}

func defer02()(x int){  //return x = ret + 2
    defer func(){
        x += 2
        fmt.Println("def02",x)
    }()
    return 5    //ret = 5, x = ret +2 , 
}

func defer01()(int){    //return = ret = 5
    x := 5  //初始化 x =5
    defer func(){
        x +=2      //这个时候 压入栈的 x = 5 ->5+2=7
        fmt.Println("def01",x )   
    }()
    return x    //ret=5, x = ret+2 
}

func defer00(){
    defer fmt.Println("001")    //6
    defer fmt.Println("002")    //5
    defer fmt.Println("003")    //4

    fmt.Println("Start")    //1
    fmt.Println("000")   //2
    fmt.Println("End")  //3
}