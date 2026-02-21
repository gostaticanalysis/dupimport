package g

import (
	"fmt"
	f "fmt" // want "fmt is duplicated import"
)

func main() {
	fmt.Println("used")
	_ = f.Sprintf("unused")
}
