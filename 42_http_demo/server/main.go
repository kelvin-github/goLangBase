package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http server
func main(){
    http.HandleFunc("/path01/01/",f1)
    http.ListenAndServe("127.0.0.1:9091",nil)
    // http.ListenAndServe("0.0.0.0:9091",nil)  //上服务器时写

}

func f1(w http.ResponseWriter,r *http.Request){
    // str := `<h1 style="color:red">hello kelvin</h1>`
    // w.Write( []byte(str) )

    b,err:=ioutil.ReadFile("./templete/test.html")
    if err!=nil{
        fmt.Println("read file fail ,err",err)
        w.Write( []byte( fmt.Sprintf("%v",err) ) )        
    }
    w.Write( b )

}