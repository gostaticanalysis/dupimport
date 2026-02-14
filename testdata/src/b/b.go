package b

import (
	f1 "fmt"
	f2 "fmt" // want "fmt is duplicated import"
)

func main() {
	f1.Println("hello")
	f2.Println("world")
}
