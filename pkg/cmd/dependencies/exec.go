package dependencies

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ansible/aap-aoc-tools/pkg/graph"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	graphgonum "gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
)

func (o *Options) complete(cmd *cobra.Command, args []string) (err error) {
	for i, n := range o.Exclude {
		if !strings.HasPrefix(n, o.Prefix) {
			o.Exclude[i] = fmt.Sprintf("%s%s", o.Prefix, n)
		}
	}
	return nil
}

func (o *Options) validate() error {
	if len(o.ManifestFile) == 0 {
		if len(o.DeploymentName) == 0 || len(o.ManifestName) == 0 {
			return fmt.Errorf("manifest-file or the pair deployment/manifest are missing")
		}
		if (len(o.DeploymentName) != 0 && len(o.ManifestName) == 0) ||
			(len(o.DeploymentName) == 0 && len(o.ManifestName) != 0) {
			return fmt.Errorf("deployment and/or manifest are missing")
		} else {
			if v := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); len(v) == 0 {
				return fmt.Errorf("environment variable GOOGLE_APPLICATION_CREDENTIALS not defined")
			}
		}
	}

	if len(o.OutputFile) == 0 && o.GraphvizOptions != "-Ttxt" {
		return fmt.Errorf("an output-file must be provided")
	}
	return nil
}

func (o *Options) run() (err error) {
	var b []byte
	if len(o.ManifestFile) != 0 {
		b, err = os.ReadFile(o.ManifestFile)
		if err != nil {
			return err
		}
	} else {
		var out bytes.Buffer
		cmd := exec.Command("gcloud", "deployment-manager", "manifests", "describe", o.ManifestName, "--deployment", o.DeploymentName)
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			return err
		}
		b = out.Bytes()
	}
	manifest := make(map[string]interface{}, 0)
	err = yaml.Unmarshal(b, &manifest)
	if err != nil {
		return err
	}
	expendedConfigs := manifest["expandedConfig"].(string)
	expendedConfig := make(map[string]interface{}, 0)
	err = yaml.Unmarshal([]byte(expendedConfigs), &expendedConfig)
	if err != nil {
		return err
	}
	graph := graph.ReadGraph(expendedConfig, o.Exclude, o.Reverse)
	p, err := o.stringGraph(graph, o.StartResource)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(o.OutputFile) == 0 {
		fmt.Print(*p)
		return
	}
	switch o.GraphvizOptions {
	case "-Ttxt":
	default:
		var out bytes.Buffer
		cmd := exec.Command(o.GraphvizLayoutEngine, o.GraphvizOptions)
		cmd.Stdin = strings.NewReader(*p)
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			return err
		}
		s := out.String()
		p = &s
	}
	if err := os.WriteFile(o.OutputFile, []byte(*p), 0600); err != nil {
		return err
	}
	return nil
}

func (o *Options) stripDeploymentName(name string) string {
	return strings.TrimPrefix(name, o.Prefix)
}

func (o *Options) stringGraph(g graph.Graph, startPoint string) (*string, error) {
	p := fmt.Sprintln("digraph G {")
	if len(startPoint) == 0 {
		for _, e := range g.Edges {
			p += fmt.Sprintf("\"%s\" -> \"%s\";\n", o.stripDeploymentName(e.Start.Name), o.stripDeploymentName(e.End.Name))
		}
		p += fmt.Sprintln("}")
		return &p, nil
	}

	var startNode *graph.Node

	for _, n := range g.Nodes.Nodes {
		if o.stripDeploymentName(n.Name) == startPoint {
			startNode = &n
			break
		}
	}

	if startNode == nil {
		return nil, fmt.Errorf("start node %s doesn't exist", startPoint)
	}

	breadthFirst := traverse.BreadthFirst{
		Visit: func(n graphgonum.Node) {
			// fmt.Printf("ID=%v,Name=%s\n", n.ID(), o.getNodeName(g, n))
		},
		Traverse: func(e graphgonum.Edge) bool {
			p += fmt.Sprintf("\"%s\" -> \"%s\";\n", o.getNodeName(g, e.From()), o.getNodeName(g, e.To()))

			return true
		},
	}

	breadthFirst.Walk(g, startNode, func(n graphgonum.Node, d int) bool {
		return false
	})
	p += fmt.Sprintln("}")
	return &p, nil
}

func (o *Options) getNodeName(g graph.Graph, n graphgonum.Node) string {
	node := g.Node(n.ID())
	if node == nil {
		return ""
	}
	return o.stripDeploymentName(node.Name)
}
