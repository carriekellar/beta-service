package main

import (
	"beta_service/db"
	"beta_service/handlers"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Open database connection
	database, err := db.Open()
	log.Print("Database Opened")
	if err != nil {
		print(err)
	}

	defer database.Close()

	h := handlers.NewHandler(database)

	r := mux.NewRouter()
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	r.HandleFunc("/assets", h.HandleGetAssets).Methods("GET")
	// r.HandleFunc("/assets", h.HandleGetAssets).Methods("GET")
	// r.HandleFunc("/kycform", handlers.h.handleUserForm).Methods("POST")

	s := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Server Open")

	// Listen for interrupt or terminal signal from the OS (e.g. Command+C)
	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-shutdown
		switch sig {
		case os.Interrupt:
			os.Exit(0)
		case syscall.SIGINT:
			os.Exit(0)
		}
	}()

	log.Print("startup", "status", "api router started", "host", &s.Addr)
	log.Fatal(s.ListenAndServe())

}
