package main

import "fmt"

// 递归函数（术语），就是自己调用自己
// 递归 一定要 设计 明确 的推出条件
// 适合处理 问题相同 ，规模越来越小 的场景

func main(){
    n := uint(5)
    value := f01(n)
    fmt.Println( n," 的阶乘结果",value)

    n1 := 5
    fmt.Println(n1,"级台阶 的 上台阶的走法数：", uplevel(n1) )

}

// 上台阶，一次可以走1步 也可以走2步，有多少种走法(不知道 有多少级台阶) 高中数学 组合：不重复的使用，排列：重复使用
func uplevel(n int) int {
    if n == 1{
        return 1
    }
    if n == 2{
        return 2
    }
    return uplevel(n-1) + uplevel(n-2)
}

// 阶乘 5! = 5*4*3*2*1
func f01(n uint) uint {
    if n<=1 {
        return 1
    }
    return n * f01(n-1)
}