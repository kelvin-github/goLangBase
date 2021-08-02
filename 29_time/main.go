package main

import (
	"fmt"
	"time"
)


func main() {
    // time.Sleep(time.Second*5)   //暂停5秒
    
	// getNow()    //获取当前
    // getTimeStamp()  //时间戳（不同时区，时间戳 不一致）
    // timeStamp2Date()    //时间戳 转 日期
    // Duration()   //时间间隔 常量
    // timeOp() //时间操作
    timeFormat()    //时间格式化
}

// 时间格式化
func timeFormat(){
    timeObj := time.Now()
    // fmt.Println( timeObj.Format("2006-01-02 03:04:05") )    //12小时制
    // fmt.Println( timeObj.Format("2006/01/02 15:04:05.000000000") )    //24小时制
    fmt.Println( timeObj.Format("Mon 2006/Jan/02 03:04:05.000 PM") )    //24小时制
    // fmt.Println( timeObj.Format(time.RFC850))    //常量 格式
}

func timeOp(){ //时间操作 加 （转成 时间对象 来操作）
    var timeStamp01 int64 = 1627784936
    timeObj := time.Unix(timeStamp01,0)
    fmt.Println("timeObj",timeObj)

    time01 := timeObj.Add(time.Hour * 24) // 后1天
    fmt.Println("追加24小时：",time01)
    
    time02 := timeObj.Add(-time.Hour * 24) // 前1天
    fmt.Println("追回24小时：",time02)
    
    // 2个时间 之间 的 差距
    var timeStamp02 int64 = 1527784936
    timeObj02 := time.Unix(timeStamp02,0)
    // d01 := timeObj.Sub(timeObj02) // 计算2个时间 的差值
    d01 := timeObj02.Sub(timeObj) // 计算2个时间 的差值
    fmt.Println("时间差：",d01.Hours()," 小时(正数表示大于；负数表示小于)")

    // 定时器  + 时间比较  Equal Before After
    timer := time.Tick(time.Second)
    laterTime := (time.Now()).Add(time.Second * 30) //x 秒 后 停止
    for t := range timer{        
        if t.After(laterTime)  {  //当 t 在 目标时间后，就打断 循环 
            break
        }
        fmt.Println("定时器：",t,"  -  ",laterTime)   //执行什么工作
    }
}

func Duration(){    //时间间隔
    fmt.Println(time.Nanosecond," 1纳秒")    //常量 
    fmt.Println(time.Microsecond," 1微秒 = 1000 * 1纳秒")    //常量 
    fmt.Println(time.Millisecond," 1毫秒 = 1000 * 1微秒")    //常量
    fmt.Println(time.Second," 1秒 = 1000 * 1毫秒")    //常量 

    fmt.Println(time.Minute,"1分钟 = 60 秒")
    fmt.Println(time.Hour,"1小时 = 60 分钟")
}

func timeStamp2Date(){
    var timeStamp01 int64 = 1627784936    //时间戳 秒 10位
    timeObj01 := time.Unix(timeStamp01,0)
    Y := timeObj01.Year()
    M := timeObj01.Month()
    D := timeObj01.Day()
    
    h := timeObj01.Hour()
    m := timeObj01.Minute()
    s := timeObj01.Second()
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n",Y,M,D,h,m,s)
    
    var timeStamp02 int64 = 1627784936078950000    //时间戳 纳秒 19位
    timeObj02 := time.Unix(0,timeStamp02)
    Y02 := timeObj02.Year()
    M02 := timeObj02.Month()
    D02 := timeObj02.Day()
    
    h02 := timeObj02.Hour()
    m02 := timeObj02.Minute()
    s02 := timeObj02.Second()
    n02 := timeObj02.Nanosecond()
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d:%02d \n",Y02,M02,D02,h02,m02,s02,n02)
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
    // fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",Y,M,D,h,m,s)
    fmt.Println("Now:",Y,M,D," ",h,":",m,":",s)
}