package f

import (
	"net/http"
	h "net/http" // want "net/http is duplicated import"
)

func main() {
	http.DefaultClient.Get("https://example.com")
	h.DefaultClient.Get("https://example.com")
}
