package http

import (
	"log"
	"net/http"

	admin "github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/admin/generated"
	client "github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/client/generated"
)

type AppServer struct {
	adminSvr  *admin.Server
	clientSvr *client.Server
}

func NewServer(adminSvr *admin.Server, clientSvr *client.Server) *AppServer {
	return &AppServer{
		adminSvr:  adminSvr,
		clientSvr: clientSvr,
	}
}

func (s *AppServer) Start() error {
	mux := http.NewServeMux()

	mux.Handle("/api/admin/v1/", http.StripPrefix("/api/admin/v1", s.adminSvr))
	mux.Handle("/api/client/v1/", http.StripPrefix("/api/client/v1", s.clientSvr))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}
	return nil
}
