package admin

import (
	"context"

	"github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/admin/generated"
)

type handler struct{}

// TenantClientsList implements generated.Handler.
func (h *handler) TenantClientsList(ctx context.Context) ([]generated.TenantClient, error) {
	return []generated.TenantClient{
		{
			Name: "example_tenant_client_1",
		},
		{
			Name: "example_tenant_client_2",
		},
	}, nil
}

// TenantClientsRead implements generated.Handler.
func (h *handler) TenantClientsRead(ctx context.Context, params generated.TenantClientsReadParams) (*generated.TenantClient, error) {
	return &generated.TenantClient{
		Name: "example_tenant_client_1",
	}, nil
}

func NewService() generated.Handler {
	return &handler{}
}
