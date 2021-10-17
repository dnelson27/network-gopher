package main

import (
	"fmt"
	"net"
	"bufio"
	"os"

	"dnelson-infosec.com/network-gopher/graph"
	"dnelson-infosec.com/network-gopher/traversal"
	"dnelson-infosec.com/network-gopher/utils"
)

func main() {
	// Init graph
	g := graph.Graph{}
    input := bufio.NewScanner(os.Stdin)

	
	for {
		fmt.Print("> ")
		input.Scan()

		inputText := input.Text()

		// validate user input
		start, err := utils.GetCurrentIP()
		end := net.ParseIP(inputText)

		if err != nil {
			panic("Failed to grab current public IP address")
		}

		if end == nil {
			panic("Failed to parse given IP address")
		}

		// call traversal method
		g, err = traversal.Traverse(g, start, end)
		if err != nil {
			fmt.Printf("Error! %v\n", err)
			panic("Failed traversing network, ending")
			break
		}
		fmt.Println(g.Vertices)
	}

}