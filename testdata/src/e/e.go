package e

import (
	"fmt"
	f "fmt" // want "fmt is duplicated import"
	"strings"
	s "strings" // want "strings is duplicated import"
)

func main() {
	fmt.Println("hello")
	f.Println("world")
	strings.Contains("hello world", "hello")
	s.Contains("hello world", "world")
}
