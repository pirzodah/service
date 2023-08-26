package go_mssqldb

import (
	"context"
	"errors"
	_ "github.com/denisenkom/go-mssqldb"
)

func (t *MSSQLTransaction) Scan(dest ...interface{}) error {
	return errors.New("Scan is not supported in MSSQLTransaction")
}

func (t *MSSQLTransaction) Next() bool {
	return false
}

func (t *MSSQLTransaction) Close() error {
	return nil
}

func (t *MSSQLTransaction) Commit(ctx context.Context) error {
	return t.tx.Commit()
}

func (t *MSSQLTransaction) Rollback(ctx context.Context) error {
	return t.tx.Rollback()
}
