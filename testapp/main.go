package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbconnTest()
}

func dbconnTest() {
	// if using mariaDB on docker-container, You have to change ip address "host.docker.internal"
	// issue https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host
	db, err := sql.Open("mysql", "root:1q2w3e4r!@tcp(host.docker.internal:3306)/GoAPIService")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
