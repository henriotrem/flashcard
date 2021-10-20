package db

import "github.com/henriotrem/flashcard/internal/card"

func (s Store) AddCard(cd card.Card) (card.Card, error) {
	return card.Card{}, nil
}

func (s Store) GetCard(id string) (card.Card, error) {
	return card.Card{}, nil
}

func (s Store) DeleteCard(id string) error {
	return nil
}
