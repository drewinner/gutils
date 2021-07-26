package server

import (
	"testing"
)

func TestServer_Run(t *testing.T) {
	Start("127.0.0.1:8090")
	select {}
}
