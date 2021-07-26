package server

import (
	"testing"
)

func TestServer_Run(t *testing.T) {
	address := "127.0.0.1:8090"
	Start(address)
	select {}
}
