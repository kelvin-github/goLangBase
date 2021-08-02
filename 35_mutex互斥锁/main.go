package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁
// 共享 公共数据  操作 时 常见 存取时机 交叉出现 偏差
// 互斥锁 会 导致 性能问题

// 读锁，锁了，其他人都可以读，但不能写
// 写锁，锁了，其他人不可以写，也不可以读

// 读写 分离，写 主库， 读 从库

var(
    wg sync.WaitGroup
    x = 0
    lock sync.Mutex //声明 一把 互斥锁
    lockrw sync.RWMutex //声明 一把 读写互斥锁
)

func writer01(){
    defer wg.Done()    
    lock.Lock() //上锁（互斥锁）
    x = x + 1  
    // time.Sleep(time.Millisecond*100)    //100毫秒
    lock.Unlock()  //解锁（互斥锁）
}

func reader01(){
    defer wg.Done()
    // lock.Lock()  //这种锁，别人 不能读，不能写 耗时
    // fmt.Println("当前x : ",x)
    // lock.Unlock()   
    
    lockrw.RLock() //读锁 上 （别人 可以读，但 不能写）
    fmt.Println("当前x : ",x)
    lockrw.RUnlock()  //读锁 解
}

func main(){
    start := time.Now()
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go writer01()
    }
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go reader01()
    }
    
    wg.Wait()
    Dt := time.Now().Sub(start)
    fmt.Println("main still alive! ，耗时：", Dt.Seconds() ,"秒")
}