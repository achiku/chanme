package main

import (
	"log"
	"os"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("chanme-wallet", "chanme wallet client.")

	createPrivKeyCmd = app.Command("createprivkey", "Create new private key.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case createPrivKeyCmd.FullCommand():
		if err := createPrivKey(); err != nil {
			app.Errorf("createPrivKey failed: %s", err)
		}
	default:
		log.Printf("no such command")
	}
	return
}
