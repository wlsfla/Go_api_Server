package DBConfig

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
)

var (
	DBconn    *gorm.DB
	schemaStr string
)

func Init() {
	schemaStr = "apiuser:1q2w3e4r!@tcp(host.docker.internal:3306)/GoAPIService"
	DBconn = getDBConn()
}

func getDBConn() *gorm.DB {
	// if using mariaDB on docker-container, You have to change ip address "host.docker.internal"
	// issue https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host
	db, _ := gorm.Open(mysql.Open(schemaStr), &gorm.Config{})

	return db
}

func Insert(query string) {
	DBconn.Exec(query)
}

func Update(query string) {
	DBconn.Exec(query)
}
