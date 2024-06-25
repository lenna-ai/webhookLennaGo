package simples

type DatabasePostgresDsn string
type DatabaseMongodbDsn string

type Database struct {
	DatabasePostgres string
	DatabaseMongoDB  string
}

type FullName struct {
	Fullname string
}

func NewDatabasePostgresDsn() *DatabasePostgresDsn {
	var dataPgsql = "pgsql"
	return (*DatabasePostgresDsn)(&dataPgsql)
}

func NewDatabaseMongodbDsn() *DatabaseMongodbDsn {
	var dataMongoDb = "mongoDb"
	return (*DatabaseMongodbDsn)(&dataMongoDb)
}

func NewDatabaseRepo(PostgresDsn *DatabasePostgresDsn, MongoDb *DatabaseMongodbDsn) *Database {
	return &Database{
		DatabasePostgres: string(*PostgresDsn),
		DatabaseMongoDB:  string(*MongoDb),
	}
}
