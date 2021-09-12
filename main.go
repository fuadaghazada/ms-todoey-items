package main

import (
	"fmt"
	"github.com/fuadaghazada/ms-todoey-items/config"
	"github.com/fuadaghazada/ms-todoey-items/db"
	"github.com/fuadaghazada/ms-todoey-items/handler"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"net/http"
)

func main() {
	config.LoadConfig()

	db.MigrateDb()
	dbCon := db.ConnectDb()
	defer dbCon.Close()

	router := chi.NewRouter()
	handler.NewHealthHandler(router)

	port := config.Props.Port
	log.Info(fmt.Sprintf("Starting server at port: %v", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
