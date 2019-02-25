package serverlib

import (
	"log"
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
	BindIP            string `json:"rpc-bind-ip"`
	BindPort          int    `json:"rpc-bind-port"`
	ClientID          string `json:"client-id"`
	ClientSecret      string
	AuthenticationURL string `json:"authentication"`
	RedirectURL       string `json:"redirect-url"`
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
