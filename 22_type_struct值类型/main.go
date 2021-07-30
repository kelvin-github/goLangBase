package main

import "fmt"

// 结构体 要 比 map 自由些
// 值类型(int float bool string [x]int struct )赋值后  修改 副本 不影响 源 深cp
// 引用类型 （map slice chan interface）赋值后 修改副本 会引起 源变化，浅cp
// go 中 函数 传参 ，全是 深cp 了一个 过去，如果 &变量，才是 源 修改
type persion struct{
    Name string `json:"name"`
    Age int `json:"age"`
    Gender string   `json:"gender"`
    Hobby []string  `json:"hobby"`
}

func main() {
    var p01,p02 persion
    p01.Age = 42
    p01.Name = "kelvin"
    p01.Gender = "男"
    p01.Hobby = []string{"篮球","足球"}
    fmt.Printf("%T %v \n",p01,p01)

    p02.Age =40
    p02.Name = "Yuang"
    p02.Gender = "男"
    p02.Hobby = []string{"排球","足球"}
    fmt.Printf("%T %v \n",p02,p02)

    // 匿名结构体(用于临时场景)
    var s01 struct{
        Name string
        Age int
    }
    s01.Name = "kelvin"
    s01.Age = 42
    fmt.Printf("%T %v \n",s01,s01)

    f1(p01)
    fmt.Println(p01.Age)    //42

    f2(&p01)    //
    fmt.Println(p01.Age)

}
func f2(p *persion){
    (*p).Age = 1
}

func f1(p persion){
    p.Age = 10
}