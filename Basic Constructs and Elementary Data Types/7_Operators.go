package main
import (
	"fmt"
)

func main(){
   a:= 5
   b := 3
   b += a
   fmt.Printf("La suma de a + b es %d\n",b)

   c := b > a
   fmt.Printf("b es mayor que a: %t\n",c)
	c = !c
   fmt.Printf("b es mayor que a: %t\n",c)

}