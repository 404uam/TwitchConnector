package server

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

/***********************EXPORT METHODS******************/
type Twitch string

func (s *Twitch) Register(args *ClientCred, reply *string) error {
	DebugLog.Println("Hi i've been called to register")

	provider, err := oidc.NewProvider(context.Background(), ServerConfig.AuthenticationURL)
	IsErr("", err)

	DebugLog.Println("Creating oauth2Config...")

	oauth2Config := oauth2.Config{
		ClientID:     ServerConfig.ClientID,
		ClientSecret: ServerConfig.ClientSecret,
		RedirectURL:  ServerConfig.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       scopes,
	}

	oidcVerifier = provider.Verifier(&oidc.Config{ClientID: ServerConfig.ClientID})

	var tokenBytes [255]byte
	if _, err := rand.Read(tokenBytes[:]); err != nil {
		return AnnotateError(err, "Couldn't generate a session!", http.StatusInternalServerError)
	}

	state := hex.EncodeToString(tokenBytes[:])
	session.AddFlash(state, stateCallbackKey)
	DebugLog.Println(state)

	*reply = oauth2Config.AuthCodeURL(state)
	return nil
}
func (s *Twitch) GetToken(args *ClientCred, reply *bool) error {

	return nil
}

/*********************End of exported methods***************/
const (
	stateCallbackKey = "oauth-state-callback"
	oauthSessionName = "oauth-oidc-session"
)

var (
	scopes       = []string{oidc.ScopeOpenID, "user_subscriptions"}
	oidcVerifier *oidc.IDTokenVerifier
	verify       string
	cookieSecret = []byte(os.Getenv("SessionKey"))
	cookieStore  = sessions.NewCookieStore(cookieSecret)
	session      = sessions.NewSession(cookieStore, oauthSessionName)
)

func StartRPCServer() {
	DebugLog.Println(fmt.Sprintf("%s:%d", ServerConfig.BindRPCIP, ServerConfig.BindRPCPort))
	DebugLog.Println(ServerConfig.ClientID)

	twitch := new(Twitch)

	server := rpc.NewServer()
	err := server.Register(twitch)
	IsErr("Failed to register RPC server", err)

	l, e := net.Listen("tcp", fmt.Sprintf("%s:%d", ServerConfig.BindRPCIP, ServerConfig.BindRPCPort))
	IsErr("Could not bind to listen", e)

	DebugLog.Printf("Server started. Receiving on %s\n", fmt.Sprintf("%s:%d", ServerConfig.BindRPCIP, ServerConfig.BindRPCPort))
	for {
		conn, _ := l.Accept()
		go server.ServeConn(conn)
	}
}
