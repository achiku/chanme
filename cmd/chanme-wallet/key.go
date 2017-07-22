package main

import (
	"crypto/rand"
	"log"

	"github.com/achiku/chanme"
)

func createPrivKey() error {
	k, err := chanme.GenerateNewPrivateKey(rand.Reader)
	if err != nil {
		return err
	}
	log.Printf("%s", k.Base58Check(chanme.VerMainPriv))
	return nil
}
