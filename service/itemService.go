package service

import (
	"github.com/fuadaghazada/ms-todoey-items/dao/repo"
	itemMapper "github.com/fuadaghazada/ms-todoey-items/mapper"
	"github.com/fuadaghazada/ms-todoey-items/model"
	log "github.com/sirupsen/logrus"
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
	log.Debugf("ActionLog.GetUserItems.start: User#%v", userID)

	itemEntities, err := i.itemRepo.GetItemsByUserID(userID)
	if err != nil {
		log.Errorf("ActionLog.GetUserItems.error: %v", err)
		return nil, err
	}

	itemDTOs := itemMapper.ToDTOs(*itemEntities)

	log.Debugf("ActionLog.GetUserItems.end: User#%v", userID)

	return &itemDTOs, nil
}


