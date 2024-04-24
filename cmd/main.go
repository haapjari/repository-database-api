package main

import (
	"github.com/haapjari/repository-database-api/internal/pkg/cfg"
	"github.com/haapjari/repository-database-api/internal/pkg/handler"
	"log/slog"
	"net/http"
	"os"
)

const (
	host = "0.0.0.0"
)

func main() {
	conf := cfg.NewConfig()

	h := handler.NewHandler(conf)
	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	slog.Info("REST API | " + host + ":" + conf.Port)

	err := http.ListenAndServe(host+":"+conf.Port, mux)
	if err != nil {
		panic("unable to start the server: " + err.Error())
	}
}
