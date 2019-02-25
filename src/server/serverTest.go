package server

import (
	"../serverlib"
	"fmt"
	"net/rpc"
)

func main() {
	serverIP := "127.0.0.1"
	serverPort := "6969"
	bo := false

	c, err := rpc.Dial("tcp", fmt.Sprintf("%s:%s", serverIP, serverPort))
	serverlib.IsErr("Cannot Dial", err)
	defer c.Close()

	err = c.Call("Twitch.Register", serverlib.ClientCred{"billy"}, &bo)
	serverlib.IsErr("", err)

	if bo {
		fmt.Println("Hi Ho i'm True")
	} else {
		fmt.Println("Ho Hi i'm False")
	}

}
