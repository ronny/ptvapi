package main

import (
	"context"
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/ronny/ptvapi/client"
	"github.com/ronny/ptvapi/client/routes"
	"github.com/ronny/ptvapi/transport"
)

func main() {
	devID := os.Getenv("PTV_DEVID")
	if devID == "" {
		panic("Missing PTV_DEVID env var")
	}

	apiKey := os.Getenv("PTV_APIKEY")
	if apiKey == "" {
		panic("Missing PTV_APIKEY env var")
	}

	transport := transport.NewSigningTransport(
		httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes),
		devID,
		apiKey,
	)
	client := apiclient.New(transport, strfmt.Default)

	resp, err := client.Routes.RoutesOneOrMoreRoutes(&routes.RoutesOneOrMoreRoutesParams{
		Context: context.Background(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
}
