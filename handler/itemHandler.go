package handler

import (
	"encoding/json"
	"github.com/fuadaghazada/ms-todoey-items/service"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const contentTypeString = "Content-Type"
const applicationJSONString = "application/json"

type itemHandler struct {
	itemService service.IItemService
}

func NewItemHandler(router *chi.Mux, itemService service.IItemService) *chi.Mux {
	handler := &itemHandler{itemService: itemService}

	router.Get("/items/{userID}", ErrorHandler(handler.GetUserItems))

	return router
}

func (i *itemHandler) GetUserItems(w http.ResponseWriter, r *http.Request) error {
	userID := chi.URLParam(r, "userID")

	res, err := i.itemService.GetUserItems(userID)
	if err != nil {
		log.Errorf("ActionLog.GetUserItems.error: %v", err)
		return err
	}

	w.Header().Add(contentTypeString, applicationJSONString)
	_ = json.NewEncoder(w).Encode(res)

	return nil
}