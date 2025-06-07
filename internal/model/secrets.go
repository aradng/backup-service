package model

type DatabaseType int

type SecretSchemaType struct {
	Name            string
	Username        string
	Password        string
	DefaultPassword string
}

const (
	MongoDB DatabaseType = iota
	MySQLDB
)

var SecretSchemas = map[DatabaseType]SecretSchemaType{
	MongoDB: {Name: "mongo", Username: "MONGO_INITDB_ROOT_USERNAME", Password: "MONGO_INITDB_ROOT_PASSWORD"},
	MySQLDB: {Name: "mysql", Username: "MYSQL_ROOT_USER", Password: "MYSQL_ROOT_PASSWORD", DefaultPassword: "root"},
}
