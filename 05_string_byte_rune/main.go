package main

import (
	"fmt"
	"strings"
	"unicode"
)

// go 中 字符串 一定要用双引号 包裹，单引号 包裹的是 字符（就是1个）
// 1个 字符 = 1个 byte 1Byte = 8Bit,
// 1个 utf8编码 的汉字 = 3Byte = 3 * 8Bit

// 用 \ 来转义 例如 要输出 " 符号 " \" "
// 多行 的字符串 用反引号 ` ` 符号包裹(原样输出)

// len(str) 求长度；
// + 或 fmt.Sprintf 拼接；
// strings.Split 分割；
// strings.contains 判断是否包含；
// strings.HasPrefix 前缀判断
// strings.HasSuffix 后缀判断
// strings.Join(a[]string,sep string) join操作

// rune (int32)类型 是 utf8 字符;  byte (uint8)类型 是 ASCII 码
func main() {
    path := "D:/'中文'//\"abc\"/123 "
    fmt.Println("path:",path)

    s1 := `
        世情薄
        人情恶
        雨送黄昏花易落
    `
    fmt.Print("换行：",s1,"\n")

    // 拼接
    s2 := fmt.Sprintf("%s%s",path,s1)   //带返回值
    fmt.Printf(s2)

    // 分割 放入 数组
    s3 := "1 2 3 4 5 hello小王子"
    s4 := strings.Split(s3," ")
    fmt.Printf("%T\n%v\n%v\n",s4,s4,s4[0])

    // 包含判断
    fmt.Println( strings.Contains(s3,"5") )

    // 前后缀判断
    fmt.Println("判断前缀", strings.HasPrefix(s3,"1") )
    fmt.Println("判断后缀", strings.HasSuffix(s3,"4") )

    // 获取 查询 值 的 index 
    fmt.Println("查询 3 的 索引 位：", strings.Index(s3,"3") )  //空格 也 算 1个索引位置


    // join操作 切片 拼接
    fmt.Println( strings.Join(s4," "))  //s4 拼接回 s4 字符串了

    count01 := 0
    for _,item := range s3{
        // fmt.Printf("item:%c；Type:%T \n",item,item)  //%s字符串 %c 字符
        
        if unicode.Is(unicode.Han,item){
            fmt.Printf("汉字：%v",item)
            count01 ++
        }
    }
    fmt.Println("汉字数：",count01)

    // 字符串 修改
    s5 := "白萝卜"
    s6 := []rune(s5)    //s5 强制转成 了 切片，rune 针对处理 utf8 字符串
    fmt.Printf( "%s,%T\n",string(s6),s6 )

    s6[0] = '红'
    fmt.Printf( "%s,%T\n",string(s6),s6 )
}