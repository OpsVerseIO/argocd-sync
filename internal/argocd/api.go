package argocd

import (
	"context"
	"io"
	"log"

	argocdclient "github.com/argoproj/argo-cd/v2/pkg/apiclient"
	applicationpkg "github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	argoio "github.com/argoproj/argo-cd/v2/util/io"
)

// Interface is an interface for API.
//
//nolint:lll // go generate is ugly.
type Interface interface {
	Sync(appName string) error
}

// API is struct for ArgoCD api.
type API struct {
	client     applicationpkg.ApplicationServiceClient
	connection io.Closer
}

// APIOptions is options for API.
type APIOptions struct {
	Address                string
	Token                  string
	DisableTlsVerification bool
}

// NewAPI creates new API.
func NewAPI(options *APIOptions) API {
	clientOptions := argocdclient.ClientOptions{
		ServerAddr: options.Address,
		AuthToken:  options.Token,
		GRPCWeb:    true,
		Insecure:   options.DisableTlsVerification,
	}

	connection, client := argocdclient.NewClientOrDie(&clientOptions).NewApplicationClientOrDie()
	if options.DisableTlsVerification {
		log.Println("Skip TLS Validation option is enabled")
	}

	return API{client: client, connection: connection}
}

// Sync syncs given application.
func (a API) Sync(appName string) error {
	request := applicationpkg.ApplicationSyncRequest{
		Name:  &appName,
		Prune: true,
	}

	_, err := a.client.Sync(context.Background(), &request)
	if err != nil {
		return err
	}

	defer argoio.Close(a.connection)

	return nil
}
