package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//TODO Get streamer tokens && Store token in a safe place
//TODO Refresh streamer token periodically

type ServerCred struct {
	client_id     string
	client_secret string
	grant_type    string
	refresh_token string
}

type ClientCred struct {
	Username string
}

type Config struct {
	BindRPCIP         string `json:"rpc-bind-ip"`
	BindRPCPort       int    `json:"rpc-bind-port"`
	BindWebIP         string `json:"web-bind-ip"`
	BindWebPort       int    `json:"web-bind-port"`
	SSLCert           string `json:"path-to-ssl-cert"`
	SSLKey            string `json:"path-to-ssl-key"`
	ClientID          string `json:"client-id"`
	ClientSecret      string
	AuthenticationURL string `json:"authentication"`
	RedirectURL       string `json:"redirect-url"`
}

var (
	DebugLog     = log.New(os.Stderr, "[Server] ", 0)
	ErrLog       = log.New(os.Stderr, "[Error] ", 0)
	ServerConfig Config
)

func IsErr(msg string, e error) {
	if e != nil {
		ErrLog.Fatalf("%s, err = %s\n", msg, e.Error())
	}
}

func LoadSettings(path string) {
	file, err := os.Open(path)
	IsErr("Config not read", err)

	buffer, err := ioutil.ReadAll(file)
	IsErr("Error Reading", err)

	err = json.Unmarshal(buffer, &ServerConfig)
	ServerConfig.ClientSecret = os.Getenv("ClientSecret")
	IsErr("Error unmarshalling json", err)
}

/********* Added from twitch oauth authorization code sample *****/
/* Can be accessed here https://github.com/twitchdev/authentication-samples */

// HumanReadableError represents error information
// that can be fed back to a human user.
//
// This prevents internal state that might be sensitive
// being leaked to the outside world.
type HumanReadableError interface {
	HumanError() string
	HTTPCode() int
}

// HumanReadableWrapper implements HumanReadableError
type HumanReadableWrapper struct {
	ToHuman string
	Code    int
	error
}

func (h HumanReadableWrapper) HumanError() string { return h.ToHuman }
func (h HumanReadableWrapper) HTTPCode() int      { return h.Code }

// AnnotateError wraps an error with a message that is intended for a human end-user to read,
// plus an associated HTTP error code.
func AnnotateError(err error, annotation string, code int) error {
	if err == nil {
		return nil
	}
	return HumanReadableWrapper{ToHuman: annotation, error: err}
}

type Handler func(http.ResponseWriter, *http.Request) error
