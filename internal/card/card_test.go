package card

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCardService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("ADD CARD TESTS", func(t *testing.T) {
		cardMockStore := NewMockStore(mockCtrl)
		id := "UUID-1"
		cardMockStore.
			EXPECT().
			AddCard(Card{
				ID: id,
			}).Return(Card{
			ID: id,
		}, nil)

		cardService := New(cardMockStore)
		cd, err := cardService.AddCard(
			context.Background(),
			Card{
				ID: id,
			})

		assert.NoError(t, err)
		assert.Equal(t, cd.ID, id)
	})
}
