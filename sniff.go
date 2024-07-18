package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)


func main() {
	var device string // if you don't know your interface name run interface.go

	handle,err := pcap.OpenLive(device, int32(1024),false,pcap.BlockForever)
	
	if err != nil{
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle,handle.LinkType())

	
	
	for packet := range packetSource.Packets(){
		fmt.Println(packet)
				
	}
	
}