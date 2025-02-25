package main

import (
	"log"

	"github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/admin"
	"github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/client"
	mhttp "github.com/kaito2/make-friends-with-typespec/multiple/internal/adapter/handler/http"
)

func main() {
	adminSvr, err := admin.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	clientSvr, err := client.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	server := mhttp.NewServer(adminSvr, clientSvr)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
