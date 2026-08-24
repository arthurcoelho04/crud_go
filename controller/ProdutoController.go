package controller

import "crud-go/service"

type ProdutoController struct {
	service *service.ProdutoService
}

func NewProdutoController(service *service.ProdutoService) *ProdutoController { // reebe um produto service
	return &ProdutoController{
		service: service,
	}
}
