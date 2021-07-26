package client

import (
	"fmt"
	"testing"
)

func TestInvoke(t *testing.T) {
	resp, err := Invoke(1, 2, "test", "test...")
	if err != nil {
		fmt.Println("invoke err", err.Error())
		return
	}
	fmt.Printf("resp:%+v", resp)
}
