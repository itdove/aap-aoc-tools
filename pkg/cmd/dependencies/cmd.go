package dependencies

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var example = `
%[1]s --prefix <prefix> --manifest-file <manifest-file> 
       [--exclude <resource_1>]... 
	   [--start-resource <start-resource>] 
	   [--output-file <output-file>] [--graph-format <graph-format>]
%[1]s --prefix <prefix> --deployment <deployment> --manifest <manifest> 
	   [--exclude <resource_1>]... 
	   [--start-resource <start-resource>]
	   [--output-file <output-file>] [--graph-format <graph-format>]

deployment and manifest require gcloud to be installed and 
	the GOOGLE_APPLICATION_CREDENTIALS environement variable to be set.
graph-format can be found at https://graphviz.org/docs/outputs/
	and graphviz must be installed https://graphviz.org/download/
`

func NewCmd(flags *pflag.FlagSet) *cobra.Command {
	o := newOptions()
	cmd := &cobra.Command{
		Use:          "dependencies",
		Aliases:      []string{"dep"},
		Short:        "Generate dependencies graph",
		Example:      fmt.Sprintf(example, os.Args[0]),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.complete(c, args); err != nil {
				return err
			}
			if err := o.validate(); err != nil {
				return err
			}
			if err := o.run(); err != nil {
				return err
			}
			return nil
		},
	}

	flags.StringVar(&o.Prefix, "prefix", "", "The resource prefix to strip")
	flags.StringVar(&o.ManifestFile, "manifest-file", "", "The manifest file")
	flags.StringVar(&o.ManifestName, "manifest", "", "The manifest name")
	flags.StringVar(&o.DeploymentName, "deployment", "", "The deployment name")
	flags.StringArrayVar(&o.Exclude, "exclude", nil, "The resources to exclude")
	flags.StringVar(&o.StartResource, "start-resource", "", "The start resource")
	flags.StringVar(&o.GraphFormat, "graph-format", "txt", "The graph format type")
	flags.StringVar(&o.OutputFile, "output-file", "", "The graph output file")
	flags.BoolVar(&o.Reverse, "reverse", false, "Reverse dependencies")

	return cmd

}
