package main

import (
	"fmt"
	"runtime"
)

func main() {
	echo("Hello World!")
}

func echo(words string) {
	fmt.Println(words, runtime.GOOS)
}
