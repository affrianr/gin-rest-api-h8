package usecase

import (
	"github.com/affrianr/gin-rest-api-h8/domain"
)

type ItemUsecase struct {
	repo domain.ItemRepository
}

func NewItemUseCase(repo domain.ItemRepository) *ItemUsecase {
    return &ItemUsecase{repo}
}

func (uc *ItemUsecase) CreateItem(item *domain.Item) error {
    return uc.repo.Create(item)
}

func (uc *ItemUsecase) GetItem(id uint) (*domain.Item, error) {
    return uc.repo.GetByID(id)
}

func (uc *ItemUsecase) UpdateItem(item *domain.Item) error {
    return uc.repo.Update(item)
}

func (uc *ItemUsecase) DeleteItem(id uint) error {
    return uc.repo.Delete(id)
}

func (uc *ItemUsecase) ListItems() ([]domain.Item, error) {
    return uc.repo.List()
}