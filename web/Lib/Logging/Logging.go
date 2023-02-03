package Logging

import (
	"fmt"
	"os"
)

var (
	Fileptr *os.File
)

func init() {
	fmt.Println("\t[*] init Logging")

	// accesslog write
	Fileptr, _ = os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func Close() {
	Fileptr.Close()
}
