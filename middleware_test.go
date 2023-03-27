package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute Handler") // kita bisa melakukan sesuatu sebelum dan sesudah memanggil handler
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(writer, "Hello Middleware")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
