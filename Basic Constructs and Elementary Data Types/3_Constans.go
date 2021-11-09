package main
import "fmt"

func main(){
	//convencion de escribir todas las constantes con mayusculas
	const PI = 3.14159
	const pi float32 = 3.14159
	var a string = "5"
	fmt.Println(a)
	fmt.Println(pi)


	const BEEF, TWO, C = "meat", 2, "veg"
	fmt.Println(BEEF)
	fmt.Println(TWO)

	//ENUMERATIONS
	const (
		UNKNOW = iota
		FEMALE = iota
		MALE = iota
	)
	fmt.Println(UNKNOW)
	fmt.Println(FEMALE)
	fmt.Println(MALE)
}