package datastore

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB Gloable database connection pool
var DB *sqlx.DB
