package main

import (
	"log"

	"github.com/henriotrem/flashcard/internal/card"
	"github.com/henriotrem/flashcard/internal/db"
)

func Run() error {
	cardStore, err := db.New()
	if err != nil {
		log.Println("Error when starting the DB")
		return err
	}

	err = cardStore.Migrate()
	if err != nil {
		log.Println("Error when performing the DB migration")
		return err
	}

	_ = card.New(cardStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
