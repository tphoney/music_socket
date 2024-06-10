package main

// simple web server

import (
	"fmt"
	"net"
	"net/http"
	"tphoney/music_socket/web"
)

var (
	TargetServer = "ws://192.168.1.150:9090"
)

func main() {
	webby := web.Config{
		TargetServer: TargetServer,
	}

	ip := GetOutboundIP()
	fmt.Printf("Starting server on http://%s:9191", ip.String())
	mux := http.NewServeMux()

	mux.HandleFunc("/", webby.PageHandler)
	mux.HandleFunc("/send", webby.SendMessageHandler)

	err := http.ListenAndServe(":9191", mux)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Failed to get local IP address")
		return nil
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
