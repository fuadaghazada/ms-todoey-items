package service

import (
	"fmt"
	"github.com/fuadaghazada/ms-todoey-items/dao/repo"
	"github.com/fuadaghazada/ms-todoey-items/exception"
	itemMapper "github.com/fuadaghazada/ms-todoey-items/mapper"
	"github.com/fuadaghazada/ms-todoey-items/model"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

type IItemService interface {
	GetUserItems(userID string) (*[]model.ItemDto, error)
	GetUserItem(id int, userID string) (*model.ItemDto, error)
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

func (i itemService) GetUserItem(id int, userID string) (*model.ItemDto, error) {
	log.Debugf("ActionLog.GetUserItem.start: Item#%v, User#%v", id, userID)

	tx, _ := i.itemRepo.GetTransaction()
	itemEntity, err := i.itemRepo.GetItemByIDAndUserID(tx, id, userID)
	if err != nil {
		log.Errorf("ActionLog.GetUserItem.error: %v", err)
		return nil, err
	}

	defer closeTransaction(tx, err, i.itemRepo.Rollback, i.itemRepo.Commit)

	if itemEntity == nil {
		log.Errorf("ActionLog.GetUserItem.error: Item not found %v", err)
		return nil, exception.NewItemNotFoundError("error.item-not-found", "Item not found")
	}

	itemDTO := itemMapper.ToDTO(itemEntity)

	log.Debugf("ActionLog.GetUserItem.end: Item#%v, User#%v", id, userID)

	return itemDTO, nil
}

func closeTransaction(tx *pg.Tx, err error, rollback func(*pg.Tx), commit func(*pg.Tx)) {
	if tx != nil {
		rec := recover()
		if err != nil || rec != nil {
			fmt.Println(err, " \n ", rec)
			rollback(tx)
		} else {
			commit(tx)
		}
	}
}
