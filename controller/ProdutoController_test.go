package controller

import (
	"bytes"
	"crud-go/database"
	"crud-go/repository"
	"crud-go/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testa se o cadastro de um produto válido retorna HTTP 201.
func TestSaveProdutoValido(t *testing.T) {
	// Conecta ao banco
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("erro ao conectar ao banco: %v", err)
	}

	// Monta as camadas da aplicação
	repo := repository.NewProdutoRepository(db)
	serv := service.NewProdutoService(repo)
	controller := NewProdutoController(serv)

	// Produto usado no teste
	body := []byte(`{"Nome":"Produto Teste","Preco":100}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/produtos",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	// Executa o método Save
	controller.Save(res, req)

	// Verifica o status HTTP
	if res.Code != http.StatusCreated {
		t.Errorf("esperado status 201, recebido %d", res.Code)
	}
}
// Testa se produto sem nome retorna HTTP 400.
func TestSaveProdutoNomeVazio(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("erro ao conectar ao banco: %v", err)
	}

	repo := repository.NewProdutoRepository(db)
	serv := service.NewProdutoService(repo)
	controller := NewProdutoController(serv)

	body := []byte(`{"Nome":"","Preco":100}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/produtos",
		bytes.NewBuffer(body),
	)

	res := httptest.NewRecorder()
	controller.Save(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("esperado status 400, recebido %d", res.Code)
	}
}

// Testa se produto com preço negativo retorna HTTP 400.
func TestSaveProdutoPrecoNegativo(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("erro ao conectar ao banco: %v", err)
	}

	repo := repository.NewProdutoRepository(db)
	serv := service.NewProdutoService(repo)
	controller := NewProdutoController(serv)

	body := []byte(`{"Nome":"Mouse","Preco":-50}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/produtos",
		bytes.NewBuffer(body),
	)

	res := httptest.NewRecorder()
	controller.Save(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("esperado status 400, recebido %d", res.Code)
	}
}
// Testa se a busca de um produto inexistente retorna HTTP 404.
func TestFindByIDProdutoNaoEncontrado(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("erro ao conectar ao banco: %v", err)
	}

	repo := repository.NewProdutoRepository(db)
	serv := service.NewProdutoService(repo)
	controller := NewProdutoController(serv)

	req := httptest.NewRequest(
		http.MethodGet,
		"/produtos/999999",
		nil,
	)

	res := httptest.NewRecorder()

	controller.FindByID(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("esperado status 404, recebido %d", res.Code)
	}
}

// Testa se a atualização de um produto inexistente retorna HTTP 404.
func TestUpdateProdutoNaoEncontrado(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("erro ao conectar ao banco: %v", err)
	}

	repo := repository.NewProdutoRepository(db)
	serv := service.NewProdutoService(repo)
	controller := NewProdutoController(serv)

	body := []byte(`{"Nome":"Teclado","Preco":250}`)

	req := httptest.NewRequest(
		http.MethodPut,
		"/produtos/999999",
		bytes.NewBuffer(body),
	)

	res := httptest.NewRecorder()

	controller.Update(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("esperado status 404, recebido %d", res.Code)
	}
}
// Testa se a exclusão de um produto inexistente retorna HTTP 404.
func TestDeleteProdutoNaoEncontrado(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("erro ao conectar ao banco: %v", err)
	}

	repo := repository.NewProdutoRepository(db)
	serv := service.NewProdutoService(repo)
	controller := NewProdutoController(serv)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/produtos/999999",
		nil,
	)

	res := httptest.NewRecorder()

	controller.Delete(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("esperado status 404, recebido %d", res.Code)
	}
}