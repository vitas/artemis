package main

import (
	"fmt"
	. "github.com/vitas/artemis/aesf"
)

func main() {
	world := NewEntityWorld()
	fmt.Println(&world)
}
