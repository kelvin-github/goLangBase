package main

import "fmt"

// 定义结构体
type dog struct {
    name string
}
// 构造函数
func newDog(name string) dog {
    return dog{
        name:name,
    }
}

// 方法是作用于 特定类型的 函数 （为这个结构体 定义 一个绑定 的 方法，表行为 ）
// 接收者(d dog) 表示 调用方法 的 具体类型 变量，多用类型首字母小写 表示
// Go 中 首字母大写的标语 都是 可以暴露出去的，相当与js export 一些变量 或 函数
func (d dog) wang() {
    fmt.Printf("%s:汪汪\n",d.name)
}

type persion struct {
    name string
    age int
}

// 值接收者 (有 返回)
func (p persion) growup01 (n int) (int){
    age := p.age + n
    return age
}

// 指针接收者 （直接 修改 源）
func (p *persion) growup02(n int){
    p.age +=n
}

func main(){
    d1 := newDog("dog01")
    d2 := newDog("dog02")
    d1.wang()
    d2.wang()

    // 指针方法 来初始化
    p1 := &persion{
        name:"kelvin",
        age:42,
    }
    fmt.Println( "返回值：",p1.growup01(2) )

    p1.growup02(2)
    fmt.Println( "源的修改：",p1.age )
}