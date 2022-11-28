package dependencies

type Options struct {
	Prefix               string
	ManifestFile         string
	ManifestName         string
	DeploymentName       string
	Exclude              []string
	StartResource        string
	GraphvizLayoutEngine string
	GraphvizFlags        string
	OutputFile           string
	Reverse              bool
	Project              string
}

func newOptions() *Options {
	return &Options{}
}
