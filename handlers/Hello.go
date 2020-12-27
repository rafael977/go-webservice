package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello define dependencies
type Hello struct {
	l *log.Logger
}

// NewHello create instance
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHttp serve http
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello wolrd")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Error  data", http.StatusBadRequest)
		return
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Error data"))
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
