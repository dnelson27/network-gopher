package traversal

import (
	"net"
	// "dnelson-infosec.com/network-gopher/graph"
	"syscall"
	"errors"
	"fmt"
)

func destAddr(dest string) ([4]byte, error) {
    destAddr := [4]byte{0, 0, 0, 0}
    addrs, err := net.LookupHost(dest)
    if err != nil {
        return destAddr, err
    }
    addr := addrs[0]

    ipAddr, err := net.ResolveIPAddr("ip", addr)
    if err != nil {
        return destAddr, err
    }
    copy(destAddr[:], ipAddr.IP.To4())
    return destAddr, nil
}

func socketAddr() ([4]byte, error) {
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



func Trace(g graph.Graph, destIp net.IP) (graph.Graph, error) {
	/*
	Send a UDP packet to the destination, with a TTL iterating up from 1, recording the `from` address ICMP `Time Limit Exceeded`
	Save the previous hop as `start`, and the new hop as `end` in the graph. If the new hop == destination, exit the loop
	*/
	var ttl int

	// 2000 Ms timeout
	timeoutValue := syscall.NsecToTimeval(1000 * 1000 * (int64)(2000))

	// Create an IPV4 UDP socket for  
	sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	defer syscall.Close(sendSocket)

	if err != nil {
		return g, err
	}	

	// Create an IPV4 ICMP socket for receiving timeout messages
	recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_IP)
	defer syscall.Close(recvSocket)

	if err != nil {
		return g, err
	}	

	// Update the sendsocket's time-to-live
	syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)

	// Update the recvSocket's timeout value
	syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVBUF, &timeoutValue)

	destAddr := [4]byte{0, 0, 0, 0}
	copy(destAddr[:], destIp.To4())

	syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: 33434, Addr: socketAddr})
	syscall.Sendto(sendSocket, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: 33434, Addr: destAddr})

	/* TODO 
		- Finish this tracert functionality
			- Test and verify ICMP responses are being properly handled
			- Figure out the best way to update the graph on-the-fly
	*/

	g.ConnectVertices(start, end)
	return g, nil
}
