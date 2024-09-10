package repository

import (
	"github.com/affrianr/gin-rest-api-h8/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id uint) (*domain.Order, error)
	Update(order *domain.Order) error
	Delete(id uint) error
	List() ([]domain.Order, error) 
}

type postgresOrderRepository struct {
    db *gorm.DB
}

func NewPostgresOrderRepository(db *gorm.DB) OrderRepository {
    return &postgresOrderRepository{db}
}

func (r *postgresOrderRepository) Create(order *domain.Order) error {
    return r.db.Create(order).Error
}

func (r *postgresOrderRepository) GetByID(id uint) (*domain.Order, error) {
    var order domain.Order
    err := r.db.Preload("Items").First(&order, id).Error
    return &order, err
}

func (r *postgresOrderRepository) Update(order *domain.Order) error {
    return r.db.Save(order).Error
}

func (r *postgresOrderRepository) Delete(id uint) error {
    return r.db.Delete(&domain.Order{}, id).Error
}

func (r *postgresOrderRepository) List() ([]domain.Order, error) {
    var orders []domain.Order
    err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

