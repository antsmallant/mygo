package main

import(
	"fmt"
	"net"
	"time"
)

func tcpGather(ip string, ports []string) map[string]string {
	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results[port] = "failed"
			// todo log handler
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}
	return results
}

func test(host string, ports []string) {
	var r = tcpGather(host, ports)
	fmt.Println("test results", host, ports)
	fmt.Println(r)
}

func main() {
	test("QGameserverPRD.lkgame.com", []string{"8888"})
	test("42.194.157.90", []string{"8888"})
}
