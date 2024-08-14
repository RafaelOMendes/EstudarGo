package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	conect := "user=postgres dbname=alura_loja password=0418 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conect)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectDB()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	selectProducts, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
