package service

import (
	"crud-go/entities"
	"errors"
	"testing"
)

// Testa se um produto sem nome é rejeitado.
func TestValidarProdutoNomeVazio(t *testing.T) {
	produto := entities.Produto{
		Nome:  "",
		Preco: 100,
	}

	err := validarProduto(produto)

	if !errors.Is(err, ErrNomeObrigatorio) {
		t.Errorf("esperado ErrNomeObrigatorio, recebido: %v", err)
	}
}

// Testa se um produto com preço negativo é rejeitado.
func TestValidarProdutoPrecoNegativo(t *testing.T) {
	produto := entities.Produto{
		Nome:  "Mouse",
		Preco: -50,
	}

	err := validarProduto(produto)

	if !errors.Is(err, ErrPrecoInvalido) {
		t.Errorf("esperado ErrPrecoInvalido, recebido: %v", err)
	}
}

// Testa se um produto válido passa pela validação sem erro.
func TestValidarProdutoValido(t *testing.T) {
	produto := entities.Produto{
		Nome:  "Teclado",
		Preco: 200,
	}

	err := validarProduto(produto)

	if err != nil {
		t.Errorf("não era esperado erro, recebido: %v", err)
	}
}