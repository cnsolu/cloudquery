package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	p := provider.Provider()
	serve.Serve(&serve.Options{
		Name:     p.Name,
		Provider: p,
	})
}
