package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

const (
	mysqlInster = iota
	mysqlQuery
)

var (
	errInvalidNoRowsAffected = errors.New("insert schedule:insert affected 0 rows")

	shiborSQLString = []string{
		`INSERT INTO shibor(infodate, overnight, 1w, 2w, 1m, 3m, 6m, 9m, 1y) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		`SELECT infodate, overnight, 1w, 2w, 1m, 3m, 6m, 9m, 1y FROM shibor `,
	}
)

func InsertShibor(
	db *sql.DB, InfoDate string, OverNight, OneWeek, TwoWeek,
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
