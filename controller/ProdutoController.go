package controller

import (
	"net/http"
	"strconv"
	"strings"

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
	idString := strings.TrimPrefix(r.URL.Path, "/produtos/")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	c.service.Delete(id)//delete por id

	w.WriteHeader(http.StatusNoContent)
}
