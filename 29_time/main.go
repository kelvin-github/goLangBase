package main

import (
	"fmt"
	"time"
)


func main() {
	// getNow()    //获取当前
    // getTimeStamp()  //时间戳（不同时区，时间戳 不一致）
    timeStamp2Date()    //时间戳 转 日期
}

func timeStamp2Date(){
    var timeStamp01 int64 = 1627784936    //时间戳 秒 10位 （如果是 19位的，截取头10位就好了）
    timeObj01 := time.Unix(timeStamp01,0)
    Y := timeObj01.Year()
    M := timeObj01.Month()
    D := timeObj01.Day()
    
    h := timeObj01.Hour()
    m := timeObj01.Minute()
    s := timeObj01.Second()
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d",Y,M,D,h,m,s)
    
}

func getTimeStamp(){
    now := time.Now()
    timeStamp01 := now.Unix()   //秒数 10位
    timeStamp02 := now.UnixNano()   //纳秒数 19位
    fmt.Println(timeStamp01,timeStamp02)
}

func getNow(){
    now := time.Now()
    fmt.Println(now)

    Y := now.Year()
    M := now.Month()
    D := now.Day()

    h := now.Hour()
    m := now.Minute()
    s := now.Second()
    // fmt.Printf("%d-%02d-%02d %02d:%02d:%02d",Y,M,D,h,m,s)
    fmt.Println("Now:",Y,M,D," ",h,":",m,":",s)
}