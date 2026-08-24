package main

import (
	"crud-go/controller"
	"crud-go/repository"
	"crud-go/service"
)

func main() {

	produtoRepository := repository.NewProdutoRepository()

	produtoService := service.NewProdutoService(produtoRepository)

	produtoController := controller.NewProdutoController(produtoService)

	_ = produtoController
}
