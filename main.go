package main

import (
	"fmt"
	"heartvoice/src/config"
	"heartvoice/src/router"
	"log"
	"net/http"
)

func main() {

	config.LoadEnvs()

	fmt.Printf("Listening at port %d!\n", config.Port)
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
