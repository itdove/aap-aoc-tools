package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ansible/aap-aoc-tools/pkg/cmd/dependencies"
	"github.com/spf13/cobra"
)

var example = `
%[1]s dependencies 
       --prefix <prefix> --manifest-file <manifest-file> 
       [--exclude <resource_1>]... 
	   [--start-resource <start-resource>] 
	   [--output-file <output-file>] [--graph-format <graph-format>]
%[1]s dependencies 
       --prefix <prefix> --deployment <deployment> --manifest <manifest> 
	   [--exclude <resource_1>]... 
	   [--start-resource <start-resource>]
	   [--output-file <output-file>] [--graph-format <graph-format>]

deployment and manifest require gcloud to be installed and 
	the GOOGLE_APPLICATION_CREDENTIALS environement variable to be set.
graph-format can be found at https://graphviz.org/docs/outputs/
	and graphviz must be installed https://graphviz.org/download/
`

func main() {
	cmd := &cobra.Command{
		Short:        "Generate graph",
		Example:      fmt.Sprintf(example, os.Args[0]),
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()

	flags.AddGoFlagSet(flag.CommandLine)
	dependencies := dependencies.NewCmd(flags)
	cmd.AddCommand(dependencies)

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
