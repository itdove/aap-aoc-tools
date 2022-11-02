package graph

import (
	"math"
	"math/rand"

	graphgonum "gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
)

type Graph struct {
	Nodes Nodes
	Edges []Edge
}

type Nodes struct {
	Nodes []Node
	Index int
}

type Node struct {
	Name string
	UID  int64
}

type Edge struct {
	Start Node
	End   Node
}

var _ traverse.Graph = &Graph{}
var _ graphgonum.Nodes = &Nodes{}
var _ graphgonum.Node = &Node{}
var _ graphgonum.Edge = &Edge{}

func (g Graph) From(id int64) graphgonum.Nodes {
	nodes := &Nodes{
		Nodes: make([]Node, 0),
		Index: 0,
	}
	for _, e := range g.Edges {
		if e.Start.UID == id {
			nodes.Nodes = append(nodes.Nodes, e.End)
		}
	}
	return nodes
}

func (g Graph) Edge(uid int64, vid int64) graphgonum.Edge {
	for _, e := range g.Edges {
		if e.Start.UID == uid &&
			e.End.UID == vid {
			return e
		}
	}
	return nil
}

func (g Graph) Node(uid int64) *Node {
	for _, n := range g.Nodes.Nodes {
		if n.UID == uid {
			return &n
		}
	}
	return nil
}

func (n *Nodes) Len() int {
	return len(n.Nodes)
}

func (n *Nodes) Next() bool {
	n.Index += 1
	return n.Index < n.Len()
}

func (n *Nodes) Reset() {
	n.Index = 0
}

func (n Nodes) Node() graphgonum.Node {
	return n.Nodes[n.Index]
}

func (n Node) ID() int64 {
	return n.UID
}

func (n Node) GetName() string {
	return n.Name
}

func (e Edge) From() graphgonum.Node {
	return e.Start
}

func (e Edge) To() graphgonum.Node {
	return e.End
}

func (e Edge) ReversedEdge() graphgonum.Edge {
	return e
}

func ReadGraph(expendedConfig map[string]interface{}, exclude []string, reverse bool) (graph Graph) {
	graph = Graph{
		Edges: make([]Edge, 0),
	}
	mainResourcesI := expendedConfig["resources"]
	mainResources := mainResourcesI.([]interface{})
	for _, r := range mainResources {
		resource := r.(map[string]interface{})
		if metadataI, ok := resource["metadata"]; ok {
			metadata := metadataI.(map[string]interface{})
			nameI := resource["name"]
			name := nameI.(string)
			node := graph.Nodes.getNode(name)
			if node == nil {
				node = &Node{
					Name: name,
					UID:  rand.Int63n(math.MaxInt64),
				}
			}
			graph.Nodes.Nodes = append(graph.Nodes.Nodes, *node)
			if dependsOnsI, ok := metadata["dependsOn"]; ok {
				dependsOns := dependsOnsI.([]interface{})
				if excluded(exclude, name) {
					continue
				}
				for _, dependsOn := range dependsOns {
					dependsOnValue := dependsOn.(string)
					if excluded(exclude, dependsOnValue) {
						continue
					}
					toNode := graph.Nodes.getNode(dependsOnValue)
					if toNode == nil {
						toNode = &Node{
							Name: dependsOnValue,
							UID:  rand.Int63n(math.MaxInt64),
						}
						graph.Nodes.Nodes = append(graph.Nodes.Nodes, *toNode)
					}
					e := Edge{
						Start: *node,
						End:   *toNode,
					}
					if reverse {
						start := e.Start
						e.Start = e.End
						e.End = start
					}
					graph.Edges = append(graph.Edges, e)
				}
			}
		}

	}
	return
}

func excluded(excluded []string, name string) bool {
	for _, n := range excluded {
		if n == name {
			return true
		}
	}
	return false
}

func (n *Nodes) getNode(name string) *Node {
	for _, n := range n.Nodes {
		if n.Name == name {
			return &n
		}
	}
	return nil
}
