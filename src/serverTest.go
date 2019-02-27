package main

import (
	"./serverlib"
	"fmt"
	"net/rpc"
)

func main() {
	serverIP := "127.0.0.1"
	serverPort := "6969"
	var bo string

	c, err := rpc.Dial("tcp", fmt.Sprintf("%s:%s", serverIP, serverPort))
	serverlib.IsErr("Cannot Dial", err)
	defer c.Close()

	err = c.Call("Twitch.Register", serverlib.ClientCred{"billy"}, &bo)
	serverlib.IsErr("", err)

	fmt.Println(bo)

}
