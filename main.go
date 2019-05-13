package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func main() {

	interfaces, err := netlink.LinkList()
	if err != nil {
		panic(err)
	}

	addrUpdates := make(chan netlink.AddrUpdate, 0)
	done := make(chan struct{}, 0)
	err = netlink.AddrSubscribeWithOptions(addrUpdates, done, netlink.AddrSubscribeOptions{
		ListExisting: true,
	})

	if err != nil {
		panic(err)
	}

	for update := range addrUpdates {
		fmt.Println(interfaces[update.LinkIndex].Attrs().Name, update.LinkAddress.String())
	}
}
