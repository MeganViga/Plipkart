package db

import "database/sql"

//to generate MOCkDB
type Store interface{
	Querier
}
type SQLStore struct{
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB)Store{
	return SQLStore{
		Queries: New(db),
		db:db,
	}
}