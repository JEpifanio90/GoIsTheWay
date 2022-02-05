package Types

import "fmt"

var globalVariable = 10

func init() {
	fmt.Println("This guy gets called like a constructor")
}

func Showcase() {
	fmt.Printf("========== Here's a global Variable %d of type: %T \n", globalVariable, globalVariable)
	nonExported()
}

func nonExported() {
	fmt.Println("This guy gets discriminated by it's own name")
}
