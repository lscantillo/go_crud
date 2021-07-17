package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	Name := "go_crud"
	conexion, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/crear", Create)
	log.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	establishConexion := conexionDB()
	insertRecords, err := establishConexion.Prepare("INSERT INTO employess(name,email) VALUES ('Oscar','correo@gmail.com')")
	if err != nil {
		panic(err.Error())
	}
	insertRecords.Exec()
	plantillas.ExecuteTemplate(w, "home", nil)
}
func Create(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola")
	plantillas.ExecuteTemplate(w, "create", nil)
}
