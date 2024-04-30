package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"

	"github.com/jmoiron/sqlx"
	"github.com/robowealth-mutual-fund/stdlog"
	go_ora "github.com/sijms/go-ora/v2"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
)

type DB struct {
	Sql        *sql.DB
	Sqlx       *sqlx.DB
	Connection string
}

func New(conf config.Config) *DB {

	var (
		hosts        = conf.Database.Host
		mainHost     = hosts[0]
		clusterHosts []string
	)

	for _, host := range hosts[1:] {
		clusterHosts = append(clusterHosts, fmt.Sprintf("%s:%s", host, conf.Database.Port))
	}

	urlOptions := map[string]string{
		"server": strings.Join(clusterHosts, ","),
	}

	port, _ := strconv.Atoi(conf.Database.Port)
	connStr := go_ora.BuildUrl(mainHost, port, conf.Database.DatabaseName, conf.Database.User, conf.Database.Password, urlOptions)

	stdlog.DebugWithAttrs("connecting to Oracle db (sqlx + go-ora):", map[string]interface{}{
		"url": connStr,
	})

	dbx, err := otelsqlx.Open("oracle", connStr,
		otelsql.WithAttributes(semconv.DBSystemOracle),
		otelsql.WithDBName(conf.Database.DatabaseName))
	if err != nil {
		stdlog.Error("failed to connect to Oracle db", err)
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	err = dbx.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

	dbx.SetConnMaxLifetime(time.Second * time.Duration(conf.Database.ConnMaxLifetime))
	dbx.SetMaxOpenConns(conf.Database.MaxConnection)
	dbx.SetMaxIdleConns(conf.Database.MaxIdleConnection)

	//set named mapping
	sqlx.BindDriver("oracle", sqlx.NAMED)

	return &DB{
		Sql:        dbx.DB,
		Sqlx:       dbx,
		Connection: connStr,
	}
}
