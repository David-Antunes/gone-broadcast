package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/David-Antunes/gone-broadcast/internal"
	"github.com/seancfoley/ipaddress-go/ipaddr"
)

var rttLog = log.New(os.Stdout, "RTT INFO: ", log.Ltime)
var DEFAULT_PORT = ":8000"
var DEFAULT_IFACE = "eth0"

func main() {
	var err error
	var ief *net.Interface
	if ief, err = net.InterfaceByName(DEFAULT_IFACE); err != nil {
		panic(err)
	}

	var addrs []net.Addr
	if addrs, err = ief.Addrs(); err != nil {
		panic(err)
	}

	var ip net.IP
	var ipNet *net.IPNet

	if ip, ipNet, err = net.ParseCIDR(addrs[0].String()); err != nil {
		panic(err)
	}

	var bcast string

	if adr, err := ipaddr.NewIPAddressFromNetIPNet(ipNet); err != nil {
		panic(err)
	} else {
		b, _ := adr.ToIPv4().ToBroadcastAddress()
		bcast = b.GetNetIPAddr().IP.String() + DEFAULT_PORT
	}

	rttLog.Println("IP address:", ip)

	rttLog.Println("Broadcast Address:", bcast)

	listenAddr, err := net.ResolveUDPAddr("udp4", DEFAULT_PORT)

	if err != nil {
		panic(err)
	}

	rttLog.Println(listenAddr)

	port, err := net.ListenUDP("udp4", listenAddr)

	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("udp4", bcast)

	if err != nil {
		panic(err)
	}

	ips := make(map[string]int)

	go func() {
		for {

			time.Sleep(time.Second)

			msg := internal.SendMsg(ip, 8000)
			conn.Write(msg.Bytes())
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			if len(ips) > 0 {
				for k, v := range ips {
					fmt.Println(k, v)
				}
			}
		}
	}()

	for {
		buf := make([]byte, 59)

		_, _, err = port.ReadFrom(buf)

		if err != nil {
			panic(err)
		}

		newBuf := bytes.NewBuffer(buf)

		msg, err := internal.ReceiveMsg(newBuf)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if ip.Equal(msg.Ip) {
			continue
		}

		if _, ok := ips[net.IP(msg.Ip).String()]; !ok {
			ips[net.IP(msg.Ip).String()] = 1
			fmt.Println(time.Now(), "Found", net.IP(msg.Ip).String(), msg.Port)
		} else {
			ips[net.IP(msg.Ip).String()] += 1
		}
	}

}
