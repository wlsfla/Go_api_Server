package ConnManager

import "fmt"

var (
	MaxConnCount  uint
	CurrConnCount uint
)

func init() {
	fmt.Println("\t[*] init ConnManager")

	MaxConnCount = 3
}

func SetMaxConnCount(value uint) {
	if value > 0 {
		MaxConnCount = value
	}
}

func Show() {

}
