package dependencies

type Options struct {
	Prefix         string
	ManifestFile   string
	ManifestName   string
	DeploymentName string
	Exclude        []string
	StartResource  string
	Format         string
	OutputFile     string
	Reverse        bool
}

func newOptions() *Options {
	return &Options{}
}
