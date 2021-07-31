package main

import "fmt"

// 使用 值接收 和 指针接收者 的区别
// 值接收者，值 和 指针 类型都能接收；
// 指针接收者，只能接收 指针类型



func (c cat) move(){
    fmt.Println("walking on road...")
}
func (c cat) eat(food string){
    fmt.Printf("cat eat %s...\n",food)
}

type cat struct{
    name string
    feet int
}

// 同一个 结构体 可以实现多个接口
type mover interface{
    move()
}
type eater interface{
    eat(string)
}

// interface 嵌套
type animal interface{
    mover
    eater
}

func main(){
    var a0 = new(cat)   //interface 嵌套

    var a1 mover
    var a2 eater
    c1 := cat{
        name: "cat01",
        feet:4,
    }
    c2 := cat{"cat02",2}
    
    a1 = c1
    a1.move() 

    a2 = c2
    a2.eat("mouse")

    // interface 嵌套
    a0.move()
    a0.eat("chiken")
    
    // 空接口 用于 关键字（任意） 类型
    var m1 map[string]interface{}   //声明变量
    m1 = make(map[string]interface{},1) //初始化
    m1["name"] = "kelvin01"
    m1["age"] = 42
    m1["hobby"] = []string{"篮球","足球","排球"} //数组

    fmt.Println("任意类型",m1["hobby"])
    mm := m1["hobby"]
    fmt.Printf("%T, %v",mm,mm)

    // 类型断言 interface{} 需要通过 断言 转换 原类型
    switch mm.(type){   //go 固定写法 用来处理 各种 类型
        case []string:
            op,ok:=mm.([]string)    //实际 是这句来转换
            fmt.Println(op,ok)
            if ok{
                op = append(op,"xxx")
                fmt.Println(op)
                mm = op
            }
        
    }
    m1["hobby"] = mm
    fmt.Println(m1)

    m2 := make(map[string]interface{},1) //声明变量 同时初始化
    m2["name"] = "kelvin02"
    m2["age"] = 42
    fmt.Println("任意类型",m2)

    // 空接口 
    fly(m1)
}

// interface 空接口 作为 函数 参数
func fly(x interface{}){   //用于 未知传值类型 的场景
    fmt.Printf("空接口 参数：%v,Type:%T",x,x)
}

