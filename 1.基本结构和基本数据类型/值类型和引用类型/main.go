package main

import "fmt"

//基础类型是值类型 int string bool 数组和结构
//指针 slices map channel 是引用类型

func main() {

	a := 5
	b := false

	fmt.Println(a, b)
	c := a
	d := b
	c = 10
	d = true

	fmt.Println(a, b, c, d)

	type User struct {
		Name string
	}

	xiaoming := User{Name: "xiaoming"}
	fmt.Println(xiaoming)

	xiaohong := xiaoming
	xiaohong.Name = "xiaohong"
	fmt.Println(xiaoming, xiaohong)

	maps := map[string]string{}
	maps["a"] = "a1"
	maps["b"] = "b1"
	maps["c"] = "c1"

	fmt.Println(maps)

	maps2 := maps
	maps2["a"] = "a2"

	fmt.Println(maps, maps2)

}
