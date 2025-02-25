package client

import (
	"context"

	"github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/client/generated"
)

type handler struct{}

// StoresList implements generated.Handler.
func (h *handler) StoresList(ctx context.Context) ([]generated.Store, error) {
	return []generated.Store{
		{
			Name: "example_store_1",
		},
		{
			Name: "example_store_2",
		},
	}, nil
}

// StoresRead implements generated.Handler.
func (h *handler) StoresRead(ctx context.Context, params generated.StoresReadParams) (*generated.Store, error) {
	return &generated.Store{
		Name: "example_store_1",
	}, nil
}

func NewService() generated.Handler {
	return &handler{}
}
