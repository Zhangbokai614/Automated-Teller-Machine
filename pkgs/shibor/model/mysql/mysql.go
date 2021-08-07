package mysql

import (
	"database/sql"
)

type ShiborTable struct {
	InfoDate   string `json:"date,omitempty"`
	OverNight  string `json:"O/N,omitempty"`
	OneWeek    string `json:"1W,omitempty"`
	TwoWeek    string `json:"2W,omitempty"`
	OneMonth   string `json:"1M,omitempty"`
	ThreeMonth string `json:"3M,omitempty"`
	SixMonth   string `json:"6M,omitempty"`
	NineMonth  string `json:"9M,omitempty"`
	OneYear    string `json:"1Y,omitempty"`
}

const (
	mysqlInster = iota
	mysqlQuery
)

var (
	execSql = []string{
		`INSERT INTO shibor(infodate, overnight, 1w, 2w, 1m, 3m, 6m, 9m, 1y) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		`SELECT infodate, overnight, 1w, 2w, 1m, 3m, 6m, 9m, 1y FROM shibor `,
	}
)

func Insert(db *sql.DB, d ShiborTable) error {
	_, err := db.Exec(execSql[mysqlInster], d.InfoDate, d.OverNight, d.OneWeek, d.TwoWeek, d.OneMonth, d.ThreeMonth, d.SixMonth, d.NineMonth, d.OneYear)
	if err != nil {
		return err
	}

	return nil
}
