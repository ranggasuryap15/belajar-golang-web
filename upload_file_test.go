package belajar_golang_web

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload_form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
