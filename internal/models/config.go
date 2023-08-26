package models

type Configuration struct {
	Version     string `json:"version"`
	Api         string `json:"api"`
	PostgresDsn string `json:"postgresDsn"`
	MssqlDsn    string `json:"mssqlDsn"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}
