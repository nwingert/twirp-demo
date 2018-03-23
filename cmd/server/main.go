package main

import (
	"net/http"

	"github.com/nwingert/twirp-demo/internal/tagbuilderserver"
	"github.com/nwingert/twirp-demo/rpc/tagbuilder"
)

func main() {
	server := &tagbuilderserver.Server{}
	twirpHandler := tagbuilder.NewTagBuilderServer(server, nil)

	http.ListenAndServe(":8085", twirpHandler)
}
