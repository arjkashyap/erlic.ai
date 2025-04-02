package utils

import (
	"crypto/x509"
	"fmt"
	"os"
)

func LoadCACert(caCertPath string) (*x509.CertPool, error) {
	fmt.Println("Importing CA Cert")
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %w", err)
	}

	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(caCert)
	if !ok {
		return nil, fmt.Errorf("failed to append CA certificate")
	}

	return caCertPool, nil
}
