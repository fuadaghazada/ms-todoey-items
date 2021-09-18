package handler

import (
	"encoding/json"
	"github.com/fuadaghazada/ms-todoey-items/exception"
	"github.com/fuadaghazada/ms-todoey-items/model"
	"github.com/fuadaghazada/ms-todoey-items/service"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

type itemHandler struct {
	itemService service.IItemService
}

func NewItemHandler(router *chi.Mux, itemService service.IItemService) *chi.Mux {
	handler := &itemHandler{itemService: itemService}

	router.Post("/items", exception.ErrorHandler(handler.CreateItem))
	router.Get("/items", exception.ErrorHandler(handler.GetUserItems))
	router.Get("/items/{id}", exception.ErrorHandler(handler.GetUserItem))
	router.Put("/items/{id}", exception.ErrorHandler(handler.UpdateItem))
	router.Delete("/items/{id}", exception.ErrorHandler(handler.DeleteItem))

	return router
}

func (i *itemHandler) CreateItem(w http.ResponseWriter, r *http.Request) error {
	userID, err := getUserID(r)
	if err != nil {
		log.Errorf("ActionLog.CreateItem.error: %v", err)
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	itemRequestDto := new(model.CreateUpdateItemDto)
	err = json.Unmarshal(body, itemRequestDto)
	if err != nil {
		log.Errorf("ActionLog.CreateItem.error: %v", err)
		return exception.NewBadRequestError("error.cannot-parse-item-data", "Cannot parse item data")
	}

	result, err := i.itemService.CreateItem(itemRequestDto, userID)
	if err != nil {
		log.Errorf("ActionLog.CreateItem.error: %v", err)
		return err
	}

	w.Header().Add(model.ContentType, model.ApplicationJSON)

	return json.NewEncoder(w).Encode(result)
}

func (i *itemHandler) GetUserItems(w http.ResponseWriter, r *http.Request) error {
	userID, err := getUserID(r)
	if err != nil {
		log.Errorf("ActionLog.GetUserItems.error: %v", err)
		return err
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

func (i *itemHandler) GetUserItem(w http.ResponseWriter, r *http.Request) error {
	userID, err := getUserID(r)
	if err != nil {
		log.Errorf("ActionLog.GetUserItem.error: %v", err)
		return err
	}
	itemIDStr := chi.URLParam(r, "id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		log.Errorf("ActionLog.GetUserItem.error: %v", err)
		return exception.NewBadRequestError("error.invalid-item-id", "Invalid item ID")
	}

	res, err := i.itemService.GetUserItem(itemID, userID)
	if err != nil {
		log.Errorf("ActionLog.GetUserItem.error: %v", err)
		return err
	}

	w.Header().Add(model.ContentType, model.ApplicationJSON)
	_ = json.NewEncoder(w).Encode(res)

	return nil
}

func (i *itemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) error {
	userID, err := getUserID(r)
	if err != nil {
		log.Errorf("ActionLog.UpdateItem.error: %v", err)
		return err
	}
	itemIDStr := chi.URLParam(r, "id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		log.Errorf("ActionLog.UpdateItem.error: %v", err)
		return exception.NewBadRequestError("error.invalid-item-id", "Invalid item ID")
	}

	body, err := ioutil.ReadAll(r.Body)
	itemRequestDto := new(model.CreateUpdateItemDto)
	err = json.Unmarshal(body, itemRequestDto)
	if err != nil {
		log.Errorf("ActionLog.UpdateItem.error: %v", err)
		return exception.NewBadRequestError("error.cannot-parse-item-data", "Cannot parse item data")
	}

	res, err := i.itemService.UpdateItem(itemRequestDto, itemID, userID)
	if err != nil {
		log.Errorf("ActionLog.UpdateItem.error: %v", err)
		return err
	}

	w.Header().Add(model.ContentType, model.ApplicationJSON)
	_ = json.NewEncoder(w).Encode(res)

	return nil
}

func (i *itemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) error {
	userID, err := getUserID(r)
	if err != nil {
		log.Errorf("ActionLog.UpdateItem.error: %v", err)
		return err
	}
	itemIDStr := chi.URLParam(r, "id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		log.Errorf("ActionLog.DeleteItem.error: %v", err)
		return exception.NewBadRequestError("error.invalid-item-id", "Invalid item ID")
	}

	res, err := i.itemService.DeleteItem(itemID, userID)
	if err != nil {
		log.Errorf("ActionLog.DeleteItem.error: %v", err)
		return err
	}

	w.Header().Add(model.ContentType, model.ApplicationJSON)
	_ = json.NewEncoder(w).Encode(res)

	return nil
}

func getUserID(r *http.Request) (string, error) {
	userID := r.Header.Get(model.HeaderKeyUserID)

	if userID == "" {
		log.Error("ActionLog.GetUserItems.error: User ID is missing")
		return userID, exception.NewBadRequestError("error.no-user-id", "No user ID found")
	}

	return userID, nil
}