package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
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

func sendRequest(encoder *json.Encoder, decoder *json.Decoder, req Request) Response { ///////??????
	encoder.Encode(req)
	var resp Response
	decoder.Decode(&resp)
	return resp
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Сервер недоступен:", err)
		return
	}
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ваш логин:")
	scanner.Scan()
	myLogin := scanner.Text()

	fmt.Print("Ваш пароль:")
	scanner.Scan()
	myPassword := scanner.Text()

	resp := sendRequest(encoder, decoder, Request{
		Action:   "login",
		Login:    myLogin,
		Password: myPassword,
	})

	if resp.Status != "ok" {
		fmt.Println("Ошибка входа: ", resp.Error)
		return
	}
	fmt.Println("Вход выполнен")

	for {
		fmt.Println("\n=== GoChat ===")
		fmt.Println("1. Отправить сообщение")
		fmt.Println("2. Входящие")
		fmt.Println("3. История переписки")
		fmt.Println("0. Выход")
		fmt.Print("> ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Кому: ")
			scanner.Scan()
			to := scanner.Text()
			fmt.Print("Сообщение: ")
			scanner.Scan()
			text := scanner.Text()

			resp := sendRequest(encoder, decoder, Request{
				Action: "send",
				From:   myLogin,
				To:     to,
				Text:   text,
			})
			if resp.Status != "ok" {
				fmt.Println("Ошибка", resp.Error)
			} else {
				fmt.Println("Отправлено!")
			}

		case "2":
			resp := sendRequest(encoder, decoder, Request{
				Action: "inbox",
				From:   myLogin,
			})
			if len(resp.Messages) == 0 {
				fmt.Println("Сообщений нет")
			} else {
				for _, msg := range resp.Messages {
					fmt.Println(msg)
				}
			}

		case "3":
			fmt.Print("С кем история: ")
			scanner.Scan()
			other := scanner.Text()

			resp := sendRequest(encoder, decoder, Request{
				Action: "history",
				From:   myLogin,
				To:     other,
			})
			if len(resp.Messages) == 0 {
				fmt.Println("Нет переписки")
			} else {
				for _, msg := range resp.Messages {
					fmt.Println(msg)
				}
			}

		case "0":

			fmt.Println("Пока!")
			return
		}

	}

}
