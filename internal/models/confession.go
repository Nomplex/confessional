package models

import (
	"database/sql"
	"time"
)

type Confession struct {
	Id      int
	Content string
	Created time.Time
}

type ConfessionModel struct {
	DB *sql.DB
}
