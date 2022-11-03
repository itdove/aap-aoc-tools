package dependencies

type Options struct {
	Prefix               string
	ManifestFile         string
	ManifestName         string
	DeploymentName       string
	Exclude              []string
	StartResource        string
	GraphvizLayoutEngine string
	GraphvizOptions      string
	OutputFile           string
	Reverse              bool
}

func newOptions() *Options {
	return &Options{}
}
