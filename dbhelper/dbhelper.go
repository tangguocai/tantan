package dbhelper

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

const (
	host = "59.110.125.199:5432"
	//port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "tantan"
)

type PgHelper struct {
	Conn *pg.Conn
}

var Pg *PgHelper

func PgInit() {
	db := pg.Connect(&pg.Options{
		Addr:     host,
		User:     user,
		Password: password,
		Database: dbname,
	})
	//defer db.Close()

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1") //test connect
	if err != nil {
		panic(err)
	}

	Pg = &PgHelper{db.Conn()}
}

func (this *PgHelper) CreateTabel(models []interface{}) error {
	for _, model := range models {
		err := this.Conn.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
