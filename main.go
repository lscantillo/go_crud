package main

import (
	"log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Home)
	log.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola")
	plantillas.ExecuteTemplate(w, "home", nil)
}
