package repository

import (
	"crud-go/entities"
	"errors"
)

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

func (r *ProdutoRepository) FindByID(id int) (entities.Produto, error) {

	for _, produto := range r.produtos {

		if produto.ID == id {
			return produto, nil
		}
	}

	return entities.Produto{}, errors.New("produto não encontrado")
}

func (r *ProdutoRepository) Update(id int, produtoAtualizado entities.Produto) (entities.Produto, error) {

	for i, produto := range r.produtos {

		if produto.ID == id {

			r.produtos[i] = produtoAtualizado

			return produtoAtualizado, nil
		}
	}

	return entities.Produto{}, errors.New("produto não encontrado")
}

func (r *ProdutoRepository) Delete(id int) { //Ele percorre a lista procurando o produto com depois deleta e finaliza o metodo
	for i, produto := range r.produtos {
		if produto.ID == id {
			r.produtos = append(r.produtos[:i], r.produtos[i+1:]...)
			return
		}
	}
}
