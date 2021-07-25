package main

import (
	_ "github.com/drewinner/gutils/cnode"
	"github.com/drewinner/gutils/rpc/client"
)

func main() {
	client.Invoke(1, 2, "test", "")
}
