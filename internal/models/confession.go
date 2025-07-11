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

func (c *ConfessionModel) Latest() ([]Confession, error) {
	stmt := `SELECT id, content, created FROM confessions 
	ORDER BY created LIMIT 3`

	rows, err := c.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var confessions []Confession
	for rows.Next() {
		var c Confession

		err = rows.Scan(&c.Id, &c.Content, &c.Created)
		if err != nil {
			return nil, err
		}

		confessions = append(confessions, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return confessions, nil
}
