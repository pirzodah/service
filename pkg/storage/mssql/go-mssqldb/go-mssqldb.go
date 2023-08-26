package go_mssqldb

import (
	"avesto-service/internal/models"
	"avesto-service/pkg/storage"
	"avesto-service/pkg/storage/mssql/go-mssqldb/mssqlafterconnect"
	"context"
	"database/sql"
	"errors"
	_ "github.com/denisenkom/go-mssqldb"
)

func NewMSSQLClient(ctx context.Context, databaseDsn models.Configuration) (storage.Database, error) {
	db, err := sql.Open("sqlserver", databaseDsn.MssqlDsn)
	if err != nil {
		return nil, err
	}

	if err := afterConnectMSSQL(ctx, db); err != nil {
		return nil, err
	}

	return &mssqlDatabase{db: db}, nil
}

func afterConnectMSSQL(ctx context.Context, conn *sql.DB) (err error) {

	err = mssqlafterconnect.AfterConnectService(ctx, conn)
	if err != nil {
		err = errors.New("AfterConnectService error: " + err.Error())
		return
	}

	err = mssqlafterconnect.AfterConnectInputs(ctx, conn)
	if err != nil {
		err = errors.New("AfterConnectInputs error: " + err.Error())
		return
	}
	err = mssqlafterconnect.AfterConnectServiceInputs(ctx, conn)
	if err != nil {
		err = errors.New("AfterConnectServiceInputs error: " + err.Error())
		return
	}
	return
}

type mssqlDatabase struct {
	db *sql.DB
}

type MSSQLTransaction struct {
	tx *sql.Tx
}

type MSSQLRows struct {
	rows *sql.Rows
}

type MSSQLRow struct {
	row *sql.Row
}

func (r *MSSQLRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}

func (r *MSSQLRow) Next() bool {
	return true
}

func (r *MSSQLRow) Close() error {
	return nil
}
func (c *mssqlDatabase) Close() {
	c.db.Close()
}

func (c *mssqlDatabase) Query(ctx context.Context, sql string, args ...interface{}) (storage.QueryResult, error) {
	rows, err := c.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &MSSQLRows{rows: rows}, nil
}

func (c *mssqlDatabase) QueryRow(ctx context.Context, sql string, args ...interface{}) storage.QueryResult {
	row := c.db.QueryRowContext(ctx, sql, args...)
	return &MSSQLRow{row: row}
}

func (c *mssqlDatabase) Exec(ctx context.Context, sql string, arguments ...interface{}) (storage.CommandTag, error) {
	result, err := c.db.ExecContext(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}
	return &MSSQLCommandTag{result: result}, nil
}

func (c *mssqlDatabase) Begin(ctx context.Context) (storage.Transaction, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &MSSQLTransaction{tx: tx}, nil
}

type MSSQLCommandTag struct {
	result sql.Result
}

func (t *MSSQLCommandTag) String() string {
	return ""
}

func (r *MSSQLRows) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}

func (r *MSSQLRows) Next() bool {
	return r.rows.Next()
}

func (r *MSSQLRows) Close() error {
	r.rows.Close()
	return nil
}
