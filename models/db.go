package models

import (
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	app "github.com/saidamir98/go-boilerplate/app"
)

type BaseModel struct {
	Id        int        `json:"id" db:"id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

func InitDB() {
	var (
		err error
	)

	dbUri := app.Conf["DATABASE_URL"]

	app.DB, err = sqlx.Connect("pgx", dbUri)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err = app.DB.Ping(); err != nil {
		log.Fatalf("%+v", err)
	}

	log.Println("connected db...")

	_, err = app.DB.Exec(Schemas)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Println(Schemas)
}

var Schemas = `
	DROP TABLE IF EXISTS users;
	CREATE TABLE users(
		id serial PRIMARY KEY,
		username VARCHAR (255) UNIQUE NOT NULL,
		email VARCHAR (255) UNIQUE NOT NULL,
		password VARCHAR (255) NOT NULL,
		role_id INTEGER,
		active BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);
	`
