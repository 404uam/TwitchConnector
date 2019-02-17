package src

import (
	"./serverlib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/rpc"
	"os"
	"path/filepath"
)

/***********************EXPORT METHODS******************/
type Twitch string

func (s *Twitch) Register(args *serverlib.ClientCred, reply *bool) error {

	return nil
}
func (s *Twitch) GetToken(args **serverlib.ClientCred, reply *bool) error {

	return nil
}

/*********************End of exported methods***************/

var (
	config serverlib.Config
)

func loadSettings(path string) {
	file, err := os.Open(path)
	serverlib.IsErr("Config not read", err)

	buffer, err := ioutil.ReadAll(file)
	serverlib.IsErr("Error Reading", err)

	err = json.Unmarshal(buffer, &config)
	serverlib.IsErr("Error unmarshalling json", err)
}

func main() {
	absPath, _ := filepath.Abs("./serverlib/settings.json")
	loadSettings(absPath)
	serverlib.DebugLog.Println(fmt.Sprintf("%s:%d", config.BindIP, config.BindPort))

	twitch := new(Twitch)

	server := rpc.NewServer()
	err := server.Register(twitch)
	serverlib.IsErr("Failed to register RPC server", err)

	l, e := net.Listen("tcp", config.BindIP)

	serverlib.IsErr("Could not bind to to listen", e)
	serverlib.DebugLog.Printf("serverlib started. Receiving on %s\n", fmt.Sprintf("%s:%d", config.BindIP, config.BindPort))

	for {
		conn, _ := l.Accept()
		go server.ServeConn(conn)
	}

}
