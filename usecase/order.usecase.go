package usecase

import (
	"github.com/affrianr/gin-rest-api-h8/domain"
	"github.com/affrianr/gin-rest-api-h8/repository"
)

type OrderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUseCase(repo repository.OrderRepository) *OrderUsecase {
    return &OrderUsecase{repo}
}

func (uc *OrderUsecase) CreateOrder(order *domain.Order) error {
    return uc.repo.Create(order)
}

func (uc *OrderUsecase) GetOrder(id uint) (*domain.Order, error) {
    return uc.repo.GetByID(id)
}

func (uc *OrderUsecase) UpdateOrder(order *domain.Order) error {
    return uc.repo.Update(order)
}

func (uc *OrderUsecase) DeleteOrder(id uint) error {
    return uc.repo.Delete(id)
}

func (uc *OrderUsecase) ListOrders() ([]domain.Order, error) {
    return uc.repo.List()
}