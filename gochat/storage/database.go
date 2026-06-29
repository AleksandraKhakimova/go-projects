package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type PostgresDB struct {
	conn *pgx.Conn ///Внутри — одно поле conn, это подключение к базе данных. Через conn мы будем отправлять SQL-запросы.
}

func NewPostgresDB(connString string) (*PostgresDB, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе: : %w", err)
	}
	return &PostgresDB{conn: conn}, nil ///& — берём адрес структуры (указатель)
}

func (db *PostgresDB) Close() error {
	return db.conn.Close(context.Background())
}

func (db *PostgresDB) AddUser(login, password string) error {
	_, err := db.conn.Exec(context.Background(),
		"INSERT INTO users (login, password) VALUES ($1, $2)",
		login, password,
	)

	if err != nil {
		return fmt.Errorf("пользователь уже существует: %w", err)
	}
	return nil
}

func (db *PostgresDB) SendMessage(from, to, text string) error {
	_, err := db.conn.Exec(context.Background(),
		"INSERT INTO messages (from_user, to_user, text) VALUES ($1, $2, $3)",
		from, to, text)

	return err
}

func (db *PostgresDB) GetInbox(login string) ([]string, error) {
	rows, err := db.conn.Query(context.Background(),
		"SELECT from_user, text, created_at FROM messages WHERE to_user= $1 ORDER BY created_at DESC",
		login,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var from, text string
		var createdAt time.Time
		rows.Scan(&from, &text, &createdAt)
		messages = append(messages, fmt.Sprintf("[%s] %s: %s", createdAt.Format("15:04"), from, text))
	}
	return messages, nil
}
func (db *PostgresDB) LoginUser(login, password string) error {
	var storedPassword string
	err := db.conn.QueryRow(context.Background(),
		"SELECT password FROM users WHERE login = $1", login,
	).Scan(&storedPassword)
	if err != nil {
		return fmt.Errorf("пользователь не найден")
	}
	if storedPassword != password {
		return fmt.Errorf("неверный пароль")
	}
	return nil
}

func (db *PostgresDB) GetHistory(user1, user2 string) ([]string, error) {
	rows, err := db.conn.Query(context.Background(),
		"SELECT from_user, text, created_at FROM messages WHERE (from_user=$1 AND to_user=$2) OR (from_user=$2 AND to_user=$1)ORDER BY created_at ASC",
		user1, user2,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var from, text string
		var createdAt time.Time
		rows.Scan(&from, &text, &createdAt)
		messages = append(messages,
			fmt.Sprintf("[%s] %s: %s", createdAt.Format("15:04"), from, text),
		)
	}
	return messages, nil
}
