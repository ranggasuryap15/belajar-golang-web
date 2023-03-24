package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// parsing dulu baru ambil post form
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	// request.PostFormValue("first_name")

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")
	
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Rangga&last_name=Prayoga")
	request := httptest.NewRequest("POST", "http://localhost/", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded") // ini sudah menjadi standar content type-nya
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	
}