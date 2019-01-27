package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	"errors"
	"strings"
)

var ConnectionTimeout = errors.New("dial tcp 80.211.15.40:2: i/o timeout")

func ScanPort(ip string, port string, timeout time.Duration) {
	target := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "refused") {
			fmt.Println(port, "filtered") 
		} else if strings.Contains(err.Error(), "timeout"){
			fmt.Println(port, "closed")
		} else {
			fmt.Println(port, "error")
			ScanPort (ip, port, 2*time.Second)
		}
		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

func main() {

	if len(os.Args) == 2 {
		fmt.Println(os.Args)
		ip := os.Args[1]
		fmt.Println(ip)
		for port := 0; port <= 65535; port++ {
			port := strconv.Itoa(port)
			go ScanPort (ip, port, 2*time.Second)
		}		
	} else {
		fmt.Println("Usage: scan ip")
	}

}