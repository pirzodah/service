package pgx

import (
	"avesto-service/internal/models"
	"avesto-service/pkg/storage"
	"avesto-service/pkg/storage/postgres/pgx/afterconnect"
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresDatabase struct {
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context, databaseDsn models.Configuration) (storage.Database, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	poolConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Println(err)
		return nil, errors.New("pgxpool.ParseConfig ERROR: " + err.Error())
	}
	poolConfig.MaxConns = 10
	poolConfig.AfterConnect = afterConnect

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Println(err)
		return nil, errors.New("pgxpool.ConnectConfig ERROR: " + err.Error())
	}
	return &postgresDatabase{pool: pool}, nil
}

func afterConnect(ctx context.Context, conn *pgx.Conn) (err error) {

	err = afterconnect.AfterConnectService(ctx, conn)
	if err != nil {
		err = errors.New("AfterConnectService error: " + err.Error())
		return
	}

	err = afterconnect.AfterConnectInputs(ctx, conn)
	if err != nil {
		err = errors.New("AfterConnectInputs error: " + err.Error())
		return
	}
	err = afterconnect.AfterConnectServiceInputs(ctx, conn)
	if err != nil {
		err = errors.New("AfterConnectServiceInputs error: " + err.Error())
		return
	}
	return
}
