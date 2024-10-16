package main

import (
	"fmt"
	"net/http"
	handlers "second-task/task2/internal/server"
	"time"
)

func main() {
	PORT := ":8001"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  0 * time.Second,
		WriteTimeout: 0 * time.Second,
	}
	mux.Handle("/version", http.HandlerFunc(handlers.GetVersion))
	mux.Handle("/decode", http.HandlerFunc(handlers.Decode))
	mux.Handle("/hard-op", http.HandlerFunc(handlers.HardOp))
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Ошибка при запуске сервера")
		return
	}
}
