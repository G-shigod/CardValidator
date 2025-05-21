package main

import (
	"card-validator/internal/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Регистрируем API-обработчик
	mux.HandleFunc("/validate", handler.ValidateCard)

	// Регистрируем отдачу статических файлов из папки static по корню "/"
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/", fileServer)

	addr := ":8080"
	log.Printf("Сервер запущен на %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
