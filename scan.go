package main

import (
	"fmt"
	"net"
	"sync"
)

func scanTCP(ip string, port int, wg *sync.WaitGroup) bool {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("%s:%d is not open(TCP)\n", ip, port)
		return false
	}
	conn.Close()
	fmt.Printf("%s:%d is open(TCP)\n", ip, port)
	return true
}

func scanUDP(ip string, port int, wg *sync.WaitGroup) bool {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Printf("%s:%d is not open(UDP)\n", ip, port)
		return false
	}
	conn.Close()
	fmt.Printf("%s:%d is open(UDP)\n", ip, port)
	return true
}

func main() {
	// ip := "101.34.38.11"
	ip := "fadinglight.cn"
	ports := []int{80, 81, 82, 83}

	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)
		go scanUDP(ip, port, &wg)
	}
	wg.Wait()
	fmt.Println("done.")
}
