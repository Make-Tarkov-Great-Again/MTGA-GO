// Package certificate acquires a self-signed certificate
package services

import (
	"MT-GO/tools"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

const certPath string = "user/cert"

// Certificate represents a certificate in the system certificate authority format
type Certificate struct {
	CertFile    string
	KeyFile     string
	Certificate tls.Certificate
}

// GetCertificate returns a Certificate for HTTPS server
func GetCertificate(ip string) *Certificate {
	cert := Certificate{
		CertFile: filepath.Join(certPath, "cert.pem"),
		KeyFile:  filepath.Join(certPath, "key.pem"),
	}

	if tools.FileExist(cert.CertFile) && tools.FileExist(cert.KeyFile) {
		return &cert
	}

	if !tools.FileExist(certPath) {
		err := os.Mkdir(certPath, 0700)
		if err != nil {
			panic(err)
		}
	}

	cert.setCertificate(ip)
	return &cert
}

// Generate SHA256 certificate for HTTPS server
func (cg *Certificate) setCertificate(ip string) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	notBefore := time.Now().UTC()
	notAfter := notBefore.AddDate(1, 0, 0)

	maxSerialNumber := new(big.Int).Lsh(big.NewInt(1), 128) // 1 << 128 = 2^128
	serialNumber, err := rand.Int(rand.Reader, maxSerialNumber)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   "MTGA",
			Organization: []string{"Make Tarkov Great Again"},
		},
		IPAddresses: []net.IP{net.ParseIP(ip)},
		DNSNames:    []string{"localhost"},
		NotBefore:   notBefore,
		NotAfter:    notAfter,
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		ExtraExtensions: []pkix.Extension{{
			Id:    asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 1}, // subjectAltName extension OID
			Value: []byte{0x05, 0x00},
		}},
	}

	// Self-sign the certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		panic(err)
	}

	// Save the certificate to files
	certOut, err := os.Create(cg.CertFile)
	if err != nil {
		panic(err)
	}
	defer certOut.Close()

	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		panic(err)
	}

	keyOut, err := os.Create(cg.KeyFile)
	if err != nil {
		panic(err)
	}
	defer keyOut.Close()
	privBytes := x509.MarshalPKCS1PrivateKey(priv)

	err = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes})
	if err != nil {
		panic(err)
	}

	fmt.Println("Certificate and key generated successfully")
}
