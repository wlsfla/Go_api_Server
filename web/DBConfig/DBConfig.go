package DBConfig

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
)

var (
	DBConn    *gorm.DB
	schemaStr string
)

func init() {
	schemaStr = "apiuser:1q2w3e4r!@tcp(host.docker.internal:3306)/GoAPIService"
	DBConn = getDBConn()
	connTest()
}

func connTest() {
	var version string
	DBConn.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

func getDBConn() *sql.DB {
	// if using mariaDB on docker-container, You have to change ip address "host.docker.internal"
	// issue https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host
	db, _ := sql.Open("mysql", schemaStr)

	return db
}

func Close() {
	DBConn.Close()
}

func Insert(query string) {
	DBConn.Exec(query)
}

func Update(query string) {
	DBConn.Exec(query)
}

func Select(query string) *sql.Rows {
	rows, _ := DBConn.Query(query)

	return rows
}
