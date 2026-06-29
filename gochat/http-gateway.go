package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type Request struct {
	Action   string `json:"action"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Text     string `json:"text,omitempty"`
}

type Response struct {
	Status   string   `json:"status"`
	Error    string   `json:"error,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

func main() {
	fmt.Println("🌌 GoChat HTTP Gateway")
	fmt.Println("🚀 Запущен на порту 8082")
	fmt.Println("📋 Откройте index.html в браузере")

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Подключаемся к TCP серверу
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			http.Error(w, "TCP server not available: "+err.Error(), http.StatusServiceUnavailable)
			return
		}
		defer conn.Close()

		encoder := json.NewEncoder(conn)
		if err := encoder.Encode(req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var resp Response
		decoder := json.NewDecoder(conn)
		if err := decoder.Decode(&resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Слушаем на localhost
	http.ListenAndServe(":8082", nil)
}
