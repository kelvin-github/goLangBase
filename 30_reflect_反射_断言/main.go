package main

import (
	"fmt"
	"reflect"
)

type cat struct{
    Name string `json:"name"`
    Age int `json:"age"`
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
    fmt.Printf("t:%v; x的类型 name 是%v; x的类型 kind 是%v \n",t,t.Name(),t.Kind())

    // 断言 判断类型
    switch x.(type){
    case int :
        fmt.Println("int", x)
        break
    case cat :
        fmt.Println("cat", x)
        // 反射 获取结构体 的 字段
        for i:=0; i < t.NumField(); i++{
            field := t.Field(i)
            fmt.Printf("name:%s, index:%d, type:%v, json tag:%v \n",field.Name, field.Index, field.Type, field.Tag.Get("json"))
        }
        break

    // case ptr :
    //     fmt.Println("指针", x)
    //     break
    
    }    
}

func reflectValue(x interface{}){
    v := reflect.ValueOf(x) 
    k := v.Kind()
    fmt.Printf("%v %v",v,k)
}

func reflectValueSet(x interface{}){
    v := reflect.ValueOf(x) 
    k := v.Kind()
    fmt.Printf("%v %v",v,k)
    if k== reflect.Int{
        v.SetInt(200)
    }
    
    
}

func main() {
    // var x01 int =10
    // reflectType(x01)
    // reflectValue(x01)
    
    x02 := cat{
        Name:"cat01",
        Age: 1,
    }
    reflectType(x02)
    // reflectValue(x02)

    // x03 := 10
    // reflectValue(&x03)
    // reflectValueSet(&x03)   //要在 方法中修改值，传指针
    // fmt.Println("修改后：",x03)


}