package exam1

import (
	"fmt"
	"log"
	"net/http"
)

type Price float32

func (p Price) String() string {
	return fmt.Sprintf("$%.2f", p)
}

type Database map[string]Price

func (d *Database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/apple":
		_, _ = fmt.Fprintf(w, "%s: %s\n", r.URL.Path[1:], (*d)["apple"])
	case "/banana":
		_, _ = fmt.Fprintf(w, "banana: %s\n", (*d)["banana"])
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "No page found for: %s\n", r.URL)
	}
}

func Exam() {
	db := &Database{
		"apple":  1.23,
		"banana": 3.12,
	}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
