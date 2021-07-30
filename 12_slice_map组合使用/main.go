package main

import "fmt"

//
func main(){
    // 值为 map 的 slice
    var s01 = make([]map[string]string,10) //初始化 一个 长度容量10 的切片，切片的值是 map 对象，这个时候的map 对象 还没有初始化
    fmt.Printf("%T\n",s01)
    s01[0] = make(map[string]string,1) //对 这个切片 指定 索引 值 赋值一个 初始化 后的 map 对象，才能直接赋值

    s01[0]["key01"] = "value01"
    s01[0]["key02"] = "value02"
    s01[0]["key03"] = "value03" //居然可以 超出 make(map[string]string,2) 定义 的长度 赋值 为什么？不知道
    s01[0]["key04"] = "value04"
    s01[0]["key05"] = "value05"
    fmt.Println(s01)
    // s01[0] = {1:"k01",2:"k02"}


    // 值为 slice 的 map
    var m01 = make(map[string][]string,1)  // map 似乎 是可以不定义 长度 和 容量的（只要 用 make 初始化 就可以随意用了）
    m01["key01"] = make([]string,10)
    m01["key01"] = []string{"北京","上海","深圳","广州"} 
    
    m01["key02"] = make([]string,2)
    m01["key02"][0] = "BJ"
    m01["key02"][1] = "SH"
    // m01["key02"][2] = "BJ"   //slice 严格 不能超出 长度
    m01["key02"] = append(m01["key02"], "GZ")   //但是 切片 嘛，还是 可以 append 的

    fmt.Println(m01)

}