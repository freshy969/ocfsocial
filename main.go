package main

import (
	"log"
	"net/http"

	"github.com/cagox/ocfsocial/app/routes"
	"github.com/cagox/ocfsocial/app/util/config"
	"github.com/cagox/ocfsocial/app/util/database"
	"github.com/cagox/ocfsocial/app/util/logger"
)

func main() {
	logger.StartLogging()
	defer logger.StopLogging()

	database.DialMongoSession()

	routes.Routes()

	log.Fatal(http.ListenAndServe("localhost:8989", config.Config.Router))

}
