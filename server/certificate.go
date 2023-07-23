package server

import (
	"MT-GO/tools"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"
)

type TLSCertKeyPair struct {
	Cert []byte
	Key  []byte
}

func generateTLS(serverIP string, serverHostname string, certificateDays int) (*TLSCertKeyPair, error) {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	// Create a certificate template
	now := time.Now()
	certTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(now.Unix()),
		Subject:               pkix.Name{CommonName: "MGTA Server", Organization: []string{"MTGA"}},
		NotBefore:             now.UTC(),
		NotAfter:              now.AddDate(0, 0, certificateDays).UTC(),
		BasicConstraintsValid: true,
		DNSNames:              []string{serverHostname},
		IPAddresses:           []net.IP{net.ParseIP(serverIP)},
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		SignatureAlgorithm:    x509.SHA256WithRSA, // specify SHA256 signature algorithm
		IsCA:                  true,
	}

	// Generate the certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, err
	}

	// Encode the certificate and private key in PEM format
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})

	// Store the certificate and private key together in a struct
	certString := string(certPEM)
	keyString := string(keyPEM)

	if err := tools.WriteToFile("./user/cert/cert.pem", certString); err != nil {
		return nil, err
	}

	if err := tools.WriteToFile("./user/cert/key.pem", keyString); err != nil {
		return nil, err
	}

	return &TLSCertKeyPair{
		Cert: certPEM,
		Key:  keyPEM,
	}, nil
}
