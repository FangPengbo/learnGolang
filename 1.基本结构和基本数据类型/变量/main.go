package main

import "fmt"

// var 关键字应当是在定义全局变量时使用 内部变量应当使用 := 语法声明

var globalNumber = 100

func main() {

	//定义内部变量
	countNumber := 20

	fmt.Println(globalNumber, countNumber)

}

// 使用全局变量
func func1() {

	fmt.Println(globalNumber)

}
