package db

import (
	"context"
	"log"

	// "gorm.io/driver/postgres"

	// "gorm.io/gorm"
	"github.com/jackc/pgx/v4"
)

type Handler struct {
	// DB *gorm.DB
	DB *pgx.Conn
}

func Init(url string) Handler {
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection Established!!!!")

	// db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// db.AutoMigrate(&models.User{})

	// defer conn.Close(context.Background())
	return Handler{conn}
}
