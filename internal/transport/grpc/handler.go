package grpc

import (
	"context"
	"log"
	"net"

	"github.com/henriotrem/flashcard/internal/card"
	cd "github.com/henriotrem/protos/card/v1"
	"google.golang.org/grpc"
)

type CardService interface {
	AddCard(context context.Context, card card.Card) (card.Card, error)
	GetCard(context context.Context, id string) (card.Card, error)
	DeleteCard(context context.Context, id string) error
}

type Handler struct {
	CardService CardService
}

func New(cardService CardService) Handler {
	return Handler{
		CardService: cardService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	cd.RegisterCardServiceServer(grpcServer, &h)

	log.Println("gRPC Server listening on 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("Failed to serve")
		return err
	}

	return nil
}

func (h Handler) AddCard(context context.Context, req *cd.AddCardRequest) (*cd.AddCardResponse, error) {
	log.Println("Add Card Method")
	return &cd.AddCardResponse{
		Card: req.Card,
	}, nil
}

func (h Handler) GetCard(context context.Context, req *cd.GetCardRequest) (*cd.GetCardResponse, error) {
	log.Println("Get Card Method")
	return &cd.GetCardResponse{}, nil
}

func (h Handler) DeleteCard(context context.Context, req *cd.DeleteCardRequest) (*cd.DeleteCardResponse, error) {
	log.Println("Delete Card Method")
	return &cd.DeleteCardResponse{}, nil
}
