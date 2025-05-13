package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ShubhamkumarAnand/melkey-go/mel_project/internal/app"
	"github.com/ShubhamkumarAnand/melkey-go/mel_project/internal/routes"
)

func main() {
	// command line parsing for the port given by user
	// eg. -port 8081
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	// instance the new application
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	r := routes.SetupRoutes(app)

	// creating a server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Up and Running at Port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
