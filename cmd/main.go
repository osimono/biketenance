package main

import (
	"embed"
	"github.com/gorilla/mux"
	"github.com/osimono/biketenance/cmd/bikes"
	"github.com/osimono/biketenance/cmd/db"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"time"
)

//go:embed ui
var embeddedFiles embed.FS

func main() {
	//db.UseDb()
	db := db.DB()
	defer db.Close()

	r := mux.NewRouter()
	subFS, _ := fs.Sub(embeddedFiles, "ui")
	r.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.FS(subFS))))

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/bikes", bikes.AllBikes)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("starting server...")
	log.Fatal(srv.ListenAndServe())
}
