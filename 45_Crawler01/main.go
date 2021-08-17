package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	// fmt.Printf("%s\n", all )
	printCityAll(all)


}

// 解析 城市列表
func printCityAll(content []byte){
	// res := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+"[^>]*>[^<]+</a>`)	//[^>]* 不是>符号的所有个; [^<]+不是< 的整个
	res := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)	//[^>]* 不是>符号的所有个; [^<]+不是< 的整个
	// matcher := res.FindAll(content, -1)
	matcher := res.FindAllSubmatch(content, -1)
	// fmt.Println(matcher)

	for _,m := range matcher{
		// fmt.Printf("%s \n",m)
		fmt.Printf("URL:%s , AD:%s\n",m[1],m[2])
	}
}
