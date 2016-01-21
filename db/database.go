package db

import (
	"database/sql"
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"log"
	"time"
)

type Today struct {
	Today dat.NullTime `db:"now"`
}

func Init(dsn string) *runner.DB {
	// create a normal database connection through database/sql
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// ensures the database can be pinged with an exponential backoff (15 min)
	runner.MustPing(db)

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// set to check things like sessions closing.
	// Should be disabled in production/release builds.
	dat.Strict = false

	// Log any query over 10ms as warnings. (optional)
	runner.LogQueriesThreshold = 10 * time.Millisecond

	DB := runner.NewDB(db, "postgres")

	// DoO a test run against the DB
	var _res Today
	dbErr := DB.SQL("select now()").QueryStruct(&_res)
	if dbErr != nil {
		log.Fatalln(dbErr.Error())
	}
	log.Println("... Test connection to DB, today =", _res.Today)

	return DB
}
