package main

import (
	"fmt"
	"net/http"

	"github.com/TheAmgadX/APIs-With-GoLang/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true) // Enables file & line number tracking
	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting Go API service...")
	fmt.Println(`  
   ██████╗  ██████╗     █████╗ ██████╗ ██╗
  ██╔════╝ ██╔═══██╗   ██╔══██╗██╔══██╗██║
  ██║  ███╗██║   ██║   ███████║██████╔╝██║
  ██║   ██║██║   ██║   ██╔══██║██╔═══╝ ██║
  ╚██████╔╝╚██████╔╝██╗██║  ██║██║     ██║
   ╚═════╝  ╚═════╝ ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝
    `)

	err := http.ListenAndServe("127.0.0.1:8080", r) // running on localhost
	if err != nil {
		log.Error(err)
	}
}
