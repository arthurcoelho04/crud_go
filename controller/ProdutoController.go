package controller

import (
	"net/http"

	"crud-go/service"
)

type ProdutoController struct {
	service *service.ProdutoService
}

func NewProdutoController(service *service.ProdutoService) *ProdutoController { // reebe um produto service "injesão de dependencia"
	return &ProdutoController{
		service: service,
	}
}

func (c *ProdutoController) Delete(w http.ResponseWriter, r *http.Request) {
	// futuramente vamos pegar o ID da URL

	c.service.Delete(1)

	w.WriteHeader(http.StatusNoContent)
}
