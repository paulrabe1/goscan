//  commment in the begining?
package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "time"
    "strings"
)

var errors int = 0
const numberOfPorts int = 65535

func scanport(ip string, port string, timeout time.Duration) {
    target := fmt.Sprintf("%s:%s", ip, port)
    conn, err := net.DialTimeout("tcp", target, timeout)
    if err != nil {
        if strings.Contains(err.Error(), "refused") == false && strings.Contains(err.Error(), "timeout") == false {
            errors++
            if errors >= 1000 {
                fmt.Println("Too many errors!!! Consider decreasing number of threads!")
                os.Exit(1)
            }
            scanport (ip, port, timeout)
        }
        return
    }
    conn.Close()
    fmt.Println(port, "open")
}

func worker(id int, jobs <-chan int, results chan<- int, ip string, timeout time.Duration) {
    for j := range jobs {
        port := strconv.Itoa(j)
        scanport(ip, port, timeout)
        results <- j * 2
    }
}

func main() {
    if len(os.Args) == 4 {
        ip := os.Args[1]
        threads, _ := strconv.Atoi(os.Args[2])
        numOfMilliseconds, _ := strconv.Atoi(os.Args[3])
        timeout := time.Duration(numOfMilliseconds)*time.Millisecond

        jobs := make(chan int, numberOfPorts)
        results := make(chan int, numberOfPorts)

        for w := 1; w <= threads; w++ {
            go worker(w, jobs, results, ip, timeout)
        }

        for j := 1; j <= numberOfPorts; j++ {
            jobs <- j
        }
        close(jobs)

        for a := 1; a <= numberOfPorts; a++ {
            <-results
        }   
    } else {
        fmt.Println("Usage: ip threads timeout(milliseconds)")
    }
}
