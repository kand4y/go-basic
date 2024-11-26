package main

import (
	"fmt"
	"net/http"
	"os"
	"unsafe"

	"github.com/joho/godotenv"
)

const the_readonly = "omg"

type fruits int

const (
	banana fruits = iota + 1
	apple
	mango
)

func main() {
	fmt.Println("Hello, World")

	// fizz_buzz()
	// test_env()

	var i int = 4
	fmt.Printf("i: %v %T\n", i, i)

	fmt.Printf("banana:%v apple:%v mango:%v \n", banana, apple, mango)

	// ポインタ変数
	var ui1 uint16
	fmt.Printf("value of ui1: %v \n", ui1)

	var p1 *uint16
	fmt.Printf("value of p1: %v \n", p1)

	p1 = &ui1
	fmt.Printf("value of p1: %v \n", p1)
	fmt.Printf("size of p1: %d[byte] \n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p \n", &p1)

	fmt.Printf("value of ui1(dereference): %v \n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v \n", ui1)
	//simple_http()
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
