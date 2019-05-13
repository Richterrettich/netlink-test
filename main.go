package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func main() {
	addrUpdates := make(chan netlink.AddrUpdate, 0)
	done := make(chan struct{}, 0)
	err := netlink.AddrSubscribeWithOptions(addrUpdates, done, netlink.AddrSubscribeOptions{
		ListExisting: true,
	})

	if err != nil {
		panic(err)
	}

	for update := range addrUpdates {
		fmt.Println(update.LinkAddress.String())
	}
}
