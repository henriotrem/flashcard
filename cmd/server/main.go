package main

import (
	"log"

	"github.com/henriotrem/flashcard/internal/card"
	"github.com/henriotrem/flashcard/internal/db"
	"github.com/henriotrem/flashcard/internal/transport/grpc"
)

func Run() error {
	cardStore, err := db.New()
	if err != nil {
		return err
	}

	err = cardStore.Migrate()
	if err != nil {
		return err
	}

	cardService := card.New(cardStore)
	cardHandler := grpc.New(cardService)

	if err = cardHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
