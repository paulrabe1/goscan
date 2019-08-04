package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "time"
    "strings"
)

const numberOfPorts int = 65535

func scanport(ip string, port string, timeout time.Duration) {
    target := fmt.Sprintf("%s:%s", ip, port)
    conn, err := net.DialTimeout("tcp", target, timeout)
    if err != nil {
        if strings.Contains(err.Error(), "refused") == false && strings.Contains(err.Error(), "timeout") == false {
            scanport (ip, port, 500*time.Millisecond)
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
    ip := os.Args[1]
    jobs := make(chan int, numberOfPorts)
    results := make(chan int, numberOfPorts)

    for w := 1; w <= 100; w++ {
        go worker(w, jobs, results, ip, 500*time.Millisecond)
    }

    for j := 1; j <= numberOfPorts; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numberOfPorts; a++ {
        <-results
    }
}
