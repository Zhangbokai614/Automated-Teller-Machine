package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Prr struct {
	InfoDate      string
	One_day       float32
	Seven_day     float32
	Fourteen_day  float32
	Twentyone_day float32
	One_month     float32
	Two_month     float32
}

const (
	mysqlInster = iota
	mysqlQuery
)

var (
	errInvalidNoRowsAffected = errors.New("insert schedule:insert affected 0 rows")

	prrSQLString = []string{
		`INSERT INTO prr(infodate, one_day, seven_day, fourteen_day, twentyone_day, one_month, two_month) VALUES(?, ?, ?, ?, ?, ?, ?)`,
		`SELECT infodate, one_day, seven_day, fourteen_day, twentyone_day, one_month, two_month FROM prr`,
	}
)

func InsertPrr(db *sql.DB, InfoDate time.Time, One_day, Seven_day, Fourteen_day, Twentyone_day, One_month, Two_month float32) error {
	sql := fmt.Sprintf(prrSQLString[mysqlInster])
	result, err := db.Exec(sql, InfoDate, One_day, Seven_day, Fourteen_day, Twentyone_day, One_month, Two_month)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func QueryPrr(db *sql.DB) ([]*Prr, error) {
	var (
		infoDate      string
		one_day       float32
		seven_day     float32
		fourteen_day  float32
		twentyone_day float32
		one_month     float32
		two_month     float32

		prrs []*Prr
	)

	sql := fmt.Sprintf(prrSQLString[mysqlQuery])
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&infoDate, &one_day, &seven_day, &fourteen_day, &twentyone_day, &one_month, &two_month); err != nil {
			return nil, err
		}

		prr := &Prr{
			InfoDate:      infoDate,
			One_day:       one_day,
			Seven_day:     seven_day,
			Fourteen_day:  fourteen_day,
			Twentyone_day: twentyone_day,
			One_month:     one_month,
			Two_month:     two_month,
		}
		prrs = append(prrs, prr)
	}

	return prrs, nil
}
