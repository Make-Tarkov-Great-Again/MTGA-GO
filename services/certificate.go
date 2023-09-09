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
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const certPath string = "user/cert"

// Certificate represents a certificate in the system certificate authority format
type Certificate struct {
	CertFile    string
	KeyFile     string
	Certificate tls.Certificate
}

const certSubject string = "O=Make Tarkov Great Again, CN=MTGA Root CA Certificate"

// GetCertificate returns a Certificate for HTTPS server
func GetCertificate(ip string) *Certificate {
	cert := Certificate{
		CertFile: filepath.Join(certPath, "cert.pem"),
		KeyFile:  filepath.Join(certPath, "key.pem"),
	}

	if cert.verifyCertificate() {
		return &cert
	} else {

		if !tools.FileExist(certPath) {
			err := os.Mkdir(certPath, 0700)
			if err != nil {
				panic(err)
			}
		}

		cert.setCertificate(ip)
		cert.installCertificate()
		return &cert
	}
}

// Generate SHA256 certificate for HTTPS server
func (cg *Certificate) setCertificate(ip string) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	notBefore := time.Now().UTC()
	notAfter := notBefore.AddDate(0, 0, 2)

	maxSerialNumber := new(big.Int).Lsh(big.NewInt(1), 128) // 1 << 128 = 2^128
	serialNumber, err := rand.Int(rand.Reader, maxSerialNumber)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   "MTGA Root CA Certificate",
			Organization: []string{"Make Tarkov Great Again"},
		},
		IPAddresses: []net.IP{net.ParseIP(ip)},
		DNSNames:    []string{"localhost"},
		NotBefore:   notBefore,
		NotAfter:    notAfter,
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		ExtraExtensions: []pkix.Extension{
			{
				Id:    asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 1}, // subjectAltName extension OID
				Value: []byte{0x05, 0x00},
			},
		},
		IsCA:                  true,
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		panic(err)
	}

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

var certFileExist bool
var keyFileExist bool

// verifyCertificate verifies the certificate to see if it is still valid
func (cg *Certificate) verifyCertificate() bool {
	certFileExist = tools.FileExist(cg.CertFile)
	keyFileExist = tools.FileExist(cg.KeyFile)

	if !certFileExist || !keyFileExist {
		cg.removeCertificate()
		return false
	}

	if cg.isCertificateInstalled() && cg.isCertificateExpired() {
		fmt.Println("Certificate is valid.")
		return true
	} else {

		cg.removeCertificate()
		return false
	}
}

const deleteCertificatePrompt string = "Certificate is expired and needs to be renewed, you will be prompted to delete the certificate. Type `Yes` if you understand, and would like to proceed."

func (cg *Certificate) removeCertificate() {

	var input string
	fmt.Println(deleteCertificatePrompt)
	for {
		fmt.Printf("> ")
		fmt.Scanln(&input)
		if strings.Contains(strings.ToLower(input), "yes") {
			if certFileExist {
				err := os.Remove(cg.CertFile)
				if err != nil {
					fmt.Println("Failed to remove the certificate")
					panic(err)
				}
			}

			if keyFileExist {
				err := os.Remove(cg.KeyFile)
				if err != nil {
					fmt.Println("Failed to remove the certificate")
					panic(err)
				}
			}

			if cg.isCertificateInstalled() {
				cmd := exec.Command("certutil", "-delstore", "-user", "Root", "*MTGA*", "-f")
				output, err := cmd.CombinedOutput()
				if err != nil {
					exitErr, _ := err.(*exec.ExitError)
					if exitErr.ProcessState.ExitCode() == exitCode {
						fmt.Println("User cancelled the deletion of the certificate")
						os.Exit(0)
					}
					fmt.Println(output)
					panic(err)
				}
				fmt.Println("Certificate removed from System")
			} else {
				return
			}
		} else {
			fmt.Println("User doesn't want to delete the expired certificate, disconnecting...")
			os.Exit(0)
		}
	}

}

const installCertificatePrompt string = "In order for Notifications/WebSocket to work in-game, we need to install the SHA256 certificate to your Trusted Root Certification Authority under `MTGA Root CA Certificate`. \n\nTLDR: Type `yes` if you want to play"
const exitCode int = 2147943623

func (cg *Certificate) installCertificate() {

	fmt.Println(installCertificatePrompt)
	var input string

	for {
		fmt.Printf("> ")
		fmt.Scanln(&input)

		if strings.Contains(strings.ToLower(input), "yes") {
			_, err := exec.Command("certutil", "-addstore", "-user", "Root", cg.CertFile).CombinedOutput()
			if err != nil {
				exitErr, _ := err.(*exec.ExitError)
				if exitErr.ProcessState.ExitCode() == exitCode {
					fmt.Println("User cancelled the installation")
					os.Exit(0)
				}
				fmt.Println("Failed to install the certificate.")
				panic(err)
			}
			fmt.Println("Certificate installed.")
			return
		} else {
			fmt.Println("User doesn't want to install the certificate, disconnecting...")
			os.Exit(0)
		}
	}

}

func (cg *Certificate) isCertificateExpired() bool {
	cmd := exec.Command("certutil", "-verifystore", "-user", "Root", "*MTGA*")

	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "Object was not found") {
			fmt.Println("Certificate is not installed.")
			return false
		}
		fmt.Println("Failed to verify if the certificate has expired.")
		return false
	}

	return !strings.Contains(string(output), "This certificate is OK.")
}

func (cg *Certificate) isCertificateInstalled() bool {
	cmd := exec.Command("certutil", "-store", "-user", "Root", "*MTGA*")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "Object was not found") {
			fmt.Println("Certificate is not installed.")
			return false
		}
		fmt.Println("Failed to verify if the certificate is installed.")
		return false
	}

	return strings.Contains(string(output), certSubject)
}
