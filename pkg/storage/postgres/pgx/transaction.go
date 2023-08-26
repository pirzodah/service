package pgx

import (
	"avesto-service/pkg/storage"
	"context"
)

func (t *PostgresTransaction) Query(ctx context.Context, sql string, args ...interface{}) (storage.QueryResult, error) {
	rows, err := t.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &PostgresRows{rows: rows}, nil
}

func (t *PostgresTransaction) QueryRow(ctx context.Context, sql string, args ...interface{}) storage.QueryResult {
	row := t.tx.QueryRow(ctx, sql, args...)
	return &PostgresRow{row: row}
}

func (t *PostgresTransaction) Exec(ctx context.Context, sql string, arguments ...interface{}) (storage.CommandTag, error) {
	result, err := t.tx.Exec(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}
	return &PostgresCommandTag{tag: result}, nil
}

func (t *PostgresTransaction) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *PostgresTransaction) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

func (t *PostgresTransaction) Scan(dest ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t *PostgresTransaction) Next() bool {
	//TODO implement me
	panic("implement me")
}

func (t *PostgresTransaction) Close() error {
	//TODO implement me
	panic("implement me")
}
