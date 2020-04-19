package main

import (
	"crypto/tls"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()
	keyPath := filepath.Join(dir,"data/key.pem")
	certPath := filepath.Join(dir,"data/certificate.cer")

	tlsConfig := tls.Config{}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		panic(err)
	}
	tlsConfig.Certificates = []tls.Certificate{cert}
}
