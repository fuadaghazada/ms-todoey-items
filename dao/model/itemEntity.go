package model

import "time"

type ItemEntity struct {
	tableName struct{} `sql:"items, alias:items"`

	ID          int       `sql:",pk" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
