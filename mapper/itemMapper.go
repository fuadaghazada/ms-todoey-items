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

func ToEntityCreate(itemDto *dto.CreateItemDto, userID string) *entity.ItemEntity {
	return &entity.ItemEntity{
		Title: itemDto.Title,
		Description: itemDto.Description,
		UserID: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
