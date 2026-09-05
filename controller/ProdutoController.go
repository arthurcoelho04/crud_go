package controller

import (
	"crud-go/entities"
	"crud-go/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type ProdutoController struct {
	service *service.ProdutoService
}

func NewProdutoController(service *service.ProdutoService) *ProdutoController {
	return &ProdutoController{
		service: service,
	}
}

func (c *ProdutoController) Save(w http.ResponseWriter, r *http.Request) {
	var produto entities.Produto

	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.service.Save(produto)
	if err != nil {
		if errors.Is(err, service.ErrNomeObrigatorio) ||
			errors.Is(err, service.ErrPrecoInvalido) {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "Erro interno ao salvar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(produto)
}

func (c *ProdutoController) FindAll(w http.ResponseWriter, r *http.Request) {
	produtos, err := c.service.FindAll()
	if err != nil {
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produtos)
}

func (c *ProdutoController) FindByID(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/produtos/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	produto, err := c.service.FindByID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar produto", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produto)
}

func (c *ProdutoController) Update(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/produtos/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var produto entities.Produto

	err = json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.service.Update(id, produto)
	if err != nil {
		if errors.Is(err, service.ErrNomeObrigatorio) ||
			errors.Is(err, service.ErrPrecoInvalido) {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "Erro interno ao atualizar produto", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produto)
}

func (c *ProdutoController) Delete(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/produtos/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		http.Error(w, "Erro ao excluir produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}