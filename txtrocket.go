package main

import (
	"os"
	"runtime"

	"github.com/blit/txtrocket/cmd"
	"github.com/codegangsta/cli"
	"github.com/joho/godotenv"
)

// Txtrocket version
const VERSION = "0.0.1 Beta"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// load .env
	godotenv.Load()
	// initialize app cli
	app := cli.NewApp()
	app.Name = "txtrocket"
	app.Usage = "hooah"
	app.Version = VERSION
	app.Commands = []cli.Command{
		cmd.Server,
		//cmd.Sms,
	}
	app.Run(os.Args)
}
