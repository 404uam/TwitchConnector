package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"os"
	"path/filepath"
)

/***********************EXPORT METHODS******************/
type Twitch string

func (s *Twitch) Register(args *dfslib.Args, reply *bool) error {

	return nil
}
func (s *Twitch) GetToken(args *dfslib.Args, reply *bool) error {

	return nil
}

/*********************End of exported methods***************/
type Config struct {
	BindIP   string `json:"rpc-bind-ip"`
	BindPort int    `json:"rpc-bind-port"`
}

var (
	config   Config
	debugLog *log.Logger = log.New(os.Stderr, "[Server] ", 0)
	errLog   *log.Logger = log.New(os.Stderr, "[Error] ", 0)
)

func loadSettings(path string) {
	file, err := os.Open(path)
	isErr("Config not read", err)

	buffer, err := ioutil.ReadAll(file)
	isErr("Error Reading", err)

	err = json.Unmarshal(buffer, &config)
	isErr("Error unmarshalling json", err)
}

func main() {
	absPath, _ := filepath.Abs("./src/Server/settings.json")
	loadSettings(absPath)
	debugLog.Println(fmt.Sprintf("%s:%d", config.BindIP, config.BindPort))

	twitch := new(Twitch)

	server := rpc.NewServer()
	err := server.Register(twitch)
	isErr("Failed to register RPC server", err)

	l, e := net.Listen("tcp", config.BindIP)

	isErr("Could not bind to to listen", e)
	debugLog.Printf("Server started. Receiving on %s\n", fmt.Sprintf("%s:%d", config.BindIP, config.BindPort))

	for {
		conn, _ := l.Accept()
		go server.ServeConn(conn)
	}

}
