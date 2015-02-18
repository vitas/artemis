package main

import (
	"fmt"
	. "github.com/vitas/artemis/aesf"
)

func main() {
	fmt.Println("Init World!")
	pWorld := new(World)
	pWorld.Initialize()

}
