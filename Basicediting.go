package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	
	var ip_info string
	var subnet_info string
	fi_ip, err_ip := os.Open("/tmp/ip")
	if err_ip != nil {
		fmt.Printf("Error: %s\n", err_ip)
		return
	}
	defer fi_ip.Close()

	br_ip := bufio.NewReader(fi_ip)
	for {
		ip, _, c_ip := br_ip.ReadLine()
		if c_ip == io.EOF {
			break
		}
		ip_info = string(ip)
	}

	fi_subnet, err_subnet := os.Open("/tmp/route")
	if err_subnet != nil {
		fmt.Printf("Error: %s\n", err_subnet)
		return
	}
	defer fi_subnet.Close()

	br_subnet := bufio.NewReader(fi_subnet)
	for {
		subnet, _, c_subnet := br_subnet.ReadLine()
		if c_subnet == io.EOF {
			break
		}
		subnet_info = string(subnet)
	}

	_,ipnetA,_ := net.ParseCIDR(subnet_info)
	ipB,_,_ := net.ParseCIDR(ip_info)

	if ipnetA.Contains(ipB) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}