package main

import "fmt"

// 结构体 模拟实现 继承
type animal struct{
    name string
}
type dog struct{
    feet uint8
    animal  //animal 的方法 也继承了
}
func (a animal) move ()(){
    fmt.Println(a.name," is moving!?")
}
func (d dog) wanwan(){
    fmt.Printf("%s ：汪汪。。。\n",d.name)
}

func main() {
    // 结构体 声明 并 初始化
    d1 := dog{
        feet: 4,
        animal: animal{name:"小狗01"},        
    }
    d1.wanwan()
    d1.move()   // move 是 继承，来自 animal
    

}