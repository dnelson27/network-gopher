package utils

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)


func GetCurrentIP () (net.IP, error) {
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	ipAddr := net.ParseIP(string(ip))

	if ipAddr == nil {
		return nil, fmt.Errorf("empty return from Ipify API")
	}

	return ipAddr, nil

}