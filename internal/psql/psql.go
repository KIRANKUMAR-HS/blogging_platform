package db

import (
	"database/sql"

	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

type PsqlClient struct {
	db *sql.DB
}

func NewPsqlClint(connStr string) (*PsqlClient, error) {

	// open database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Err(err).Msg("failed to open psql connection")
		return nil, err
	}

	// close database
	//defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		log.Err(err).Msg("failed to ping db")
		return nil, err
	}

	return &PsqlClient{
		db: db,
	}, err
}



func (p *PsqlClient) Close() error {
	return p.db.Close()
}


