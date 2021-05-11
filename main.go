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
	handler := routers.NewRouter()
	log.Fatal("server started on port : 5000", http.ListenAndServe(":5000", handler))
}
