package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	apiKeys struct {
		publicKey string
		secretKey string
	}
}

type application struct {
	config        config
	infolog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf("localhost:%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infolog.Println(fmt.Sprintf("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviorment {development|production}")

	flag.Parse()

	cfg.apiKeys.publicKey = "xnd_public_development_WEGjYBpLutnVZGqP4XilQ6OGzWO7pyzIrx1ytB58AwfVvkGr3bUn4pIOWsm3C331"
	cfg.apiKeys.secretKey = "xnd_development_WQP1xVMTB9liR4xziHLUMVFcVPNfTO20g7iOX8qanFhrAUyd7cIFXknYEDNMo83"

	info_log := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	error_log := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	template_cache := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infolog:       info_log,
		errorLog:      error_log,
		templateCache: template_cache,
		version:       version,
	}

	err := app.serve()

	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}

}
