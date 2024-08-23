package main

// import (
// 	"net/http"

// 	"loja/routes"
// )

import (
	// "database/sql"
	"fmt"
	"log"
	"loja/routes"

	// "html/template"
	"net/http"
	// _ "github.com/lib/pq"
)

// func conectaComBancoDeDados() *sql.DB {
// 	conexao := "user=postgres dbname=loja password=dba1234 host=localhost sslmode=disable"
// 	db, err := sql.Open("postgres", conexao)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// type Produto struct {
// 	//	Id         int
// 	Nome       string
// 	Descricao  string
// 	Preco      float64
// 	Quantidade int
// }

// var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	sistema := "Loja"
	versao := 1.1
	fmt.Println("Sistema: ", sistema)
	fmt.Println("Versão: ", versao)
	routes.CarregaRotas()
	fmt.Println("Carregou rotas....")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("Finalizou.... ")

}

// func main() {
// 	sistema := "Loja"
// 	versao := 1.1
// 	fmt.Println("Sistema: ", sistema)
// 	fmt.Println("Versão: ", versao)

// 	fmt.Println("Conecta DB.....")
// 	conexao := "user=postgres dbname=loja password=dba1234 host=localhost sslmode=disable"
// 	db, err := sql.Open("postgres", conexao)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	fmt.Println("Aitiva HTTP Server.....")

// 	http.HandleFunc("/", index)
// 	http.ListenAndServe(":8080", nil)
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	db := conectaComBancoDeDados()

// 	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	p := Produto{}
// 	produtos := []Produto{}

// 	for selectDeTodosOsProdutos.Next() {
// 		var id, quantidade int
// 		var nome, descricao string
// 		var preco float64

// 		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		p.Nome = nome
// 		p.Descricao = descricao
// 		p.Preco = preco
// 		p.Quantidade = quantidade

// 		produtos = append(produtos, p)
// 	}

// 	temp.ExecuteTemplate(w, "Index", produtos)
// 	defer db.Close()
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	produtos := []Produto{
// 		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
// 		{"Tenis", "Confortável", 89, 3},
// 		{"Fone", "Muito bom", 59, 2},
// 		{"Produto novo", "Muito legal", 1.99, 1},
// 	}

// 	temp.ExecuteTemplate(w, "Index", produtos)
// }
