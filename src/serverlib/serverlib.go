package serverlib

import (
	"log"
	"os"
)

type serverCred struct {
	client_id     string
	client_secret string
	grant_type    string
	refresh_token string
}

type clientCred struct {
	username string
}

type Config struct {
	BindIP   string `json:"rpc-bind-ip"`
	BindPort int    `json:"rpc-bind-port"`
}

var (
	DebugLog = log.New(os.Stderr, "[Server] ", 0)
	ErrLog   = log.New(os.Stderr, "[Error] ", 0)
)

func IsErr(msg string, e error) {
	if e != nil {
		ErrLog.Fatalf("%s, err = %s\n", msg, e.Error())
	}
}
