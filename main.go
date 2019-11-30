package main

import "github.com/cloudingcity/golab/cmd"

var version = "dev"

func main() {
	cmd.Version = version
	cmd.Execute()
}
