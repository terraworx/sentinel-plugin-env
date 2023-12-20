package main

import (
	"github.com/hashicorp/sentinel-sdk/rpc"
	"github.com/sentinel-plugin-file/env"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		PluginFunc: env.New,
	})
}