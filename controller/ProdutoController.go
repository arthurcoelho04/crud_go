package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"crud-go/entities"
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

// obs// Recebe os dados do produto e manda salvar POST
func (c *ProdutoController) Save(w http.ResponseWriter, r *http.Request) {
	var produto entities.Produto

	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.service.Save(produto)
	if err != nil {
		http.Error(w, "Erro ao salvar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produto)
}

// obs// Recebe o ID da URL e manda deletar o produto DELETE
func (c *ProdutoController) Delete(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/produtos/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// obs// Busca todos os produtos cadastrados GET
func (c *ProdutoController) FindAll(w http.ResponseWriter, r *http.Request) {
	produtos, err := c.service.FindAll()
	if err != nil {
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produtos)
}

// obs// Busca um produto pelo ID recebido na URL GET
func (c *ProdutoController) FindByID(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/produtos/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	produto, err := c.service.FindByID(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(produto)
}

// obs// Atualiza um produto pelo ID recebido na URL PUT
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
		http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produto)
}