package main

import "fmt"

// 数组 是 值类型， 深cp， 后面的变量 改变 不会 影响 源
// 切片 是 引用类型 ，浅cp，后面的变量 改变 会 影响 源
func main() {
	s1 := make([]int, 5, 10) //用make 函数 定义 一个 长度5，容量10 的切片
	fmt.Printf("len:%d,cap:%d,value:%v \n",len(s1),cap(s1),s1)

    var s01 []int //len:0, cap:0, s01==nil
    s02 := []int{}  //len:0, cap:0, s02 != nil
    s03 := make([]int,0) //len:0, cap:0, s03 !=nil  //容量不写，默认等于长度
    fmt.Printf("len:%d,cap:%d,value:%v,nil=?%v \n",len(s01),cap(s01),s01,s01==nil)
    fmt.Printf("len:%d,cap:%d,value:%v,nil=?%v \n",len(s02),cap(s02),s02,s02==nil)
    fmt.Printf("len:%d,cap:%d,value:%v,nil=?%v \n",len(s03),cap(s03),s03,s03==nil)

    s04 := []string{"北京","上海","深圳"}
    s04 = append(s04, "广州","成都","杭州") //append 用回原来的 变量 接收（容量 会 自己增加，如原来小于1024 就2倍，大于1024 就1/4 地增加）
    fmt.Printf("len:%d,cap:%d,value:%s,nil=?%v \n",len(s04),cap(s04),s04,s04==nil)

    s05 := []string{"珠海","中山"}
    s04 = append(s04, s05...)   // ... 表示把 s05 拆开 
    fmt.Printf("len:%d,cap:%d,value:%s,nil=?%v \n",len(s04),cap(s04),s04,s04==nil)

    // copy 是 深 cp
    s06 := s04  //引用变量
    
    var s07 = make([]string,3)  //make 是 要 自定义 长度，定义的长度不够 就要用 append 来追加了
    copy(s07, s04)  //复制（目标，来自）
    fmt.Printf("len:%d,cap:%d,value:%s,nil=?%v \n",len(s07),cap(s07),s07,s07==nil)
        
    s04[0] = "BJ"
    fmt.Print(s04,"\n",s06,"\n",s07,"\n")

    // 测试题
    var a1 = make([]int,5,10)   //{0 0 0 0 0}
    for i := 0; i < 10; i++ {
        a1 = append(a1, i)
    }
    fmt.Println(a1) // 扩容 -> {0 0 0 0 0 0 1 2 3 4 5 6 7 8 9}
}