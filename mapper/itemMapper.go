package mapper

import (
	entity "github.com/fuadaghazada/ms-todoey-items/dao/model"
	dto "github.com/fuadaghazada/ms-todoey-items/model"
	"time"
)

func ToDTO(itemEntity *entity.ItemEntity) *dto.ItemDto {
	return &dto.ItemDto{
		ID:          itemEntity.ID,
		Title:       itemEntity.Title,
		Description: itemEntity.Description,
		CreatedAt:   itemEntity.CreatedAt,
		UpdatedAt:   itemEntity.UpdatedAt,
	}
}

func ToDTOs(itemEntities []entity.ItemEntity) []dto.ItemDto {
	itemDTOs := make([]dto.ItemDto, len(itemEntities))

	for i, itemEntity := range itemEntities {
		itemDTOs[i] = *ToDTO(&itemEntity)
	}

	return itemDTOs
}

func ToEntityCreate(itemDto *dto.CreateUpdateItemDto, userID string) *entity.ItemEntity {
	return &entity.ItemEntity{
		Title: itemDto.Title,
		Description: itemDto.Description,
		UserID: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToEntityUpdate(itemDto *dto.CreateUpdateItemDto, itemEntity *entity.ItemEntity) *entity.ItemEntity {
	updated := false
	title := itemEntity.Title
	description := itemEntity.Description
	updatedAt := itemEntity.UpdatedAt

	if itemDto.Title != "" {
		title = itemDto.Title
		updated = true
	}
	if itemDto.Description != "" {
		description = itemDto.Description
		updated = true
	}
	if updated {
		updatedAt = time.Now()
	}

	return &entity.ItemEntity{
		ID: itemEntity.ID,
		Title: title,
		Description: description,
		UserID: itemEntity.UserID,
		CreatedAt: itemEntity.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
