# Subnet mapper in Go with Graph visualization
This project is a work-in-progress. **Network Gopher** is a small tool used to record the path of traffic and map it to a graph database. I am doing two implementations, one with https://github.com/cayleygraph and one with Neo4j. I am using this to evaluate a different Graph database engine (Neo4j is my daily driver) and work on my skills with Go and Graph Databases.

I am taking inspiration and ideas from:
- https://github.com/caffix/netmap
- https://github.com/kalbhor/tracesite

## TO DO
- Finish traceroute functionality (verify ICMP responses from each hop are being properly handled)
- Get some basic stdout output for the graph
- Figure out the visualization part


