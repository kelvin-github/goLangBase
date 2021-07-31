package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// http server
func main(){
    http.HandleFunc("/",f1)
    http.HandleFunc("/api",f2)
    http.ListenAndServe("127.0.0.1:9091",nil)
    // http.ListenAndServe("0.0.0.0:9091",nil)  //上服务器时写

}

func f2(w http.ResponseWriter,reve *http.Request){
    receContent := reve.URL.Query()
    fmt.Println(receContent) //get 在这里 收取

    // 根据请求body创建一个json解析器实例
    decoder := json.NewDecoder(reve.Body)   //post 在这里 收取
    fmt.Println(decoder)

    // 处理逻辑

    // 返回数据
    // str := `get it xxxx `
    // w.Write( []byte(str) )
}

func f1(w http.ResponseWriter,r *http.Request){
    // str := `<h1 style="color:red">hello kelvin</h1>`
    // w.Write( []byte(str) )

    b,err:=ioutil.ReadFile("./templete/index.html")
    if err!=nil{
        fmt.Println("read file fail ,err",err)
        w.Write( []byte( fmt.Sprintf("%v",err) ) )        
    }
    w.Write( b )
}