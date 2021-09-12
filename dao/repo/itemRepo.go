package repo

import (
	"github.com/fuadaghazada/ms-todoey-items/dao/model"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

type IItemRepository interface {
	GetItemsByUserID(userID string) (*[]model.ItemEntity, error)
	GetTransaction() (*pg.Tx, error)
	Commit(tx *pg.Tx)
	Rollback(tx *pg.Tx)
}

type itemRepository struct {
	db *pg.DB
}

func NewItemRepository(orm *pg.DB) IItemRepository {
	return &itemRepository{db: orm}
}

func (i itemRepository) GetItemsByUserID(userID string) (*[]model.ItemEntity, error) {
	log.Debug("ActionLog.GetItemsByUserID.start")

	itemList := make([]model.ItemEntity, 0)
	err := i.db.Model(&itemList).
		ColumnExpr("distinct items.id, items.title, items.description").
		ColumnExpr("items.user_id, items.created_at, items.updated_at").
		Where("items.user_id = ?", userID).
		Select()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Debug("ActionLog.GetItemsByUserID.end")

	return &itemList, nil
}

func (i *itemRepository) GetTransaction() (*pg.Tx, error) {
	t, err := i.db.Begin()
	return t, err
}

func (i itemRepository) Commit(tx *pg.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Error("Failed to commit current transaction ", err)
	}
}

func (i itemRepository) Rollback(tx *pg.Tx) {
	err := tx.Rollback()
	if err != nil {
		log.Error("Failed to rollback current transaction ", err)
	}
}
