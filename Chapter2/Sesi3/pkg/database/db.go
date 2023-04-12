package database

import (
	cm "SWSGolang/pkg/common"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


func Connection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cm.Host, cm.Port,cm.User, cm.Password, cm.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}