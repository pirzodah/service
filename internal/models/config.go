package models

type Configuration struct {
	Version     string `json:"version"`
	Api         string `json:"api"`
	DatabaseUrl string `json:"DATABASE_URL"`
	MssqlDsn    string `json:"mssqlDsn"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}
