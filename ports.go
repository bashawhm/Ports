package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	ip := os.Args[1]
	maxPort := 1000
	if len(os.Args) == 3 {
		maxPort, _ = strconv.Atoi(os.Args[2])
	}

	maxPort++
	var wg sync.WaitGroup
	foundCount := 0

	fmt.Println("Starting Report...\n-------------------------")

	for i := 1; i < maxPort; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			ipAddr := ip
			ipAddr += ":"
			ipAddr += strconv.Itoa(port)

			conn, _ := net.Dial("tcp", ipAddr)
			if conn != nil {
				fmt.Printf("tcp\t%d\topen\n", port)
				foundCount++
				return
			}
			return
		}(i)
	}
	wg.Wait()

	if foundCount == 0 {
		fmt.Println("No open connection in range")
	}

}
