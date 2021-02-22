package main

import (
	"goMux-api/routers"
	"goMux-api/utils"
	"log"
	"net/http"
)

func main() {
	log.Println("Entered Main")
	utils.Init()
	r := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", r))
}
