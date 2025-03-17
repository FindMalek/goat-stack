package db

import (

	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
	 "github.com/jmoiron/sqlx"
)


var schema = `
create table if not exists vote (
    id integer primary key,
    timestamp integer,
    vote text not null
);
`

func New() *sqlx.DB {

	db, err := sqlx.Open("sqlite", "file:db.sqlite")

	if err != nil {
		log.Fatal().Err(err).Msg("failed to open database")
		return nil
	}

	if _, err = db.Exec(schema); err != nil {
		return nil
	}
	return db

}
