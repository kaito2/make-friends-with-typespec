package client

import (
	"fmt"

	clientgen "github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/client/generated"
)

func NewServer() (*clientgen.Server, error) {
	svc := NewService()
	svr, err := clientgen.NewServer(svc)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	return svr, nil
}
