package c

import (
	myfmt "fmt"
	"fmt" // want "fmt is duplicated import"
)

func main() {
	myfmt.Println("hello")
	fmt.Println("world")
}
