package cmd

import (
	"fmt"

	"os"

	"github.com/blit/txtrocket/controllers/website"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// Server command
var Server = cli.Command{
	Name:        "server",
	ShortName:   "s",
	Usage:       "txtrocket server",
	Description: "Starts the txtrocket webserver",
	Action:      runServer,
}

func runServer(ctx *cli.Context) {
	router := mux.NewRouter()

	router.HandleFunc("/", website.Index)

	addr := os.Getenv("TXTR_SERVER_BIND")
	if addr == "" {
		addr = ":3000" // default port
	}

	fmt.Println("booting txtrocker server on ", addr)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(addr)
}
