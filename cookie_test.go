package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// menyimpan datanya bukan di server tapi di client namun jangan diterima bulat-bulat, harus aman datanya.
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Generus-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success Create Cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Generus-Name")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
