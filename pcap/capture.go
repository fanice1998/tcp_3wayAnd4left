package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// 觀測 3 次握手

func main() {
	const (
		// EthDev  string         = "Net device name"
		EthDev string = "\\Device\\NPF_{A2B1B76B-513F-4DE9-A112-EEA1C9ECEEE3}"
		SnapLen   int32          = 65536
		Promisc   bool           = false
		Timeout   time.Duration  = time.Second * 3
		Port      layers.TCPPort = 80
		ipAddress string         = "192.168.80.118"
	)

	log.Println("Capture dev: ", EthDev, "port: ", Port)
	handler, err := pcap.OpenLive(EthDev, SnapLen, Promisc, Timeout)
	if err != nil {
		log.Fatalln(err)
	}
	defer handler.Close()

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for pk := range source.Packets() {
		if layer2 := pk.NetworkLayer(); layer2 != nil {
			ipLayer, ok := layer2.(*layers.IPv4)
			if !ok {
				continue
			}
			if layer3 := pk.TransportLayer(); layer3 != nil {

				if tcpLayer, ok := layer3.(*layers.TCP); ok {
					if (fmt.Sprintf("%v", ipLayer.DstIP) == ipAddress || fmt.Sprintf("%v", ipLayer.SrcIP) == ipAddress) && (tcpLayer.DstPort == Port || tcpLayer.SrcPort == Port) {
						fmt.Printf("%s:%d --> %s:%d, SYN=%v, ACK=%v, RST=%v, PSH=%v, FIN=%v, URG=%v, ECE=%v,"+
							" Payload length=%d, seq=%v, ackNum=%v, windows=%v \n",
							ipLayer.SrcIP,
							tcpLayer.SrcPort,
							ipLayer.DstIP,
							tcpLayer.DstPort,
							tcpLayer.SYN, tcpLayer.ACK, tcpLayer.RST, tcpLayer.PSH, tcpLayer.FIN, tcpLayer.URG, tcpLayer.ECE,
							len(tcpLayer.Payload),
							tcpLayer.Seq,
							tcpLayer.ACK,
							tcpLayer.Window,
						)
					}
				}
			}
		}
	}
}
