package main

import (
	"fmt"
	// "net"
	"strconv"
	"strings"
	"regexp"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(fmt.Sprintf("Uncaught Error! %d", e))
	}
}

func getIpsListFromCidr(targetSubnet string, targetCidr int) []string {
	var result []string
	// finalIpAddress := getFinalIp() // Gets the last available IP
	// firstIpAddress := getFirstIp() // Gets the first available IP for this block
	// result = iterateIps(firstIpAddress, finalIpAddress) // Get all available IPs between first and last available IP in block
	return result
}

func scanCidrRange(targetSubnet string, targetCidr int){
	// targetList := getIpsListFromCidr()
	fmt.Printf("Scanning subnet %s with CIDR range /%d", targetSubnet, targetCidr)
}

func scanSingleIP(target string){
	fmt.Printf("Scanning single IP %s", target)
}

func validateIp(ipAddress string) bool {
	ipMatchString := `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	match, err := regexp.MatchString(ipMatchString, ipAddress)
	errCheck(err)
	if match {
		return true
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		panic("Invalid arguments!!")
	}
	targetIp := os.Args[1]
	cidrMatch, err := regexp.MatchString(`.*/.*`, targetIp)
	errCheck(err)
	if cidrMatch {
		resultList := strings.Split(targetIp, "/")
		targetSubnet := resultList[0]
		targetCidr, err := strconv.Atoi(resultList[1])
		errCheck(err)
		if validateIp(targetSubnet) && targetCidr <= 32 && targetCidr > 0 {
			scanCidrRange(targetSubnet, targetCidr)
		} else {
			panic(fmt.Sprintf("Invalid IP or CIDR mask %d", targetIp))
		}
	} else if validateIp(targetIp) {
		scanSingleIP(targetIp)
	} else {
		panic(fmt.Sprintf("Invalid IP or CIDR mask %d", targetIp))	
	}
}