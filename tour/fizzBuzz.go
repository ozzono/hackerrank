package main

import "fmt"

func main() {
	fizzBuzz(int32(15))
}

func fizzBuzz(n int32) {
	for i := 1; i <= int(n); i++ {
		mod3 := (i % 3) == 0
		mod5 := (i % 5) == 0
		switch true {
		case mod3 && mod5:
			fmt.Println("FizzBuzz")
		case mod3 && !mod5:
			fmt.Println("Fizz")
		case !mod3 && mod5:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
