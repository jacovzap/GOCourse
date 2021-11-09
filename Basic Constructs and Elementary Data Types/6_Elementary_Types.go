package main
import (
	"fmt" 
	"math/rand"
)

func main() {
    //var a int
    var b int32
    //a = 15
   // b = a + a // compiler error
    b = b + 5 // ok: 5 is a constant

	var n int16 = 34    // int16 variable
    var m int32         // int32 variable

    m = int32(n)        // explicit typing
    fmt.Printf("32 bit int is: %d\n", m)
    fmt.Printf("16 bit int is: %d\n", n)
	//%d is used for integers
	//%g is used fot float types
	//%s is used for strings
	//%t is used for booleans

	//NUMEROS COMPLEJOS
	var c1 complex64 = 5 + 10i        // Declaring complex num (real +imaginary(¡))
    fmt.Printf("The value is: %v\n", c1)

	//NUMEROS RANDOM
	a := rand.Int()               // generates a random number
    r := rand.Intn(8)             // generates a random number in [0, n)
    fmt.Printf("a is: %d\n", a)
    fmt.Printf("r is: %d\n", r)
}