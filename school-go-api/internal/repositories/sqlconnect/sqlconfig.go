package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("Trying to connect to database")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")
	//connectionString := "root:pass@tcp(127.0.0.1:3306)/" + dbname // username:pass@tcp(ip:port)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		//panic(err)
		return nil, err
	}
	fmt.Println("Connected to maria db")
	return db, nil
}
