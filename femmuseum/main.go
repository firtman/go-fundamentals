package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"frontendmuseum.com/api"
	"frontendmuseum.com/data"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello, HTTP!\n")
}

func getTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.template")
	if err != nil {
		fmt.Printf("Template error: %v", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = html.Execute(w, data.GetAll())
	if err != nil {
		return
	}
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/hello", getHello)
	server.HandleFunc("/templates", getTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err == nil {
		print("Error opening the server")
	}
}
