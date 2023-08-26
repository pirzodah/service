package clients

import (
	"avesto-service/internal/database"
	pkgpostgres "avesto-service/pkg/storage"
)

type db struct {
	postgres pkgpostgres.Database
}

func New(postgres pkgpostgres.Database) database.Database {
	return &db{postgres: postgres}
}
