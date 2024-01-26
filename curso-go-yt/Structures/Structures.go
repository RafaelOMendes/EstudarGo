package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo metodo:", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Pais string
}

func main() {
	cliente := Cliente{
		Nome:  "Rafa",
		Email: "w@w.com",
		CPF:   04765526054,
	}
	fmt.Println(cliente)

	cliente2 := Cliente{"Gabriel", "m@m.com", 12345678900}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "d@d.com",
			CPF:   12312312300,
		},
		Pais: "EUA",
	}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d. Pais: %s\n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)

	cliente.Exibe()
	cliente3.Exibe()

	clienteJson, err := json.Marshal(cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))

	jsonCliente4 := `{"Nome":"Davi","Email":"d@d.com","CPF":12312312300,"Pais":"EUA"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}
