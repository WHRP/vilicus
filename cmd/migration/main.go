package main

import (
	"log"
	"os"

	"github.com/edersonbrilhante/ccvs"
	"github.com/edersonbrilhante/ccvs/pkg/util/postgres"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func main() {
	db, err := postgres.New(os.Getenv("DATABASE_URL"), 0, false)
	checkErr(err)
	err = createSchema(db)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createSchema(db *pg.DB) error {
    models := []interface{}{
        (*ccvs.Analysis)(nil),
    }
	
	_, err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public")
	if err != nil {
		return err
	}

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		return err
    }
    return nil
}
