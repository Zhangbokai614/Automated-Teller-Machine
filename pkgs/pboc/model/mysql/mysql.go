package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Pboc struct {
	InfoDate       string
	Period         int32
	Deal_amount    int32
	Rate           float32
	Trading_method string
}

const (
	mysqlInster = iota
	mysqlQuery
)

var (
	errInvalidNoRowsAffected = errors.New("insert schedule:insert affected 0 rows")

	pbocSQLString = []string{
		`INSERT INTO pboc(infodate, period, deal_amount, rate, trading_method) VALUES(?, ?, ?, ?, ?)`,
		`SELECT infodate, period, deal_amount, rate, trading_method FROM pboc`,
	}
)

func InsertPboc(db *sql.DB, InfoDate time.Time, Period, Deal_amount int32, Rate float32, Trading_method string) error {
	sql := fmt.Sprintf(pbocSQLString[mysqlInster])
	result, err := db.Exec(sql, InfoDate, Period, Deal_amount, Rate, Trading_method)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func QueryPboc(db *sql.DB) ([]*Pboc, error) {
	var (
		infoDate       string
		period         int32
		deal_amount    int32
		rate           float32
		trading_method string

		pbocs []*Pboc
	)

	sql := fmt.Sprintf(pbocSQLString[mysqlQuery])
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&infoDate, &period, &deal_amount, &rate, &trading_method); err != nil {
			return nil, err
		}

		pboc := &Pboc{
			InfoDate:       infoDate,
			Period:         period,
			Deal_amount:    deal_amount,
			Rate:           rate,
			Trading_method: trading_method,
		}
		pbocs = append(pbocs, pboc)
	}

	return pbocs, nil
}
