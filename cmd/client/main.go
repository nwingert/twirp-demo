package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/nwingert/twirp-demo/rpc/tagbuilder"
)

func main() {
	client := tagbuilder.NewTagBuilderProtobufClient("http://localhost:8085", &http.Client{})

	tagContainer, err := client.GetTags(context.Background(), &tagbuilder.Merchant{MerchantId: 2305})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Tag container received: %v", tagContainer.Tags)
}
