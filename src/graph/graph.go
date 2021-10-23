package graph

import "net"

type Graph struct {
	Vertices map[string]Vertex
}

type Vertex struct {
	VertexIP net.IP
	Edges    map[string]Edge
}

type Edge struct {
	To   Vertex
	From Vertex
}

func (g *Graph) SaveNewVertex(ip net.IP) Vertex {

	// add vertex to graph
	vertex := Vertex{VertexIP: ip}

	// Init the vertices map
	if g.Vertices == nil {
		g.Vertices = make(map[string]Vertex)
	}

	g.Vertices[ip.String()] = vertex
	vertex.Edges = make(map[string]Edge)

	return vertex
}

func (g *Graph) ConnectVertices(toAddr, fromAddr net.IP) {
	var toVertex Vertex
	var fromVertex Vertex
	toAddrString := toAddr.String()
	fromAddrString := fromAddr.String()

	if value, ok := g.Vertices[toAddrString]; ok {
		toVertex = value
	} else {
		toVertex = g.SaveNewVertex(toAddr)
	}

	if value, ok := g.Vertices[fromAddrString]; ok {
		fromVertex = value
	} else {
		fromVertex = g.SaveNewVertex(fromAddr)
	}

	// (from: Vertex)-[:Edge]->(to: Vertex)
	fromVertex.Edges[toAddrString] = Edge{To: fromVertex, From: toVertex}

}
