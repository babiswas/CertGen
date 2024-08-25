package main

import (
	"CAService/CAUtil"
	"CAService/CertGenUtil"
	"CAService/ClientUtil"
	"fmt"
)

func main() {
	// Create CA's private key and self-signed certificate
	caPrivateKey, caCert, err := CAUtil.CreateCACertificate()
	if err != nil {
		panic(err)
	}

	// Generate client's private key and CSR
	clientPrivateKey, clientCSRDER, err := ClientUtil.CreateClientCSR()
	if err != nil {
		panic(err)
	}

	// CA signs the client's CSR to issue a certificate
	err = CertGenUtil.SignCSR(caPrivateKey, caCert, clientCSRDER)
	if err != nil {
		panic(err)
	}

	fmt.Println(clientPrivateKey)

	println("CA and client certificates have been generated successfully.")
}
