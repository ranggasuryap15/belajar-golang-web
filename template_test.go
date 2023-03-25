package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//t, err :=
	//if err != nil {
	//	panic(err)
	//}

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
