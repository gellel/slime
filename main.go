package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gellel/w3g"
	"github.com/google/uuid"
)

type request struct{}

type response struct{}

// cookieUUIDName is the name for the UUID http.Cookie.
const cookieUUIDName string = "uuid"

// faviconPath is the HTTP route for the favicon file.
const faviconPath string = fileNameFavicon

// fileNameFavicon is the name for the favicon ico file.
const fileNameFavicon string = "favicon.ico"

// fileNameIFrame is the name for the iframe HTML file.
const fileNameIFrame string = "iframe.html"

// flagPortName is the flag name for the port flag.
const flagPortName string = "port"

// flagPortValue is the default value for the port flag.
const flagPortValue int = 5000

// flagPort is the port flag argument for the application. flagPort controls the network port the application uses.
var flagPort *int = flag.Int(flagPortName, flagPortValue, (fmt.Sprintf("-%s %d", flagPortName, flagPortValue)))

// flagVerboseName is the flag name for the verbose flag.
const flagVerboseName string = "verbose"

// flagVerboseValue is the flag value for the verbose flag.
const flagVerboseValue bool = true

// flagVerbose is the flag name for the verbose flag. flagVerbose controls the level of output the application generates.
var flagVerbose *bool = flag.Bool(flagVerboseName, flagVerboseValue, (fmt.Sprintf("-%s %t", flagVerboseName, flagVerboseValue)))

// filename is the name of the file being executed.
var _, filename, _, _ = runtime.Caller(0)

// filefolder is the name of the folder for the file being executed.
var filefolder string = filepath.Dir(filename)

// cacheHandler is the HTTP handler for all cache requests.
func cacheHandler(w http.ResponseWriter, r *http.Request) {

}

// defaultHandler is the HTTP handler for all root requests.
func defaultHandler(w http.ResponseWriter, r *http.Request) {}

// faviconHandler is the HTTP handler for all favicon requests.
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(filefolder, "favicon", fileNameFavicon))
}

// iframeHandler is the HTTP handle for all iframe requests.
func iframeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(filefolder, fileNameIFrame))
}

// jsHandler is the HTTP handler for all JS requests.
func jsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(filefolder, "main.js"))
}

// newCacheUUID creates a new uuid.UUID from either a request.Header, http.Cookie or url.Query.
func newCacheUUID(r *http.Request) (UUID uuid.UUID) {
	var ok bool
	if UUID, ok = getCookieUUID(r); ok {
		return UUID
	} else if UUID, ok = getIfNoneMatchUUID(r); ok {
		return UUID
	} else if UUID, ok = getQueryUUID(r); ok {
		return UUID
	}
	UUID = (uuid.New())
	return UUID
}

// getIfNoneMatchValue gets the If-None-Match HTTP header value.
func getIfNoneMatchValue(r *http.Request) (s string, ok bool) {
	s = (r.Header.Get(w3g.IfNoneMatch))
	ok = (!(len(s) == 0))
	return s, ok
}

// getIfNoneMatchUUID tries to get a uuid.UUID from the http.Header If-None-Match if the HTTP header is found.
func getIfNoneMatchUUID(r *http.Request) (UUID uuid.UUID, ok bool) {
	var err error
	var s string
	s, ok = getIfNoneMatchValue(r)
	if !ok {
		return UUID, ok
	}
	UUID, err = uuid.Parse(s)
	ok = (err == nil)
	return UUID, ok
}

// getCookieValue gets a http.Cookie value.
func getCookieValue(r *http.Request, name string) (s string, ok bool) {
	var cookie *http.Cookie
	var err error
	cookie, err = r.Cookie(name)
	ok = (err == nil)
	if ok {
		s = cookie.Value
	}
	return s, ok
}

// getCookieUUID tries to get a uuid.UUID from the UUID http.Cookie if the UUID http.Cookie is found.
func getCookieUUID(r *http.Request) (UUID uuid.UUID, ok bool) {
	var err error
	var s string
	s, ok = getCookieValue(r, cookieUUIDName)
	if !ok {
		return UUID, ok
	}
	UUID, err = uuid.Parse(s)
	ok = (err == nil)
	return UUID, ok
}

// getQueryValue gets a url.Query value.
func getQueryValue(r *http.Request, name string) (s string, ok bool) {
	s = r.URL.Query().Get(name)
	ok = (!(len(s) == 0))
	return s, ok
}

func main() {
	flag.Parse()
	if *flagVerbose {
		log.Println(os.Args)
	}
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc((fmt.Sprintf("/%s", faviconPath)), faviconHandler)
	http.HandleFunc((fmt.Sprintf("/%s", "i")), iframeHandler)
	http.HandleFunc((fmt.Sprintf("/%s", "j")), jsHandler)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), nil))
}
