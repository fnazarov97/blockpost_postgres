package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

const auth = "user=farhod password=f@rhod666997 dbname=blockpost sslmode=disable"

func InitDB() (p *Postgres, err error) {
	db, err := sqlx.Connect("postgres", auth)
	if err != nil {
		return p, err
	}
	fmt.Println("ok-------------------")
	return &Postgres{
		DB: db,
	}, nil
}
