package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"wsserver/configs"
)

var (
	dbConn *sql.DB
)

func InitMysql() {
	mysql_name := configs.GConf.DbUserName
	mysql_password := configs.GConf.DbPw
	mysql_host := configs.GConf.DbAddr
	mysql_port := configs.GConf.DbPort
	mysql_app_db := configs.GConf.DbDatabase

	//	dbConn, err = sql.Open("mysql", "root:lx123321@tcp(localhost:3306)/taptap?charset=UTF8")
	dbConn, err := sql.Open("mysql", mysql_name+":"+mysql_password+"@tcp("+mysql_host+":"+strconv.Itoa(mysql_port)+")/"+mysql_app_db+"?charset=UTF8")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("初始化mysql:", dbConn.Ping(), err)

}

func GetMysqlConn() (c *sql.DB) {
	return dbConn
}
