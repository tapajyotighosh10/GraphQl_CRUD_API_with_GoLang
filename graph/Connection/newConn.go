package postgres

import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/lib/pq"
)

const (
	host     = "database-2.cwwzs37worn7.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "ghosh252"
	dbname   = "myDb"
	sslmode = "require"
)

func Conn() string {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s ", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("111successful conn")
	//sqlstatement:=`I`
	return psqlinfo
}
