package d

import (
	"fmt"
	fmt2 "fmt" // want "fmt is duplicated import"
	fmt3 "fmt" // want "fmt is duplicated import"
)

func main() {
	fmt.Println("one")
	fmt2.Println("two")
	fmt3.Println("three")
}
