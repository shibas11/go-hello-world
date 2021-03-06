package exam2

import (
	"fmt"
	"github.com/shibas11/go-hello-world/multiplexer/exam1"
	"log"
	"net/http"
)

type Database map[string]exam1.Price

func (d Database) apple(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "apple: %s\n", d["apple"])
}

func (d Database) banana(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "banana: %s\n", d["banana"])
}

func (d Database) grape(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "baz: %s\n", d["grape"])
}

func Exam() {
	db := &Database{
		"apple":  1.23,
		"banana": 3.12,
		"grape":  9.99,
	}

	mux := http.NewServeMux()
	mux.Handle("/apple", http.HandlerFunc(db.apple))
	mux.Handle("/banana", http.HandlerFunc(db.banana))
	mux.HandleFunc("/grape", db.grape) // convenience method for longer form mux.Handle

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
