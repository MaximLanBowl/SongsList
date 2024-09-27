package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfigToConnect struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLmode  string
}

var DB *sqlx.DB

func NewPostgres(config ConfigToConnect) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBname, config.SSLmode))
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удается подключиться к базе данных: %v", err)
	}
	log.Println("Connected to the database")
	return db, nil
}
