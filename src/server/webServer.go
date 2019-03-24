package server

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func RunWebServer() {
	/* Imported from Twitch sample code
	Provides useful and logical handling of HTTP errors and handleFunc function to host endpoints
	*/
	/*	var middleware = func(h Handler) Handler {
			return func(w http.ResponseWriter, r *http.Request) (err error) {
				// parse POST body, limit request size
				if err = r.ParseForm(); err != nil {
					return AnnotateError(err, "Something went wrong! Please try again.", http.StatusBadRequest)
				}

				return h(w, r)
			}
		}

		// errorHandling is a middleware that centralises error handling.
		// this prevents a lot of duplication and prevents issues where a missing
		// return causes an error to be printed, but functionality to otherwise continue
		// see https://blog.golang.org/error-handling-and-go
		var errorHandling = func(handler func(w http.ResponseWriter, r *http.Request) error) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := handler(w, r); err != nil {
					var errorString string = "Something went wrong! Please try again."
					var errorCode int = 500

					if v, ok := err.(HumanReadableError); ok {
						errorString, errorCode = v.HumanError(), v.HTTPCode()
					}

					log.Println(err)
					w.Write([]byte(errorString))
					w.WriteHeader(errorCode)
					return
				}
			})
		}

		var handleFunc = func(path string, handler Handler) {
			http.Handle(path, errorHandling(middleware(handler)))
		}
		// End of imported Twitch sample code*/

	//go http.ListenAndServeTLS(fmt.Sprintf("%s:443",ServerConfig.BindWebIP),ServerConfig.SSLCert,ServerConfig.SSLKey,nil)
	DebugLog.Println("Listening on SSL")
	DebugLog.Println("Started running on http://localhost:2222")
	path, _ := filepath.Abs("./html")
	DebugLog.Println(path)
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)
	DebugLog.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", ServerConfig.BindWebIP, ServerConfig.BindWebPort), nil))
	//IsErr("",err)

}
