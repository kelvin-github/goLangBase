package main

import "fmt"

// 数组 定义时必须 定义 容量,定义后 不能改变
func main() {
    var a1 [3]bool  //默认 bool:false int:0 string:'' 
    var a2 [4]string  //容量 不一样 不能对比
    fmt.Printf("%T,%T\n",a1,a2)

    // 数组初始化(数组赋值)
    a1 = [3]bool{true,true} //先付2个
    fmt.Println(a1)

    a3 := [...]int{0,1,2,3,4}   // ... 自动根据后面 {} 的数量来设置
    fmt.Println(a3)

    a4 := [5]int{0:1,3:2}   //根据索引来赋值
    fmt.Println(a4)

    // 数组遍历
    citys := []string{"北京","上海","深圳"}
    for i,item := range citys{
        fmt.Println(citys[i],item)
    }

    // 多维数组
    // var a11 [3][2]int
    a10:=[...]int{1,3}
    a11 := [...][2]int{
        [2]int{1,2},
        [2]int{3,4},
        [2]int{5,6},
        a10,
    }
    fmt.Println(a11)
    // 多维数组遍历
    for _,item01 :=range a11{        
        for _,item02 := range item01{
            fmt.Printf("%v ",item02)
        }
        fmt.Println(item01," - ")
    }

    // 数组 是 值类型 (赋值时，深copy; 引用类型, 浅copy) 
    a12 := a10
    a12[0] = 20
    fmt.Println(a10,a12)

    a13:=arraySum(a10)  //数组内 求和
    fmt.Println(a13)

    // 切片 slice（可变长度的序列，支持自动扩容）定义数组时 [] 里面 不写数字就是切片
    var s01 []int   //定义了一个存放 int 类型元素的 切片
    var s02 []string //定义了一个存放 string 类型元素 的切片    
    fmt.Println(s01==nil,s02==nil)    //切片 默认值 nil

    // 切片初始化
    s01 = []int{1,2,3}
    s02 = []string{"长江","黄河","珠江"}
    fmt.Println(s01,s02) 

    // 切片 的长度lang() 容量cap(切片中的容量 是 基础 的数组 的 容量) 
    
    // 切片 来自 数组
    a01 := [...]int{1,2,3,4,5}
    s03 := a01[0:2]     //左包 右不包（左闭右开），索引 0-3 赋值{1,3}  
    fmt.Printf("a01:%v, %T\n",a01,a01)
    fmt.Printf("长度：%v, 容量： %v\n",len(a01),cap(a01)) //a01 数组中的 长度=容量=5

    fmt.Printf("s03:%v, %T\n",s03,s03)  //{3,4}
    fmt.Printf("长度：%v, 容量： %v\n",len(s03),cap(s03)) //s03 切片中的 长度=引用数量，容量=原数组 :前索引号 开始计数到最后一个=5：（从原数组首个 索引号 算到最后一个，即使没有全部赋值过来）

    s04 := a01[:]
    fmt.Printf("s04:%v, %T\n",s04,s04)  //全部
    fmt.Printf("长度：%v, 容量： %v\n",len(s04),cap(s04)) //切片中的 容量5：（从原数组首个 索引号 算到最后一个，即使没有全部赋值过来）
    
    s05 := a01[3:]
    fmt.Printf("s05:%v, %T\n",s05,s05)  //从索引3到 结束 {4,5}
    fmt.Printf("长度：%v, 容量： %v\n",len(s05),cap(s05)) //切片中的 容量2：（从原数组首个 索引号 算到最后一个，即使没有全部赋值过来）

}

func arraySum(x [2]int) int{
    total := 0
    for _,v1 := range x{
        total = total + v1 
    }
    return total
}