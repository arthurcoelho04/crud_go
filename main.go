package main

import (
	"fmt"
	"net/http"

	"crud-go/controller"
	"crud-go/repository"
	"crud-go/service"
)

func main() {

	// Cria o Repository
	produtoRepository := repository.NewProdutoRepository()

	// Cria a Service usando o Repository
	produtoService := service.NewProdutoService(produtoRepository)

	// Cria o Controller usando a Service
	produtoController := controller.NewProdutoController(produtoService)

	// Rota de produtos
	http.HandleFunc("/produtos", produtoController.Delete)

	fmt.Println("Servidor rodando em http://localhost:8080")

	http.ListenAndServe(":8080", nil)

	_ = produtoController
}
