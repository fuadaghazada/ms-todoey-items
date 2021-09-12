package main

import (
	"fmt"
	"github.com/fuadaghazada/ms-todoey-items/config"
	"github.com/fuadaghazada/ms-todoey-items/db"
	log "github.com/sirupsen/logrus"

	"net/http"
)

func main() {
	config.LoadConfig()

	db.MigrateDb()
	dbCon := db.ConnectDb()
	defer dbCon.Close()

	port := config.Props.Port

	log.Info(fmt.Sprintf("Starting server at port: %v", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
