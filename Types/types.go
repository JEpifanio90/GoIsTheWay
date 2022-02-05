package Types

import "fmt"

var globalVariable = 10

func init() {
	fmt.Println("This guy gets called like a constructor")
}

func Showcase() {
	fmt.Printf("========== Here's a global Variable %d of type: %T \n", globalVariable, globalVariable)
	intPlayground()
}

func intPlayground() {
	fmt.Println("==================== Int Playground ====================")
	fmt.Println("//////////////////// Unsigned (0 to inf) ///////////////")
	var ubit int8 = 10
	fmt.Printf("8 Bit %T : %v \n", ubit, ubit)
	//var u16V int8 = 10
	//fmt.Printf("8 Bit %T : %d", ubit, ubit)
}

func floatPlayground() {

}
