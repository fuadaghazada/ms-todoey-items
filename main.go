package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"ms-todoey-items/config"
	"net/http"
)

func main() {
	config.LoadConfig()

	port := config.Props.Port

	log.Info(fmt.Sprintf("Starting server at port: %v", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
