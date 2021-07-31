package main

import (
	"fmt"
	"io/ioutil"
)

func readFromFileByIoutil(){
    ret,err:=ioutil.ReadFile("./logs.txt")
    if err!=nil{
        fmt.Println("open file fail ,err",err)
        return
    }
    fmt.Println( string(ret) )

}
func main() {
	readFromFileByIoutil()
}