package handler

import (
	"encoding/json"
	"github.com/fuadaghazada/ms-todoey-items/exception"
	"github.com/fuadaghazada/ms-todoey-items/model"
	"github.com/fuadaghazada/ms-todoey-items/service"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type itemHandler struct {
	itemService service.IItemService
}

func NewItemHandler(router *chi.Mux, itemService service.IItemService) *chi.Mux {
	handler := &itemHandler{itemService: itemService}

	router.Get("/items", exception.ErrorHandler(handler.GetUserItems))

	return router
}

func (i *itemHandler) GetUserItems(w http.ResponseWriter, r *http.Request) error {
	userID := r.Header.Get(model.HeaderKeyUserID)

	if userID == "" {
		log.Error("ActionLog.GetUserItems.error: User ID is missing")
		return exception.NewBadRequestError("error.no-user-id", "No user ID found")
	}

	res, err := i.itemService.GetUserItems(userID)
	if err != nil {
		log.Errorf("ActionLog.GetUserItems.error: %v", err)
		return err
	}

	w.Header().Add(model.ContentType, model.ApplicationJSON)
	_ = json.NewEncoder(w).Encode(res)

	return nil
}