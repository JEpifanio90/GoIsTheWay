package playground

import "fmt"

var x int
var y string
var z bool

func init() {
	fmt.Println("Setup...")
}

func RunExample(number uint8) {
	switch number {
	case 1:
		example1()
	case 2:
		example2()
	default:
		fmt.Println("Yo! We only support 1 through 5")
	}
}

// Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// Now print the values stored in those variables using a single print statement multiple print statements
func example1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
}

// Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to
// the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following
// TYPE (meaning they can store VALUES of that TYPE).
func example2() {
	fmt.Println(x, y, z)
}
