package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
)

type Args struct {}

func (a Args) Parse() error {
	parser := flags.NewParser(nil, flags.Default)
	parser.AddCommand("run", "Runs the proxy", "Runs the proxy", &run)
	parser.AddCommand("reconfigure", "Reconfigures the proxy", "Reconfigures the proxy using information stored in Consul", &reconfigure)
	if _, err := parser.ParseArgs(os.Args[1:]); err != nil {
		return fmt.Errorf("Could not parse command line arguments\n%v", err)
	}
	return nil
}