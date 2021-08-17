package regexp

import (
	"fmt"
	"regexp"
)

const test = `
my email is 530979104@qq.com
email is aaabc@126.com
email is 123@gmail.com
email is bbb@hotmail.com.cn
`

func Test() {
	// res := regexp.MustCompile(`[A-Za-z0-9]+@[A-Za-z0-9]+\.[A-Za-z0-9]+`)   //匹配邮箱(\. 是对 . 符号 转义)
    // match := res.FindString(test)    //首个匹配
    // match := res.FindAllString(test,-1)  //多个匹配 ,返回 的是 切片

    res := regexp.MustCompile(`([A-Za-z0-9]+)@([A-Za-z0-9]+)\.([A-Za-z0-9.]+)+`)   //匹配邮箱(\. 是对 . 符号 转义)
    match := res.FindAllStringSubmatch(test,-1) // () 再次 子匹配 

    // fmt.Printf("%T,%s",match,match)
    for _,m:= range match{
        fmt.Println(m)
    }
}
