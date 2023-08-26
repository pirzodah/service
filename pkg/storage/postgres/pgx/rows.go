package pgx

import (
	"avesto-service/pkg/storage"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func (c *postgresDatabase) Close() {
	c.pool.Close()
}

func (c *postgresDatabase) Query(ctx context.Context, sql string, args ...interface{}) (storage.QueryResult, error) {
	rows, err := c.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &PostgresRows{rows: rows}, nil
}

func (c *postgresDatabase) QueryRow(ctx context.Context, sql string, args ...interface{}) storage.QueryResult {
	row := c.pool.QueryRow(ctx, sql, args...)
	return &PostgresRow{row: row}
}

func (c *postgresDatabase) Exec(ctx context.Context, sql string, arguments ...interface{}) (storage.CommandTag, error) {
	result, err := c.pool.Exec(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}
	return pgconn.CommandTag(result), nil
}

func (c *postgresDatabase) Begin(ctx context.Context) (storage.Transaction, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &PostgresTransaction{tx: tx}, nil
}

type PostgresRows struct {
	rows pgx.Rows
}

func (r *PostgresRows) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}

func (r *PostgresRows) Next() bool {
	return r.rows.Next()
}

func (r *PostgresRows) Close() error {
	r.rows.Close()
	return nil
}

type PostgresRow struct {
	row pgx.Row
}

func (r *PostgresRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}

func (r *PostgresRow) Next() bool {
	return true
}

func (r *PostgresRow) Close() error {
	return nil
}

type PostgresCommandTag struct {
	tag pgconn.CommandTag
}

func (ct *PostgresCommandTag) String() string {
	return ct.tag.String()
}

type PostgresTransaction struct {
	tx pgx.Tx
}
