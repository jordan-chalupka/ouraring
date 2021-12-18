package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"ouraring"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func main() {
	var (
		withTrailingSlash string = "https://api.ouraring.com/v1"
	)
	apiKey, err := securityprovider.NewSecurityProviderApiKey("query", "access_token", os.Getenv("OURA_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	client, err := ouraring.NewClientWithResponses(
		withTrailingSlash,
		ouraring.WithRequestEditorFn(apiKey.Intercept))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	start, end := "2021-12-17", "2021-12-17"
	params := &ouraring.ActivityParams{
		Start: &start,
		End:   &end,
	}

	resp, err := client.ActivityWithResponse(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Activity Score: %v\n", *resp.JSON200.Activity[0].Score)
}
