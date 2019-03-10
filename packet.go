package main

import(
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func forgeReset(packet gopacket.Packet) gopacket.Packet{
	ethLayer := packet.Layer(layers.LayerTypeEthernet)
	eth, _ := ethLayer.(*layers.Ethernet)

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	ip, _ := ipLayer.(*layers.IPv4)

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	tcp, _ := tcpLayer.(*layers.TCP)			

	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths: true,
	}

	tcp.RST = true
	tcp.NS = false
	tcp.CWR = false
	tcp.ECE = false
	tcp.URG = false
	tcp.PSH = false
	tcp.SYN = false
	tcp.FIN = false	

	if(tcp.ACK){		
		tcp.ACK = false
		tcp.Seq = tcp.Ack
		tcp.Ack = 0		
	}	else{
		tcp.ACK = true
		tcp.Ack = tcp.Seq + uint32(len(packet.Data()))
		tcp.Seq = 0
		}			
		
	tcp.Window = 0
	tcp.Urgent = 0
	tcp.Options = tcp.Options[:0]
	tcp.Payload = tcp.Payload[:0]

	tcp.SrcPort, tcp.DstPort = tcp.DstPort, tcp.SrcPort			
	ip.SrcIP, ip.DstIP = ip.DstIP, ip.SrcIP
	eth.SrcMAC, eth.DstMAC = eth.DstMAC, eth.SrcMAC

	tcp.SetNetworkLayerForChecksum(ip)

	resetPacketBuffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializePacket(resetPacketBuffer, options, packet)
	if err != nil {
		panic(err)
	}
	resetPacket := gopacket.NewPacket(resetPacketBuffer.Bytes(), layers.LayerTypeEthernet, gopacket.Default)
	return resetPacket
}