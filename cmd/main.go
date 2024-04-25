package main

import (
	"github.com/haapjari/repository-database-api/internal/pkg/cfg"
	"github.com/haapjari/repository-database-api/internal/pkg/database"
	"github.com/haapjari/repository-database-api/internal/pkg/handler"
	"log/slog"
	"net/http"
	"os"
)

const (
	host = "0.0.0.0"
)

func main() {
	var (
		conf = cfg.NewConfig()
		db   = database.NewDatabase(&database.Options{
			Host:     conf.DBHost,
			Port:     conf.DBPort,
			Username: conf.DBUser,
			Password: conf.DBPassword,
			Database: conf.DBName,
			TimeZone: conf.DBTimeZone,
		})

		mux = http.NewServeMux()
	)

	d, err := db.ConnectAndMigrate()
	if err != nil {
		panic("unable to connect to the database: " + err.Error())
	}

	h := handler.NewHandler(conf, d)
	h.SetupRoutes(mux)

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	slog.Info("REST API | " + host + ":" + conf.Port)

	if err = http.ListenAndServe(host+":"+conf.Port, mux); err != nil {
		panic("unable to start the server: " + err.Error())
	}
}
