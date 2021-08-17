package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("err:",resp.StatusCode)
		return
	}

	// resp.Body 如果 这个网页 事 GBK 的 要变成UTF8
	// utf8Reader := transform.NewReader(resp.Body , simplifiedchinese.GBK.NewDecoder() )
	utf8Reader := resp.Body

	// 获取到 网页源码
	all,err:=ioutil.ReadAll(utf8Reader)	
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n", all )
}

// 自动获取网页中的编码
func determinEncoding(r io.Reader) encoding.Encoding{
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err!=nil{
		panic(err)
	}
	ec,_,_ := charset.DetermineEncoding(bytes, "")
	return ec
}