package graph

import "net"

type Graph struct {
	Vertices map[string]Vertex
}

type Vertex struct {
	PublicIP net.IP
	Edges    map[string]Edge
}

type Edge struct {
	To   Vertex
	From Vertex
}

func (g *Graph) SaveNewVertex(publicIp net.IP) (Vertex, error) {

	// add vertex to graph
	vertex := Vertex{PublicIP: publicIp}

	// Init the vertices map
	if g.Vertices == nil {
		g.Vertices = make(map[string]Vertex)
	}

	g.Vertices[publicIp.String()] = vertex
	vertex.Edges = make(map[string]Edge)

	return vertex, nil
}

func (g *Graph) ConnectVertices(toAddr, fromAddr net.IP) error {
	var err error
	var toVertex Vertex
	var fromVertex Vertex
	toAddrString := toAddr.String()
	fromAddrString := fromAddr.String()

	if value, ok := g.Vertices[toAddrString]; ok {
		toVertex = value
	} else {
		toVertex, err = g.SaveNewVertex(toAddr)
	}

	if value, ok := g.Vertices[fromAddrString]; ok {
		fromVertex = value
	} else {
		fromVertex, err = g.SaveNewVertex(fromAddr)
	}

	if err != nil {
		return err
	}

	// (from: Vertex)-[:Edge]->(to: Vertex)
	fromVertex.Edges[toAddrString] = Edge{To: fromVertex, From: toVertex}

	return nil

}
