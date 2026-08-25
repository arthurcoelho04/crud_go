package main

import (
	"crud-go/controller"
	"crud-go/entities"
	"crud-go/repository"
	"crud-go/service"
	"fmt"
)

func main() {

	produtoRepository := repository.NewProdutoRepository()

	produtoService := service.NewProdutoService(produtoRepository)

	produtoController := controller.NewProdutoController(produtoService)

	_ = produtoController

	// Criando um produto
	produto := entities.Produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 3500,
	}

	// Salvando o produto
	produtoService.Save(produto)

	// Buscando o produto pelo ID
	produtoEncontrado, err := produtoService.FindByID(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(produtoEncontrado)
	}
}
