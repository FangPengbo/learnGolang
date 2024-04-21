package pack1

import (
	"fmt"
	"strconv"
)

var p1 int

func init() {
	fmt.Println("pck1 init")
	p1 = 10
}

func Pack1() {
	fmt.Println("pack1 p1:" + strconv.Itoa(p1))
}
