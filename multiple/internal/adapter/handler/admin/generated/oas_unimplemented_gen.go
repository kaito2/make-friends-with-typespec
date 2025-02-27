// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// TenantClientsList implements TenantClients_list operation.
//
// GET /tenant-clients
func (UnimplementedHandler) TenantClientsList(ctx context.Context) (r []TenantClient, _ error) {
	return r, ht.ErrNotImplemented
}

// TenantClientsRead implements TenantClients_read operation.
//
// GET /tenant-clients/{id}
func (UnimplementedHandler) TenantClientsRead(ctx context.Context, params TenantClientsReadParams) (r *TenantClient, _ error) {
	return r, ht.ErrNotImplemented
}
