package pack2

import (
	p1 "bianliang_zuoyongyu/pack1"
	"fmt"
)

func Execute() {
	fmt.Println("pack2 pack1 name :", p1.PublicName)
	p1.PublicPack1()
}
