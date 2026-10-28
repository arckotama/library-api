package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	mux := http.NewServeMux() // resepsionis atau router

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"message": "Library API v0.1"})
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		books := []Book{
			{ID: 1, Title: "スライム倒して300年、知らないうちにレベルMAXになってました", Author: "Kisetsu Morita", Year: 2017},
			{ID: 2, Title: "ヘルモード ～やり込み好きのゲーマーは廃設定の異世界で無双する～", Author: "Hamuo", Year: 2020},
		}
		writeJSON(w, http.StatusOK, books)
	})

	log.Println("Server jalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
