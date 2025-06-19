package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse-converter: ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)

	logger.Println("Сервер запущен на http://localhost:8080")
	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal("Server failed to start: ", err)
		return
	}
}
