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
}


func defer04()(x int){  //return x = x ??
    defer func(x int){
        x+=2
        fmt.Println("def04",x)  //x=0+2 =2
    }(x)    //x = 0 自己的函数
    return 5    //ret=5, x=0+2 
}
func defer0301()(x int){    //return x = x = 5+2
    x=5
    defer func(){
        x+=2
        fmt.Println("def0301",x)
    }()
    return x    //ret=5, x=5+2
}

func defer03()(y int){  //retrun y = ret
    x:=5
    defer func(){
        x+=2
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
    x := 5
    defer func(){
        x +=2      //修改 的 是 x 但竟然不是 返回 的那个 x
        fmt.Println("def01",x )   
    }()
    return x    //ret=5, x=ret+2, 
}

func defer00(){
    defer fmt.Println("001")    //6
    defer fmt.Println("002")    //5
    defer fmt.Println("003")    //4

    fmt.Println("Start")    //1
    fmt.Println("01")   //2
    fmt.Println("End")  //3
}