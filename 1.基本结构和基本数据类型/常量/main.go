package main

import "fmt"

// Pi 隐式定义
const Pi = 3.1415926

// Pi2 显示定义
const Pi2 float64 = 3.1415926

// 多行定义
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

// 枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

// iota是一个自增值，从0开始，每在新的一行使用时值都会增加一，并且没有赋值的常量默认会应用上一行的赋值表达式
// 在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0
const (
	a = iota
	b = iota
	c
	d
	e = iota + 50
)

func main() {
	fmt.Println(a, b, c, d, e)
}
