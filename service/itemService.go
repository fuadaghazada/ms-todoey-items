package service

import (
	entity "github.com/fuadaghazada/ms-todoey-items/dao/model"
	"github.com/fuadaghazada/ms-todoey-items/dao/repo"
	"github.com/fuadaghazada/ms-todoey-items/exception"
	itemMapper "github.com/fuadaghazada/ms-todoey-items/mapper"
	"github.com/fuadaghazada/ms-todoey-items/model"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

type IItemService interface {
	CreateItem(itemDto *model.CreateUpdateItemDto, userID string) (*model.ItemDto, error)
	GetUserItems(userID string) (*[]model.ItemDto, error)
	GetUserItem(id int, userID string) (*model.ItemDto, error)
	UpdateItem(itemDto *model.CreateUpdateItemDto, itemID int, userID string) (*model.ItemDto, error)
}

type itemService struct {
	itemRepo repo.IItemRepository
}

func NewItemService(itemRepo repo.IItemRepository) IItemService {
	return &itemService{itemRepo: itemRepo}
}

func (i *itemService) CreateItem(itemDto *model.CreateUpdateItemDto, userID string) (*model.ItemDto, error) {
	log.Debugf("ActionLog.CreateItem.start: User#%v, Item#%v", userID, itemDto)

	itemEntity := itemMapper.ToEntityCreate(itemDto, userID)

	tx, _ := i.itemRepo.GetTransaction()
	createdItem, err := i.itemRepo.SaveItem(tx, itemEntity)
	if err != nil {
		log.Error("ActionLog.CreateItem.error: Database error")
		return nil, exception.NewDatabaseError()
	}

	defer i.itemRepo.CloseTransaction(tx, err)

	resultItemDto := itemMapper.ToDTO(createdItem)

	log.Debugf("ActionLog.CreateItem.end")

	return resultItemDto, nil
}

func (i *itemService) GetUserItems(userID string) (*[]model.ItemDto, error) {
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

func (i *itemService) GetUserItem(id int, userID string) (*model.ItemDto, error) {
	log.Debugf("ActionLog.GetUserItem.start: Item#%v, User#%v", id, userID)

	itemEntity, tx, err := i.getItem(id, userID)
	if err != nil {
		log.Errorf("ActionLog.GetUserItem.error: %v", err)
		return nil, err
	}

	defer i.itemRepo.CloseTransaction(tx, err)

	itemDTO := itemMapper.ToDTO(itemEntity)

	log.Debugf("ActionLog.GetUserItem.end: Item#%v, User#%v", id, userID)

	return itemDTO, nil
}

func (i *itemService) UpdateItem(itemDto *model.CreateUpdateItemDto, itemID int, userID string) (*model.ItemDto, error) {
	log.Debugf("ActionLog.UpdateItem.start: User#%v, Item#%v", userID, itemDto)

	itemEntity, tx, err := i.getItem(itemID, userID)
	if err != nil {
		log.Errorf("ActionLog.UpdateItem.error: %v", err)
		return nil, err
	}

	itemEntity = itemMapper.ToEntityUpdate(itemDto, itemEntity)

	updatedItem, err := i.itemRepo.SaveItem(tx, itemEntity)
	if err != nil {
		log.Errorf("ActionLog.UpdateItem.error: Database error %v", err)
		return nil, exception.NewDatabaseError()
	}

	defer i.itemRepo.CloseTransaction(tx, err)

	resultItemDto := itemMapper.ToDTO(updatedItem)

	log.Debugf("ActionLog.CreateItem.end")

	return resultItemDto, nil
}

func (i *itemService) getItem(id int, userID string) (*entity.ItemEntity, *pg.Tx, error) {
	log.Debug("ActionLog.getItem.start")

	tx, _ := i.itemRepo.GetTransaction()

	itemEntity, err := i.itemRepo.GetItemByIDAndUserID(tx, id, userID)
	if err != nil {
		log.Errorf("ActionLog.getItem.error: %v", err)
		return nil, tx, err
	}

	if itemEntity == nil {
		log.Errorf("ActionLog.getItem.error: Item not found %v", err)
		return nil, tx, exception.NewItemNotFoundError("error.item-not-found", "Item not found")
	}

	log.Debug("ActionLog.getItem.end")

	return itemEntity, tx, nil
}
