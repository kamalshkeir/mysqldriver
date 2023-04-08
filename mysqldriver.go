package mysqldriver

import (
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

var Used = false
var cd = mysql.MySQLDriver{}

func Use() driver.Driver {
	Used = true
	return &cd
}

func GetDriver() driver.Driver {
	return &cd
}
