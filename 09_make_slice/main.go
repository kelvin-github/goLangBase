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


}