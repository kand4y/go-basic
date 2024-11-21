package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World")

	// fizz_buzz()
	// test_env()

	var i int = 4
	fmt.Printf("i: %v %T\n", i, i)

	simple_http()
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

func simple_http() {
	req := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello Go\n")
	}

	http.HandleFunc("/", req)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
