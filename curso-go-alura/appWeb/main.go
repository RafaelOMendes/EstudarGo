package main

import (
	"html/template"
	"net/http"

	"github.com/RafaelOMendes/EstudarGo/curso-go-alura/appWeb/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SelectAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
