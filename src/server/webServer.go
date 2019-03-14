package server

import (
	"../serverlib"
	"fmt"
	"net/http"
)

func RunWebServer() {

	serverlib.DebugLog.Println("Started running on http://localhost:6352")
	http.ListenAndServeTLS(fmt.Sprintf("%s:%d"))
}
