package main

import "fmt"

/*
50枚金币，这群人分 Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth
规则：
1.名字中 e | E 给1枚
2.名字中 i | I 给2枚
3.名字中 o | O 给3枚
4.名字中 u | U 给4枚
看看每人分多少？分完后，剩下多少？
*/

var (
    coins = 50
    user = []string{
        "Matthew","Sarah","Augustus","Heidi","Emilie","Peter","Giana","Adriano","Aaron","Elizabeth",
    }
    distribution = make(map[string]int,len(user))   //定义一个 map 来存储 每人能分到多少金币
)

func getUserList( u []string ){
    for _,name := range u {
        fmt.Println(name)
        distribution[name] = counter(name)
    }
}
func counter(name string)(int) {
    var cout = 0
    for _,chr := range name {        
        if chr == 'e' || chr == 'E'{
            coins -= 1
            cout += 1
        }
        if chr == 'i' || chr == 'I'{
            coins -= 2
            cout += 2
        }
        if chr == 'o' || chr == 'O'{
            coins -= 3
            cout += 3
        }
        if chr == 'u' || chr == 'U'{
            coins -= 4
            cout += 4
        }
    }
    fmt.Println(name," = ",cout)
    return cout
}


func main(){
    getUserList(user)
    fmt.Println(distribution,coins)

}
