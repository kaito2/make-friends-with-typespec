package admin

import (
	"fmt"

	admingen "github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/admin/generated"
)

func NewServer() (*admingen.Server, error) {
	svc := NewService()
	svr, err := admingen.NewServer(svc)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	return svr, nil
}
