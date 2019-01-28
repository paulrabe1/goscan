package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	// "errors"
	"strings"
)

func ScanPort(ip string, port string, timeout time.Duration) {
	target := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "refused") {
	// 		fmt.Println(port, "filtered") 
	// 	} else if strings.Contains(err.Error(), "timeout"){
	// 		fmt.Println(port, "closed")
	// 	} else {
	// 		// fmt.Println(port, "error")
	// 		ScanPort (ip, port, 500*time.Millisecond)
	// 	}
	// 	return
	// }

	if err != nil {
		if strings.Contains(err.Error(), "refused") == false && strings.Contains(err.Error(), "timeout") == false {
			ScanPort (ip, port, 500*time.Millisecond)
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
			go ScanPort (ip, port, 500*time.Millisecond)
		}		
	} else {
		fmt.Println("Usage: scan ip")
	}

}