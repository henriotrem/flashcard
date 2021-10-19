//go:generate mockgen -destination=card_mocks_test.go -package=card github.com/henriotrem/flashcard/internal/card Store
package card

import "context"

type Card struct {
	ID       string
	Label    string
	Type     string
	Content  string
	Created  int32
	Modified int32
}

type Store interface {
	AddCard(card Card) (Card, error)
	GetCard(id string) (Card, error)
	DeleteCard(id string) error
}

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) AddCard(context context.Context, card Card) (Card, error) {
	card, err := s.Store.AddCard(card)
	if err != nil {
		return Card{}, err
	}
	return card, nil
}

func (s Service) GetCard(context context.Context, id string) (Card, error) {
	card, err := s.Store.GetCard(id)
	if err != nil {
		return Card{}, err
	}
	return card, nil
}

func (s Service) DeleteCard(context context.Context, id string) error {
	err := s.Store.DeleteCard(id)
	if err != nil {
		return err
	}
	return nil
}
