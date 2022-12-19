package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./log.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// db.Exec("CREATE TABLE `parent` (`id` int(11) NOT NULL, `name` varchar(100) NOT NULL, PRIMARY KEY (`id`))")
	db.Exec("CREATE TABLE 'update_log' ('host_ip' varchar(15) NOT NULL,'host_name' varchar(255) NOT NULL,'winver' char(5) NOT NULL,'result' int(5) NOT NULL,'created_time' datetime NOT NULL,'updated_time' datetime NOT NULL,PRIMARY KEY ('host_ip'))")
	db.Exec("INSERT INTO 'update_log' VALUES ('1.1.1.1', 'host_name', '1901', 0, datetime('now', 'localtime'))")

	rows, err := db.Query("SELECT * FROM update_log")
	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
