package main

import (
	"log"
	"net/http"
	"ghostproxy/proxy"
	"ghostproxy/dashboard"
	"ghostproxy/utils"
)

func main() {
	utils.LoadEnv()

	go dashboard.StartDashboard() // Serves dashboard on :8081

	proxyHandler := http.HandlerFunc(proxy.ProxyHandler)
	server := &http.Server{
		Addr:      ":443",
		Handler:   proxyHandler,
		TLSConfig: utils.GetTLSConfig(),
	}

	log.Println("[+] Starting ghostproxy on port 443")
	log.Fatal(server.ListenAndServeTLS("", ""))
}
