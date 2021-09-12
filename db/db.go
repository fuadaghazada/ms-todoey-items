package db

import (
	"database/sql"
	"fmt"
	"github.com/fuadaghazada/ms-todoey-items/config"
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

var pgDb *pg.DB

func ConnectDb() *pg.DB {
	opts := getDBProperties()

	pgDb = pg.Connect(&opts)
	if pgDb == nil {
		log.Fatal("ActionLog.ConnectDb.error: Failed to connect")
	}

	return pgDb
}

func getDBProperties() pg.Options {
	dbTimeout, err := time.ParseDuration(config.Props.DbConnectionTimeout)
	if err != nil {
		log.Warn("ActionLog.getDBProperties.warn: Failed to parse timeout duration value")
		dbTimeout = time.Minute
	}

	opts := pg.Options{
		Addr:        config.Props.DbUrl,
		User:        config.Props.DbUser,
		Password:    config.Props.DbPass,
		Database:    config.Props.DbName,
		PoolSize:    config.Props.DbPoolSize,
		DialTimeout: dbTimeout,
		MaxRetries:  2,
		MaxConnAge:  time.Minute * 15,
	}

	return opts
}

func MigrateDb() {
	log.Debug("MigrateDb.Start")

	opts := getDBProperties()

	host := opts.Addr[:strings.LastIndex(opts.Addr, ":")]
	port, err := strconv.Atoi(opts.Addr[strings.LastIndex(opts.Addr, ":")+1:])
	if err != nil {
		log.Fatal("ActionLog.MigrateDb.error: ParsePort", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s %s=%s dbname=%s sslmode=disable",
		host, port, opts.User, "password", opts.Password, opts.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("ActionLog.MigrateDb.error: OpenConnection: ", err)
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("ActionLog.MigrateDb.error: Apply", err)
	}

	log.Debug("ActionLog.MigrateDb.End")
}