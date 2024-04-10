package database

import (
	"database/sql"
	"fmt"
	"module/internal/models"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

// путь до .env миграций
var envConnvertion string = "internal/config/postgres.env"
var migrationRoute string = "internal/database/migrations"

var GlobalInstance *PostgressCon

type PostgressCon struct {
	db *sqlx.DB
}

func (pg *PostgressCon) NewConnection() {

	connStr := getDBconnStr()
	poolConn, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	poolConn.SetMaxOpenConns(10)
	poolConn.SetMaxIdleConns(8)
	poolConn.SetConnMaxLifetime(20 * time.Second)

	pg.db = poolConn

	// defer poolConn.Close()

}

func (pg *PostgressCon) GetConnect() *sqlx.DB {
	return pg.db
}

func (pg *PostgressCon) PingReconnect() {
	duration := time.Second * 5
	time.Sleep(duration)

	if err := pg.db.Ping(); err != nil {
		fmt.Println("trying to reconnect")
		pg.NewConnection()
	}

	pg.PingReconnect()
}

func getDBconnStr() string {
	godotenv.Load(envConnvertion)

	envUser := os.Getenv("User")
	envPass := os.Getenv("Pass")
	envHost := os.Getenv("Host")
	envPort := os.Getenv("Port")
	envName := os.Getenv("Name")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envHost, envPort, envName)
	return connStr
}

// установка соединения с базой данных
func ConnectToDb() *sql.DB {

	connStr := getDBconnStr()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("database connection error")
		log.Debug("there is not connection with database")
		models.CheckError(err)
	}

	return db
}

// получить адрес внешнего сервера
func GetExternalRoutes(address *string) {
	godotenv.Load(envConnvertion)
	*address = os.Getenv("ExtAddress")
}

// начать миграцию
func MigrateStart() {

	duration := time.Second * 5
	time.Sleep(duration)

	db := ConnectToDb()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Debug("dont get migration dialect")
		models.CheckError(err)
	}

	if err := goose.Up(db, migrationRoute); err != nil {
		log.Error("migration connection error")
		log.Debug("no connection with database, or wrong migration route")
		models.CheckError(err)
	}
}
