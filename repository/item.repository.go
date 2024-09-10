package repository

import (
	"github.com/affrianr/gin-rest-api-h8/domain"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *domain.Item) error
	GetByID(id uint) (*domain.Item, error)
	Update(item *domain.Item) error
	Delete(id uint) error
	List() ([]domain.Item, error) 
}

type postgresItemRepository struct {
    db *gorm.DB
}

func NewPostgresItemRepository(db *gorm.DB) ItemRepository {
    return &postgresItemRepository{db}
}

func (r *postgresItemRepository) Create(item *domain.Item) error {
    return r.db.Create(item).Error
}

func (r *postgresItemRepository) GetByID(id uint) (*domain.Item, error) {
    var item domain.Item
    err := r.db.First(&item, id).Error
    return &item, err
}

func (r *postgresItemRepository) Update(item *domain.Item) error {
    return r.db.Save(item).Error
}

func (r *postgresItemRepository) Delete(id uint) error {
    return r.db.Delete(&domain.Item{}, id).Error
}

func (r *postgresItemRepository) List() ([]domain.Item, error) {
    var items []domain.Item
    err := r.db.Find(&items).Error
	return items, err
}

