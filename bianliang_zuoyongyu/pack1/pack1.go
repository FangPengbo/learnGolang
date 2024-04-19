package pack1

import "fmt"

const privateName = "privateName"
const PublicName = "PublicName"

func privatePack1() {
	fmt.Println(privateName)
}

func PublicPack1() {
	fmt.Println("public pack1")
	privatePack1()
}
