package main

import (
	"encoding/json"
	"fmt"
)

// 1. 把 结构体 转 json 格式的字符串(序列号)
// 2. 把 json 格式 的字符串 转 结构体（反序列化）

type Persion struct{
    Name string `json:"username" db:"name" ini:"name"`
    Age int `json:"age"`
}
func main() {
    // 指针 声明 并 初始化（赋值）
    p1 := &Persion{
        Name: "kelvin",
        Age: 42,
    }

    // 结构体 转 json
    b,err:= json.Marshal(p1)
    if err!=nil{
        fmt.Println("marshal err",err)
        return
    }
    fmt.Println( string(b) )
    fmt.Printf( "%v \n", string(b) )

    // json 转 结构体
    str := string(b)
    // var p2 Persion  //接收者 要传 指针地址
    // json.Unmarshal( []byte(str), &p2)    
    var p2 = new(Persion)   // new 来声明，得到的就是 指针地址
    json.Unmarshal( []byte(str), p2)
    fmt.Printf("json2struct: %v, %T",p2,p2)
}