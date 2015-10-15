package main

import (
	"net"
	"os/exec"
	"strings"
)

func resolve(targetDomain string, nameServerDomain string) (net.IP, error) {

	// dig targetDomain +norec @nameServerDomain | head 1
	out, err := exec.Command("/bin/bash", "hack/resolve.sh", targetDomain, nameServerDomain).Output()

	if err != nil {
		return nil, err
	}

	ipaddr := string(out[:])
	ipaddr = strings.TrimRight(ipaddr, "\n")
	resolveIP := net.ParseIP(ipaddr)

	return resolveIP, nil

}
