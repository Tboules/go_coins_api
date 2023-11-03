package main

import (
	"fmt"
	"net/http"

	"github.com/Tboules/go_coins_api/internal/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go Api Services...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
