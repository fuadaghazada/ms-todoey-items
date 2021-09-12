package service

import (
	"github.com/fuadaghazada/ms-todoey-items/dao/repo"
	"github.com/fuadaghazada/ms-todoey-items/model"
)

type IItemService interface {
	GetUserItems(userID string) (*[]model.ItemDto, error)
}

type itemService struct {
	itemRepo repo.IItemRepository
}

func NewItemService(itemRepo repo.IItemRepository) IItemService {
	return &itemService{itemRepo: itemRepo}
}

func (i itemService) GetUserItems(userID string) (*[]model.ItemDto, error) {
	panic("implement me")
}


