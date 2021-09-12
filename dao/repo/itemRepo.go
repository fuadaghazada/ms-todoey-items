package repo

import (
	"errors"
	"github.com/fuadaghazada/ms-todoey-items/dao/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	log "github.com/sirupsen/logrus"
)

type IItemRepository interface {
	GetItemsByUserID(userID string) (*[]model.ItemEntity, error)
	GetItemByIDAndUserID(tx *pg.Tx, id int, userID string) (*model.ItemEntity, error)
	GetTransaction() (*pg.Tx, error)
	Commit(tx *pg.Tx)
	Rollback(tx *pg.Tx)
	CloseTransaction(tx *pg.Tx, err error)
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
	err := getItemsInfo(i.db.Model(&itemList), userID).Select()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Debug("ActionLog.GetItemsByUserID.end")

	return &itemList, nil
}

func (i itemRepository) GetItemByIDAndUserID(tx *pg.Tx, id int, userID string) (*model.ItemEntity, error) {
	log.Debug("ActionLog.GetItemByIDAndUserID.start")

	item := new(model.ItemEntity)
	err := tx.Model(item).
		Where("user_id = ?", userID).
		Where("id = ?", id).Select()

	if errors.Is(err, pg.ErrNoRows) {
		log.Debugf("ActionLog.GetItemsByIDAndUserID.info: No item row found with id: %v", id)
		return nil, nil
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Debug("ActionLog.GetItemByIDAndUserID.end")

	return item, nil
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

func (i itemRepository) CloseTransaction(tx *pg.Tx, err error) {
	if tx != nil {
		rec := recover()
		if err != nil || rec != nil {
			log.Error(err, " \n ", rec)
			i.Rollback(tx)
		} else {
			i.Commit(tx)
		}
	}
}

func getItemsInfo(q *orm.Query, userID string) *orm.Query {
	query := q.
		ColumnExpr("distinct items.id, items.title, items.description").
		ColumnExpr("items.user_id, items.created_at, items.updated_at")

	if userID != "" {
		return query.Where("items.user_id = ?", userID)
	}

	return query
}
