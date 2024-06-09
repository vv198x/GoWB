package pgsql

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/vv198x/GoWB/config"
	"io/ioutil"
	"log/slog"
	"path/filepath"
)

func TlsConn() *tls.Config {
	certPath := filepath.Join(config.Get().LogDir, "server.crt")
	keyPath := filepath.Join(config.Get().LogDir, "server.key")

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		slog.Error("failed to load client certificate: %v", err)
	}

	CAFilePath := filepath.Join(config.Get().LogDir, "root.crt")

	CACert, err := ioutil.ReadFile(CAFilePath)
	if err != nil {
		slog.Error("failed to load server certificate: %v", err)
	}

	CACertPool := x509.NewCertPool()
	CACertPool.AppendCertsFromPEM(CACert)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            CACertPool,
		InsecureSkipVerify: true,
	}
	return tlsConfig
}
