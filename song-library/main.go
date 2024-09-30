package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {
    
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Ошибка загрузки .env файла")
    }

    r := mux.NewRouter()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    log.Printf("Сервер запущен на порту %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
