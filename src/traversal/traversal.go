package traversal

import (
	"net"
	"dnelson-infosec.com/network-gopher/graph"
	"syscall"
	"errors"
	"fmt"
)


func getDestAddr(dest net.IP) ([4]byte, error) {
	// Placeholder bytes
    destAddr := [4]byte{0, 0, 0, 0}

	// Get net.IPAddr and check host resolves properly
    ipAddr, err := net.ResolveIPAddr("ip", dest.String())

    if err != nil {
        return destAddr, err
    }

	// Copy IP address bytes to placeholder
    copy(destAddr[:], ipAddr.IP.To4())
    return destAddr, nil
}

func getSocketAddr() ([4]byte, error) {
    socketAddr := [4]byte{0, 0, 0, 0}

	// Get current system's interface addresses
    addrs, err := net.InterfaceAddrs()

    if err != nil {
        return socketAddr, err
    }

	// iterate over addrs: []net.Addr
    for _, addr := range addrs {

		// Check if address interface is of type net.IPNet, and it is not a loopback address
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

			// If the 4-byte representation of the ipnet is a valid IPv4 address
            if len(ipnet.IP.To4()) == net.IPv4len {

				// Move the 4-byte representation of the IPnet into all values of the socketAddr
                copy(socketAddr[:], ipnet.IP.To4())
                return socketAddr, nil
            }
        }
    }

	// If there are no addresses or an error was thrown
    err = errors.New("internet connection failed")
    return socketAddr, err
}



func Traverse(g graph.Graph, destIp net.IP) (graph.Graph, error) {
	/*
	Send a UDP packet to the destination, with a TTL iterating up from 1, recording the `from` address ICMP `Time Limit Exceeded`
	Save the previous hop as `start`, and the new hop as `end` in the graph. If the new hop == destination, exit the loop
	*/

	// 2000 Ms timeout, from nanoseconds
	timeoutMs := int64(2000)
	timeoutValue := syscall.NsecToTimeval(timeoutMs * int64(1_000_000))

	// Create an IPV4 UDP socket for  the local interface
	sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	defer syscall.Close(sendSocket)

	if err != nil {
		return g, err
	}	

	// Create an IPV4 ICMP socket for receiving timeout messages
	recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	defer syscall.Close(recvSocket)

	if err != nil {
		return g, err
	}	

	// Get local interface address
	socketAddr, err := getSocketAddr()
	if err != nil {
		return g, err
	}	

	destAddr, err := getDestAddr(destIp)

	fmt.Println(destAddr)
	ttl := 1
	if err != nil {
		return g, err
	}	

	// previousVertex := g.SaveNewVertex()
	
	for {
		// Update the sendsocket's time-to-live
		syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)

		// RESET the recvSocket's timeout value
		syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVBUF, &timeoutValue)
		
		syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: 33434, Addr: socketAddr})
		syscall.Sendto(sendSocket, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: 33434, Addr: destAddr})
		
		// Define packet size
		packetSize := make([]byte, int(56))

		// Receive ICMP packet
		// (n int, from Sockaddr, err error)
        _, from, err := syscall.Recvfrom(recvSocket, packetSize, 0)
		
		if err != nil {
			return g, err
		}

        ip := from.(*syscall.SockaddrInet4).Addr
        ipString := fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])

		// Get reverse lookup for a given address
		// Returns (names []string, err error)
        // dnsNames, err := net.LookupAddr(ipString)
		fmt.Println(ipString)

		// Will check if a matching vertex already exists
		// newVertex = g.SaveNewVertex(net.ParseIP(ipString))

		// previousVertex.OutboundEdges = newVertex
		// newVertex.InboundEdges = previousVertex
		// previousVertex = newVertex
 
        // We stop our loop if we reach destination or reach max value for ttl
        if ipString == destIp.String() || ttl >= 56 { 
			return g, nil
        }

        ttl += 1
	}
}
