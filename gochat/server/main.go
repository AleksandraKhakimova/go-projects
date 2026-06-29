package main

import (
	"encoding/json"
	"fmt"
	"gochat/storage"
	"net"
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
	// Подключаемся к базе
	db, err := storage.NewPostgresDB("postgres://localhost:5432/gochat?sslmode=disable")
	if err != nil {
		fmt.Println("Ошибка базы:", err)
		return
	}
	defer db.Close()

	// Слушаем на localhost
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка запуска:", err)
		return
	}
	defer listener.Close()
	fmt.Println("✅ Сервер запущен на порту 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn, db)
	}
}

func handleClient(conn net.Conn, db *storage.PostgresDB) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	for {
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			break
		}

		var resp Response

		switch req.Action {
		case "login":
			err := db.AddUser(req.Login, req.Password)
			if err != nil {
				err = db.LoginUser(req.Login, req.Password)
				if err != nil {
					resp = Response{Status: "error", Error: err.Error()}
				} else {
					resp = Response{Status: "ok"}
				}
			} else {
				resp = Response{Status: "ok"}
			}
		case "send":
			err := db.SendMessage(req.From, req.To, req.Text)
			if err != nil {
				resp = Response{Status: "Error", Error: err.Error()}
			} else {
				resp = Response{Status: "ok"}
			}
		case "inbox":
			msgs, err := db.GetInbox(req.Login)
			if err != nil {
				resp = Response{Status: "Error", Error: err.Error()}
			} else {
				resp = Response{Status: "ok", Messages: msgs}
			}
		case "history":
			msgs, err := db.GetHistory(req.From, req.To)
			if err != nil {
				resp = Response{Status: "error", Error: err.Error()}
			} else {
				resp = Response{Status: "ok", Messages: msgs}
			}
		default:
			resp = Response{Status: "error", Error: "неизвестное действие"}
		}
		encoder.Encode(resp)
	}
}
