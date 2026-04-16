package sqlconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(dbname string) (*sql.DB, error) {
	fmt.Println("Trying to connect to database")
	connectionString := "root:pass@tcp(127.0.0.1:3306)/" + dbname // username:pass@tcp(ip:port)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		//panic(err)
		return nil, err
	}
	fmt.Println("Connected to maria db")
	return db, nil
}
