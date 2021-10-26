package db

import (
	"errors"
	"log"

	"github.com/henriotrem/flashcard/internal/card"
)

func (s Store) AddCard(cd card.Card) (card.Card, error) {
	_, err := s.db.NamedQuery(
		`INSERT INTO cards
		(id, content)
		VALUES (:id, :content)`,
		cd,
	)
	if err != nil {
		return card.Card{}, errors.New("failed to instert into database")
	}
	return card.Card{
		ID:      cd.ID,
		Content: cd.Content,
	}, nil
}

func (s Store) GetCard(id string) (card.Card, error) {
	var cd card.Card
	row := s.db.QueryRow(
		`SELECT id, content FROM cards WHERE id=$1;`,
		id,
	)
	err := row.Scan(&cd.ID, &cd.Content)
	if err != nil {
		log.Println(err.Error())
		return card.Card{}, err
	}
	return cd, nil
}

func (s Store) DeleteCard(id string) error {
	return nil
}
