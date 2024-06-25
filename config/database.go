package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Postgres() *gorm.DB {
	createDirStorageLogsDatabase()
	file, err := os.OpenFile("./storage/logs/database/"+time.Now().Format("01-02-2006")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(fmt.Sprintf("error opening file database: %v", err))
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))
	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "omnichannel.",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func createDirStorageLogsDatabase() {
	dir := "./storage/logs/database"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Println(dir, "can't created directory")
		}
		fmt.Println("success created directory", dir)
	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
}
