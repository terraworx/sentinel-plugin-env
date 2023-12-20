package main

import (
	"github.com/hashicorp/sentinel-sdk/rpc"
	env "github.com/terrawork/sentinel-env/plugin"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		PluginFunc: env.New,
	})
}
