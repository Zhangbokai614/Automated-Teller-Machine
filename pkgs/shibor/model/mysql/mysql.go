package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Shibor struct {
	InfoDate   string
	OverNight  float32
	OneWeek    float32
	TwoWeek    float32
	OneMonth   float32
	ThreeMonth float32
	SixMonth   float32
	NineMonth  float32
	OneYear    float32
}

const (
	mysqlInster = iota
	mysqlQuery
)

var (
	errInvalidNoRowsAffected = errors.New("insert schedule:insert affected 0 rows")

	shiborSQLString = []string{
		`INSERT INTO shibor(infodate, overnight, one_week, two_week, one_month, three_month, six_month, nine_month, one_year) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		`SELECT infodate, overnight, one_week, two_week, one_month, three_month, six_month, nine_month, one_year FROM shibor`,
	}
)

func InsertShibor(
	db *sql.DB, InfoDate time.Time, OverNight, OneWeek, TwoWeek,
	OneMonth, ThreeMonth, SixMonth, NineMonth, OneYear float32,
) error {
	sql := fmt.Sprintf(shiborSQLString[mysqlInster])
	result, err := db.Exec(sql, InfoDate, OverNight, OneWeek, TwoWeek, OneMonth, ThreeMonth, SixMonth, NineMonth, OneYear)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func QueryShibor(db *sql.DB) ([]*Shibor, error) {
	var (
		infoDate   string
		overNight  float32
		oneWeek    float32
		twoWeek    float32
		oneMonth   float32
		threeMonth float32
		sixMonth   float32
		nineMonth  float32
		oneYear    float32

		shibors []*Shibor
	)

	sql := fmt.Sprintf(shiborSQLString[mysqlQuery])
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&infoDate, &overNight, &oneWeek, &twoWeek, &oneMonth,
			&threeMonth, &sixMonth, &nineMonth, &oneYear); err != nil {
			return nil, err
		}

		shibor := &Shibor{
			InfoDate:   infoDate,
			OverNight:  overNight,
			OneWeek:    oneWeek,
			TwoWeek:    twoWeek,
			OneMonth:   oneMonth,
			ThreeMonth: threeMonth,
			SixMonth:   sixMonth,
			NineMonth:  nineMonth,
			OneYear:    oneYear,
		}
		shibors = append(shibors, shibor)
	}

	return shibors, nil
}
