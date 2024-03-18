package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main()  {
	runtime.GOMAXPROCS(2)

	number := 10
	word := "hello world"

	go typeOfInt(number)
	typeOfStr(word)

	var input string
	fmt.Scanln(&input)

}

func typeOfInt(n int) {
	fmt.Println(reflect.TypeOf(n))
}

func typeOfStr(str string){
	fmt.Println(reflect.TypeOf(str))
}