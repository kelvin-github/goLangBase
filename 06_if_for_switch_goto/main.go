package main

import "fmt"

func main() {
	var a1 int = 10
	if a1 > 5 {
		fmt.Println("大于5")
	}else if a1<10{
        fmt.Println("小于10")
    }else{
        fmt.Println("a1=",a1)
    }

    // 无限循环
    // for {
    //     a1++
    //     fmt.Println("hi", a1 )
    // }

    // for 的变种，a1 赋值 作用域 有所不同
    for a1:=15;a1<20; a1++{        
        fmt.Println(a1)
        a1++        
    }
    fmt.Println("a1:",a1)

    s1 := "hello world"
    for index,item :=  range s1 {
        fmt.Printf("索引：%v, AscII值：%v , 字符值：%v \n",index,item,string(item) )
    }

    // var flag = false
    for i:=0;i<10;i++{
        for j:='A';j<'Z';j++{
            if j=='C'{
                // flag = true
                break
            }
            fmt.Printf("%v--%c \n",i,j)
        }
        // if flag{
        //     break
        // }        
    }

    for i:=0;i<10;i++{
        if i==5{            
            // break   //达到当前条件，就整体 退出循环；
            continue    //跳过当前条件，继续下一次循环
        }
        fmt.Println("i == 5 ",i)
    }

    // switch
    for i:=0;i<5;i++{
        switch i{
            case 1,2,3 :
                fmt.Println("case A:",i)
            
            case 4:
                fmt.Println("case B:",i)
            default:
                fmt.Println("case C:",i)
        }
           
    }

    

    // goto
    for i:=0;i<10;i++{
        for j:='A';j<'Z';j++{
            if j=='C'{
                goto x01    //跳去 标签 处，执行标签处的代码，同时 也打断了当前的 for嵌套
            }
            fmt.Printf("%v--%c \n",i,j)
        }
    }

    x01:goto01()

    fmt.Print("hi ?")
   

}

func goto01(){
    var x1 = "x01"    
    fmt.Println("label = ",x1)    
}