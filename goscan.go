//  commment in the begining?
package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "time"
    "strings"
    "github.com/korovkin/limiter"
)

var errors int = 0
const numberOfPorts int = 65535

func ScanPort(ip string, port string, timeout time.Duration) {
    target := fmt.Sprintf("%s:%s", ip, port)
    conn, err := net.DialTimeout("tcp", target, timeout)
    if err != nil {
        if strings.Contains(err.Error(), "refused") == false && strings.Contains(err.Error(), "timeout") == false {
            errors++
            if errors >= 1000 {
                fmt.Println("Too many errors!!! Consider decreasing number of threads!")
                os.Exit(1)
            }
            ScanPort (ip, port, timeout)
        }
        return
    }
    conn.Close()
    fmt.Println(port, "open")
}

func main() {
    if len(os.Args) == 4 {
        ip := os.Args[1]
        threads, _ := strconv.Atoi(os.Args[2])
        numOfMilliseconds, _ := strconv.Atoi(os.Args[3])   
        timeout := time.Duration(numOfMilliseconds)*time.Millisecond

		limit := limiter.NewConcurrencyLimiter(threads)

		for i := 1; i < 65535; i++ {
			port := strconv.Itoa(i)
			limit.Execute(func() {
					ScanPort(ip, port, timeout)
  				})
  		}
  		limit.Wait()

    } else {
        fmt.Println("Usage: ip threads timeout(milliseconds)")
    }
}

