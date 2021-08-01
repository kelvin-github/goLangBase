package main

import (
	"fmt"
	"strconv"
)


const (
    // RYB = 111
	R int = 1  //二进制：001
	Y int = 2  //二进制：010
	B int = 4  //二进制：100
)

func getColorMax(arg int) {
	max := fmt.Sprintf("%b",arg)    //十进制 转 二进制（字符串）
    
    switch max {
    case "001":
        fmt.Println("red")
        break
    case "010":
        fmt.Println("yellow")
        break
    case "100":
        fmt.Println("blue")
        break
    case "011":
        fmt.Println("red | yellow")
        break
    case "111":
        fmt.Println("red | yellow | blue")
        break
    case "110":
        fmt.Println("yellow | blue")
        break
    }

    switch arg {
    case 1: //001
        fmt.Println("red")
        break
    case 2: //010
        fmt.Println("yellow")
        break
    case 4: //100
        fmt.Println("blue")
        break
    case 3: //011
        fmt.Println("red | yellow")
        break
    case 7: //111
        fmt.Println("red | yellow | blue")
        break
    case 6: //110
        fmt.Println("yellow | blue")
        break
    }

    fmt.Println([]byte(strconv.Itoa(arg))," ",max)
}

func main() {
    getColorMax( Y | B)
}