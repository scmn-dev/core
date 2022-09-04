package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/scmn-dev/core/app"
	"github.com/scmn-dev/core/config"
	"github.com/scmn-dev/core/constants"
	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/router"
)

func main() {
	logger := log.New(os.Stdout, "[sm-core] ", 0)

	cfg, err := config.Init(constants.ConfigPath, constants.ConfigName)

	if err != nil {
		log.Fatal(err)
	}

	_db, err := db.DBConn(&cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	s := db.New(_db)

	app.MigrateSystemTables(s)

	srv := &http.Server{
		MaxHeaderBytes: 10, // 10 MB
		Addr:           ":" + cfg.Server.Port,
		WriteTimeout:   time.Second * time.Duration(cfg.Server.Timeout),
		ReadTimeout:    time.Second * time.Duration(cfg.Server.Timeout),
		IdleTimeout:    time.Second * 60,
		Handler:        router.New(s),
	}

	logger.Printf("📡 Server listening on %s", cfg.Server.Port)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
