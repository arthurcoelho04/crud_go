package main

import (
	"fmt"
	"net/http"

	"crud-go/controller"
	"crud-go/database"
	"crud-go/repository"
	"crud-go/service"
)

func main() {

	// obs// Conecta ao banco de dados
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// obs// Cria o Repository usando a conexão com o banco
	produtoRepository := repository.NewProdutoRepository(db)

	// obs// Cria a Service usando o Repository
	produtoService := service.NewProdutoService(produtoRepository)

	// obs// Cria o Controller usando a Service
	produtoController := controller.NewProdutoController(produtoService)

	// obs// Rotas para listar todos os produtos e cadastrar um produto
	http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			produtoController.Save(w, r)

		case http.MethodGet:
			produtoController.FindAll(w, r)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// obs// Rotas que recebem o ID do produto na URL
	http.HandleFunc("/produtos/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			produtoController.FindByID(w, r)

		case http.MethodPut:
			produtoController.Update(w, r)

		case http.MethodDelete:
			produtoController.Delete(w, r)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando em http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}