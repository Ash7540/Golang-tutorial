package main

import (
	"fmt"
)

func main(){
	// Arithmetic Operators
	fmt.Println("Arithmetic Operators")

	a := 5
	b := 3
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modules:", a%b)
	a++
	fmt.Println("Increment:", a)
	b--
	fmt.Println("Decrement:", b)
	fmt.Println()


	// Assignment Operators
	fmt.Println("Assignment Operators")
	var x  = 10
	x += 5
	fmt.Println("x += 5:",x)
	x -= 5
	fmt.Println("x -= 5:",x)
	x *= 5
	fmt.Println("x *= 5:",x)
	x /= 5
	fmt.Println("x /= 5:",x)
	x %= 5
	fmt.Println("x %= 5:",x)
	x &= 3
    fmt.Printf("x now is %03b \n", x) // 001
	x |= 3	
    fmt.Printf("x now is %03b \n", x) // 111
	x ^= 3
	fmt.Printf("x now is %03b\n", x) // 100
	x <<= 2
	fmt.Printf("x now is %03b \n", x) // 10000
	x >>= 2
	fmt.Printf("x now is %03b\n", x) // 100

	fmt.Println()
	// Comparison operators
	fmt.Println("Comparison operators")
	fmt.Println("5 == 3:", 5 == 3)
	fmt.Println("5 != 3:", 5 != 3)
	fmt.Println("5 > 3:", 5 > 3)
	fmt.Println("5 < 3:", 5 < 3)
	fmt.Println("5 >= 3:", 5 >= 3)
	fmt.Println("5 <= 3:", 5 <= 3)

	fmt.Println()
	// Logical operators
	fmt.Println("Logical operators")
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)

	fmt.Println()
	// Bitwise Operators
	fmt.Println("Bitwise Operators")
	fmt.Println("5 & 3:", 5 & 3)
	fmt.Println("5 | 3:", 5 | 3)
	fmt.Println("5 ^ 3:", 5 ^ 3)
	fmt.Println("5 &^ 3:", 5 &^ 3)
	fmt.Println("5 << 3:", 5 << 3)
	fmt.Println("5 >> 3:", 5 >> 3)

}