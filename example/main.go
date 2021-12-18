package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"ouraring"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func getCurrentDate() string {
	currentTime := time.Now()
	return currentTime.Format("2017-09-27")
}

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
	// Get the activity score for today
	start, end := getCurrentDate(), getCurrentDate()
	params := &ouraring.ActivityParams{
		Start: &start,
		End:   &end,
	}

	resp, err := client.ActivityWithResponse(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	// Since start date and end date are the same, we should only get 1 result
	fmt.Printf("Activity Score: %v\n", *resp.JSON200.Activity[0].Score)
}
