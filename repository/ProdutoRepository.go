package repository

import "crud-go/entities"

type ProdutoRepository struct {
	produtos []entities.Produto
}

func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{
		produtos: []entities.Produto{},
	}
}

func (r *ProdutoRepository) Save(produto entities.Produto) {
	r.produtos = append(r.produtos, produto)
}

func (r *ProdutoRepository) FindAll() []entities.Produto { //retorna uma lista de produtos
	return r.produtos
}
