package handlers

import (
	"net/http"
	"fmt"
	"log"
	"io"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	
	h.l.Println("Hello World - From Logging")

	d, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}