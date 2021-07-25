package client

import "testing"

func TestInvoke(t *testing.T) {
	Invoke(1, 2, "test", "{}")
}
