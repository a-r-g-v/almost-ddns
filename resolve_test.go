package main

import (
	"fmt"
	"testing"
)

func TestResolve(t *testing.T) {

	ipaddr, err := resolve("www.arg.vc", "ns.arg.vc")

	if err != nil {
		fmt.Println(err)
		t.Error("unexpected to resolve error")
	}

	if ipaddr.String() != "127.0.0.1" {
		t.Error("www.arg.vc a-record is not 127.0.0.1 get:", ipaddr.String())
	}

}

func TestNxdomainResolve(t *testing.T) {

	ipaddr, err := resolve("a20482.arg.vc", "ns.arg.vc") // nxdomain

	if err != nil {
		fmt.Println(err)
		t.Error("unexpected to resolve error")
	}

	if ipaddr != nil {
		t.Error("When resolve NXDomain, unexpected to !nil ")
	}

}
