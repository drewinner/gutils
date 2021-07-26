package client

import (
	"fmt"
	"testing"
)

func TestInvoke(t *testing.T) {
	resp, err := Invoke("127.0.0.1:8090",1, 2, "test", "test...")
	if err != nil {
		fmt.Println("invoke err", err.Error())
		return
	}
	fmt.Printf("resp:%+v", resp)
}
