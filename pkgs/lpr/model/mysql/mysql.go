package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Lpr struct {
	InfoDate string
	OneYear  float32
	FiveYear float32
}

const (
	mysqlInster = iota
	mysqlQuery
)

var (
	errInvalidNoRowsAffected = errors.New("insert schedule:insert affected 0 rows")

	lprSQLString = []string{
		`INSERT INTO lpr(infodate, one_year, five_year) VALUES(?, ?, ?)`,
		`SELECT infodate, one_year, five_year FROM lpr`,
	}
)

func InsertLpr(db *sql.DB, InfoDate time.Time, OneYear, FiveYear float32) error {
	sql := fmt.Sprintf(lprSQLString[mysqlInster])
	result, err := db.Exec(sql, InfoDate, OneYear, FiveYear)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func QueryLpr(db *sql.DB) ([]*Lpr, error) {
	var (
		infoDate string
		oneYear  float32
		fiveYear float32

		lprs []*Lpr
	)

	sql := fmt.Sprintf(lprSQLString[mysqlQuery])
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&infoDate, &oneYear, &fiveYear); err != nil {
			return nil, err
		}

		lpr := &Lpr{
			InfoDate: infoDate,
			OneYear:  oneYear,
			FiveYear: fiveYear,
		}
		lprs = append(lprs, lpr)
	}

	return lprs, nil
}
