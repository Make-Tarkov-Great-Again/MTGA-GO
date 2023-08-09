// Package certificate acquires a self-signed certificate
package certificate

import (
	"MT-GO/tools"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
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
	CertFile string
	KeyFile  string
}

// GetCertificate returns a Certificate for HTTPS server
func GetCertificate(ip string, hostname string) *Certificate {
	cert := Certificate{
		CertFile: filepath.Join(certPath, "cert.pem"),
		KeyFile:  filepath.Join(certPath, "key.pem"),
	}

	if tools.FileExist(cert.CertFile) && tools.FileExist(cert.KeyFile) {
		return &cert
	}

	if !tools.FileExist(certPath) {
		os.Mkdir(certPath, 0700)
	}

	cert.generateCertificate(ip, hostname)
	return &cert
}

// Generate SHA256 certificate for HTTPS server
func (cg *Certificate) generateCertificate(ip string, hostname string) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
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
		IPAddresses:           []net.IP{net.ParseIP(ip)},
		DNSNames:              []string{hostname},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	cert, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		panic(err)
	}

	certFile, err := os.Create(cg.CertFile)
	if err != nil {
		panic(err)
	}
	defer certFile.Close()

	pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: cert})

	keyFile, err := os.Create(cg.KeyFile)
	if err != nil {
		panic(err)
	}
	defer keyFile.Close()

	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		panic(err)
	}

	pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})
	fmt.Println("Certificate and key generated successfully")
}
