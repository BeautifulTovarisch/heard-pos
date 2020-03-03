// Package provides shared utilities for performing database operations.
// Intended to abstract away lower-level database interaction from 'CRUD' modules.
//
// Contains opt-in database schema creation.
//
// TODO:
//   - Fetch database url, password from kv-store (consul, etcd, riak etc...)
//   - Decide on recovery strategy
//   - Incorporate proper logging
//
package database

import (
	"fmt"
	"os"

	"heard/product"
	"heard/ticket"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

var create_user = fmt.Sprintf(`
CREATE USER pos PASSWORD '%s';
`, os.Getenv("USER_PASSWORD"))

const create_schema = `
CREATE SCHEMA IF NOT EXISTS AUTHORIZATION pos;
`

const create_database = `
CREATE DATABASE pos OWNER pos;
`

// Centralize database connection for convenience
func Connect(user, database, password string) *sqlx.DB {
	conn_string := fmt.Sprintf("postgres://%s:%s@database-service/%s", user, password, database)
	conn, err := sqlx.Connect("pgx", conn_string)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func setup_database(conn *sqlx.DB) {
	defer conn.Close()

	_, err := conn.Exec(create_user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating 'pos' user: %v\n", err)
	}

	_, err = conn.Exec(create_schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating 'pos' schema: %v\n", err)
	}

	_, err = conn.Exec(create_database)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating 'pos' database: %v\n", err)
	}
}

func SetupSchema() {
	conn := Connect("postgres", "postgres", os.Getenv("POSTGRES_PASSWORD"))
	setup_database(conn)

	conn = Connect("pos", "pos", os.Getenv("USER_PASSWORD"))
	defer conn.Close()

	_, err := conn.Exec(ticket.Schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating 'ticket' schema: %v\n", err)
	}

	_, err = conn.Exec(product.Schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating 'product' schema: %v\n", err)
	}
}
