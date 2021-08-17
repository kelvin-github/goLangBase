package fetch

import (
	"bufio"
	"encoding"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
    if resp.StatusCode!= http.StatusOK{
        return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
    }
    e := determinEncoding(resp.Body)
    utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
    return ioutl.ReadAll(utf8Reader)
}

// 自动获取网页中的编码
func determinEncoding(r io.Reader) encoding.Encoding{
    bytes,err := bufio.NewReader(r).Peek(1024)
    if err!=nil{
        log.Printf("fetcher err %v", err)
        return unicode.UTF8
    }
    e,_,_ := charset.DetermineEncoding(bytes,"")
    return e
}
