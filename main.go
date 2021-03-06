package main

import (
	"github.com/wraix/shopit/cmd"
	"github.com/wraix/shopit/env"
)

var (
	name        string = "shopit"
	version     string = "0.0.0"
	environment string = "development"
	commit      string
	date        string
	tag         string
)

func init() {
	env.Env.Build.Name = name
	env.Env.Build.Version = version
	env.Env.Build.Tag = tag
	env.Env.Build.Commit = commit
	env.Env.Build.Date = date
	env.Env.Build.Environment = environment
}

func main() {
	cmd.Execute()
}
