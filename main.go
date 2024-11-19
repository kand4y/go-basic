package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World")

	fizz_buzz()
}

func fizz_buzz() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz!!")
		} else if i%3 == 0 {
			fmt.Println("Fizz!!")
		} else if i%5 == 0 {
			fmt.Println("Buzz!!")
		} else {
			fmt.Println(i)
		}
	}
}

func test_env() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
}
