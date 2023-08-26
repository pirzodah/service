package storage

import (
	"context"
)

//type Database interface {
//	Close()
//	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
//	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
//	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
//	Begin(ctx context.Context) (Transaction, error)
//}
//
//type Transaction interface {
//	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
//	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
//	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
//	Commit(ctx context.Context) error
//	Rollback(ctx context.Context) error
//}

type QueryResult interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type CommandTag interface {
	String() string
}

type Transaction interface {
	QueryResult
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	QueryRow(ctx context.Context, sql string, args ...interface{}) QueryResult
	Query(ctx context.Context, sql string, args ...interface{}) (QueryResult, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (CommandTag, error)
}

type Database interface {
	Query(ctx context.Context, sql string, args ...interface{}) (QueryResult, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) QueryResult
	Exec(ctx context.Context, sql string, arguments ...interface{}) (CommandTag, error)
	Begin(ctx context.Context) (Transaction, error)
	Close()
}
