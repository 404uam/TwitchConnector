package server

import (
	"../serverlib"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
)

/***********************EXPORT METHODS******************/
type Twitch string

func (s *Twitch) Register(args *serverlib.ClientCred, reply *string) error {
	serverlib.DebugLog.Println("Hi i've been called to register")

	provider, err := oidc.NewProvider(context.Background(), config.AuthenticationURL)
	serverlib.IsErr("", err)

	serverlib.DebugLog.Println("Creating oauth2Config...")

	oauth2Config := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       scopes,
	}

	oidcVerifier = provider.Verifier(&oidc.Config{ClientID: config.ClientID})

	var tokenBytes [255]byte
	if _, err := rand.Read(tokenBytes[:]); err != nil {
		return serverlib.AnnotateError(err, "Couldn't generate a session!", http.StatusInternalServerError)
	}

	state := hex.EncodeToString(tokenBytes[:])
	session.AddFlash(state, stateCallbackKey)
	serverlib.DebugLog.Println(state)

	*reply = oauth2Config.AuthCodeURL(state)
	return nil
}
func (s *Twitch) GetToken(args *serverlib.ClientCred, reply *bool) error {

	return nil
}

/*********************End of exported methods***************/
const (
	stateCallbackKey = "oauth-state-callback"
	oauthSessionName = "oauth-oidc-session"
)

var (
	config       serverlib.Config
	scopes       = []string{oidc.ScopeOpenID, "user_subscriptions"}
	oidcVerifier *oidc.IDTokenVerifier
	verify       string
	cookieSecret = []byte(os.Getenv("SessionKey"))
	cookieStore  = sessions.NewCookieStore(cookieSecret)
	session      = sessions.NewSession(cookieStore, oauthSessionName)
)

func loadSettings(path string) {
	file, err := os.Open(path)
	serverlib.IsErr("Config not read", err)

	buffer, err := ioutil.ReadAll(file)
	serverlib.IsErr("Error Reading", err)

	err = json.Unmarshal(buffer, &config)
	config.ClientSecret = os.Getenv("ClientSecret")
	serverlib.IsErr("Error unmarshalling json", err)
}

func StartRPCServer() {
	absPath, _ := filepath.Abs("../src/serverlib/settings.json")
	loadSettings(absPath)
	serverlib.DebugLog.Println(fmt.Sprintf("%s:%d", config.BindRPCIP, config.BindRPCPort))
	serverlib.DebugLog.Println(config.ClientID)

	twitch := new(Twitch)

	server := rpc.NewServer()
	err := server.Register(twitch)
	serverlib.IsErr("Failed to register RPC server", err)

	l, e := net.Listen("tcp", fmt.Sprintf("%s:%d", config.BindRPCIP, config.BindRPCPort))
	serverlib.IsErr("Could not bind to listen", e)

	serverlib.DebugLog.Printf("Server started. Receiving on %s\n", fmt.Sprintf("%s:%d", config.BindRPCIP, config.BindRPCPort))
	for {
		conn, _ := l.Accept()
		go server.ServeConn(conn)
	}
}
