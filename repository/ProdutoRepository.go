package repository

import (
	"crud-go/entities"

	"gorm.io/gorm"
)

type ProdutoRepository struct {
	db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) *ProdutoRepository {
	return &ProdutoRepository{
		db: db,
	}
}

func (r *ProdutoRepository) Save(produto *entities.Produto) error {
	return r.db.Create(produto).Error
}

func (r *ProdutoRepository) FindAll() ([]entities.Produto, error) {
	var produtos []entities.Produto

	err := r.db.Find(&produtos).Error

	return produtos, err
}

func (r *ProdutoRepository) FindByID(id int) (entities.Produto, error) {
	var produto entities.Produto

	err := r.db.First(&produto, id).Error

	return produto, err
}

func (r *ProdutoRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Produto{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Atualiza um produto pelo ID, incluindo campos com valor zero.
func (r *ProdutoRepository) Update(id int, produto entities.Produto) error {
	result := r.db.Model(&entities.Produto{}).
		Where("id = ?", id).
		Select("Nome", "Preco").
		Updates(produto)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}