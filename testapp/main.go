package main // https://velog.io/@tae2089/Go%EC%97%90%EC%84%9C-Mysql-%EC%97%B0%EB%8F%99%ED%95%98%EA%B8%B0

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ConnPath struct {
	// db 계정이름
	username string
	// db 계정 패스워드
	password string
	// Socket 파일 위치
	socketPath string
	// 연결할 데이터베이스 이름
	database string
}

func GetSocketConn() string {
	connInfo := ConnPath{
		username:   "root",
		password:   "root",
		socketPath: "/var/run/mysqld/mysqld.sock",
		database:   "WinUpdate",
	}
	conn := connInfo.username + ":" + connInfo.password +
		"@unix(" + connInfo.socketPath + ")" + "/" +
		connInfo.database + "?charset=utf8"
	return conn
}

func main() {
	conn, err := sql.Open("mysql", GetSocketConn())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)
}
