package a

import (
	"fmt"
	fmt2 "fmt" // want "fmt is duplicated import"
)

func main() {
	fmt.Println("hoge")
	fmt2.Println("hoge")
}
