package DBConfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbobj *sql.DB
var schemaStr string

func init() {
	schemaStr = "apiuser:1q2w3e4r!@tcp(host.docker.internal:3306)/GoAPIService"
	dbobj = getDBConn()
	connTest()
}

func connTest() {
	var version string
	dbobj.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

func getDBConn() *sql.DB {
	// if using mariaDB on docker-container, You have to change ip address "host.docker.internal"
	// issue https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host
	db, _ := sql.Open("mysql", schemaStr)

	return db
}

func Close() {
	dbobj.Close()
}

func Insert(query string) {
	dbobj.Exec(query)
}

func Update(query string) {
	dbobj.Exec(query)
}

func Select(query string) *sql.Rows {
	rows, _ := dbobj.Query("select * from GoAPIService.target_winver")

	return rows
}
