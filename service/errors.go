package service

import "errors"

var (
	ErrNomeObrigatorio       = errors.New("nome do produto é obrigatório")
	ErrPrecoInvalido         = errors.New("preço do produto não pode ser negativo")
	ErrProdutoNaoEncontrado  = errors.New("produto não encontrado")
)