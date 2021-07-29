package main

import "fmt"

// 定义常量 及 iota(获取到当前 定义常量 的个数，包括自己,每 新增 一行 iota+1,同一行 定义多个常量 不加+1 )
const (    
    a1 = iota   // a1 = 0
    
    pi01 = 3.1416   
    pi02
    pi03
    
    a2=iota  //a2 = a1 +1
    a3  //a3 = a2 +1    
)
//新 的 const 出现 ，iota 重置 为0 (理解为 const 下 的索引，但是 获取 是 要把自己 算上，很奇怪的函数)
const(
    b1 = iota
    b2
    _
    b3  //b3 =3 
)
// iota 特性 ：每 新增 一行 iota + 1,同一行 定义多个常量 不加+1
// 毒：iota 本质是一个函数方法，不是 变量，因此 赋值时，它只是记录了const 中的，定义变量 的函数
const(
    c1,c2 = iota+1,iota+2   //0 + 1=1, 0 + 2=2
    c3,c4 = iota+1,iota+2   //1 + 1=2, 1 + 2=3
)

// 定义数量级？
const(
    _ = iota
    KB = 1 << (10 * iota)   // << 表示二进制 1 左移10位，二进制结果->10000000000 = 十进制 2的10次方 = 1024 * 1
    MB = 1 << (10 * iota)   // 10 *2 = 20，表示二进制 1左移20位，二进制 结果 100000000 0000000000，转换成十进制，就是1024 * 1024 = 1024KB =1MB
    GB = 1 << (10 * iota)   // 10* 3 = 30，二进制1 左移30位 与上同理
    TB = 1 << (10 * iota)   
    PB = 1 << (10 * iota)   
)

func main() {
	fmt.Printf("%v,%T\n",pi01,pi01)
	fmt.Printf("%v,%T\n",pi02,pi02)
	fmt.Printf("%v,%T\n",pi03,pi03)

    fmt.Printf("%v,%T\n",a1,a1)
	fmt.Printf("%v,%T\n",a2,a2)
	fmt.Printf("%v,%T\n",a3,a3)

    fmt.Printf("%v,%T\n",b1,b1)
	fmt.Printf("%v,%T\n",b2,b2)
	fmt.Printf("%v,%T\n",b3,b3)

	fmt.Printf("%v,%T\n",c1,c1)
	fmt.Printf("%v,%T\n",c2,c2)
	fmt.Printf("%v,%T\n",c3,c3)
	fmt.Printf("%v,%T\n",c4,c4)

	fmt.Printf(" -%v\n -%v\n -%v\n -%v\n -%v\n",KB,MB,GB,TB,PB)
	
}