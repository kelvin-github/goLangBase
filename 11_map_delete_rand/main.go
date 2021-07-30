package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map 提供映射关系，内部使用散列表（hash）实现，key-value 结构，引用类型（声明时/后，就必须初始化，即赋值）
func main() {
    var m01 map[string]int
    m01 = make(map[string]int,0)   //要估算好 map 容量，避免 动态扩容
    m01["key01"] = 1
    m01["key02"] = 2

    fmt.Println(m01,m01["key01"])

    // 判断有无这个 key
    v,ok := m01["key02"]
    if ok {
        fmt.Println(v)
    }else{
        fmt.Println("查无此key")
    }

    // map 的遍历 原本数组 或 切片中的索引号 参数1，这里是 key 名字, 参数2 是内容值
    for k, v := range m01{
        fmt.Println(k, v)
    }

    // 删除 map 中的key 
    delete(m01,"key02")
    fmt.Println(m01)

    delete(m01,"不存在的key")   //也不会出错，就是删不到而已，不操作 no-op
    fmt.Println(m01)

    // map 将key 排序后 ，再遍历输出（思路是取出 key 排好，得到这个数组，来 遍历 map）
    rand.Seed( time.Now().UnixNano() ) //初始化随机数种子
    var scoreMap = make(map[string]int,10)  //初始化 一个 map 长度10，容量10
    for i := 0; i < 10; i++ {
        keyName := fmt.Sprintf("stu_%02d",i)    //动态生成 keyName 
        scoreMap[keyName] = rand.Intn(100)  //生成 0-99 的随机整数
    }
    keyNames := keySort(scoreMap)
    for _,value := range keyNames{
        fmt.Println( value, scoreMap[value] ) 
    }

}

func keySort( m map[string]int ) []string {
    var keyNames = make([]string,0)    //定义一个切片，用于存放 keyName
    for keyName := range m {    // 将传过来的 map 提取 keyName，
        keyNames = append(keyNames, keyName)    //放到 切片slice          
    }
    
    sort.Strings(keyNames)  //利用 go 自带 的sort 排序包 对 切片 的字符串 进行排序(默认 升值 asc )
    return keyNames //返回这个切片（按这个排序好的切片，遍历 对应 的key value）
}