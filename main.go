package main

import "line-halo-stats/cloud_run/infrastructure/waf/server"

func main() {
	server.NewServer().Run()
}
