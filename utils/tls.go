package utils

import (
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func GetTLSConfig() *tls.Config {
	manager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("certs"),
		HostPolicy: autocert.HostWhitelist(), // optionally enforce subdomain
	}
	go http.ListenAndServe(":80", manager.HTTPHandler(nil))
	return manager.TLSConfig()
}
