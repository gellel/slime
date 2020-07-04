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

// cookieUUIDHTTPOnly is the HTTP flag for the UUID http.Cookie.
const cookieUUIDHTTPOnly bool = true

// cookieUUIDName is the name for the UUID http.Cookie.
const cookieUUIDName string = "uuid"

// cookieUUIDPath is the path for the UUID http.Cookie.
const cookieUUIDPath string = "/"

// favicon is the common namespace for the favicon.
const favicon string = "favicon"

// faviconPath is the HTTP route for the favicon file.
const faviconPath string = fileNameFavicon

// fileImageGIF is the representation expression of a 1x1 transparent gif.
const fileImageGIF string = "" +
	"\x47\x49\x46\x38" +
	"\x39\x61\x01\x00" +
	"\x01\x00\x80\x00" +
	"\x00\x00\x00\x00" +
	"\x00\x00\x00\x21" +
	"\xF9\x04\x01\x00" +
	"\x00\x00\x00\x2C" +
	"\x00\x00\x00\x00" +
	"\x01\x00\x01\x00" +
	"\x00\x02\x02\x44" +
	"\x01\x00\x3B"

// fileNameFavicon is the name for the favicon ico file.
const fileNameFavicon string = (favicon + ".ico")

// fileNameIFrame is the name for the iframe HTML file.
const fileNameIFrame string = "iframe.html"

// fileNameJavaScript is the name for the JavaScript file.
const fileNameJavaScript string = "main.js"

// flagPortName is the flag name for the port flag.
const flagPortName string = "port"

// flagPortValue is the default value for the port flag.
const flagPortValue int = 5000

// flagVerboseName is the flag name for the verbose flag.
const flagVerboseName string = "verbose"

// flagVerboseValue is the flag value for the verbose flag.
const flagVerboseValue bool = true

var cookieUUID http.Cookie = http.Cookie{
	Name: cookieUUIDName,
	Path: cookieUUIDPath}

// flagPort is the port flag argument for the application. flagPort controls the network port the application uses.
var flagPort *int = flag.Int(flagPortName, flagPortValue, (fmt.Sprintf("-%s %d", flagPortName, flagPortValue)))

// flagVerbose is the flag name for the verbose flag. flagVerbose controls the level of output the application generates.
var flagVerbose *bool = flag.Bool(flagVerboseName, flagVerboseValue, (fmt.Sprintf("-%s %t", flagVerboseName, flagVerboseValue)))

// filename is the name of the file being executed.
var _, filename, _, _ = runtime.Caller(0)

// filefolder is the name of the folder for the file being executed.
var filefolder string = filepath.Dir(filename)

// fileImageGIFBytes is the byte sequence for a gif serve via HTTP.
var fileImageGIFBytes = ([]byte(fileImageGIF))

// imageHeaders is a map of HTTP headers for image HTTP requests.
var imageHeaders = http.Header{
	w3g.CacheControl:       {"max-age=0", "must-revalidate", "public"},
	w3g.ContentDisposition: {"inline"},
	w3g.ContentType:        {"image/gif"},
	w3g.TimingAllowOrigin:  {"*"}}

// cacheHandler is the HTTP handler for all cache requests.
func cacheHandler(w http.ResponseWriter, r *http.Request) {
	var UUID uuid.UUID = newCacheUUID(r)
	w.Header().Add(w3g.ETag, UUID.String())
}

// defaultHandler is the HTTP handler for all root requests.
func defaultHandler(w http.ResponseWriter, r *http.Request) {}

// faviconHandler is the HTTP handler for all favicon requests.
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(filefolder, favicon, fileNameFavicon))
}

// iframeHandler is the HTTP handler for all iframe requests.
func iframeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(filefolder, fileNameIFrame))
}

// imageHandler is the HTTP handler for all image requests.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	var UUID uuid.UUID = newCacheUUID(r)
	for k, substrings := range imageHeaders {
		for _, v := range substrings {
			w.Header().Add(k, v)
		}
	}
	w.Header().Set(w3g.ETag, UUID.String())
	w.Header().Set(w3g.XRequestID, (uuid.New().String()))
	w.WriteHeader(http.StatusOK)
	w.Write(fileImageGIFBytes)
	if *flagVerbose {
		log.Println(r)
	}
}

// jsHandler is the HTTP handler for all JS requests.
func jsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(filefolder, fileNameJavaScript))
}

// newCacheUUID creates a new uuid.UUID from either a request.Header, http.Cookie or url.Query.
func newCacheUUID(r *http.Request) (UUID uuid.UUID) {
	var ok bool
	if UUID, ok = getCookieUUID(r); ok {
		return
	} else if UUID, ok = getIfNoneMatchUUID(r); ok {
		return
	} else if UUID, ok = getQueryUUID(r); ok {
		return
	}
	UUID = (uuid.New())
	return
}

// getIfNoneMatchValue gets the If-None-Match HTTP header value.
func getIfNoneMatchValue(r *http.Request) (s string, ok bool) {
	s = (r.Header.Get(w3g.IfNoneMatch))
	ok = (!(len(s) == 0))
	return
}

// getIfNoneMatchUUID tries to get a uuid.UUID from the http.Header If-None-Match if the HTTP header is found.
func getIfNoneMatchUUID(r *http.Request) (UUID uuid.UUID, ok bool) {
	var err error
	var s string
	s, ok = getIfNoneMatchValue(r)
	if !ok {
		return
	}
	UUID, err = uuid.Parse(s)
	ok = (err == nil)
	return
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
	return
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
	return
}

// getQueryValue gets a url.Query value.
func getQueryValue(r *http.Request, name string) (s string, ok bool) {
	s = r.URL.Query().Get(name)
	ok = (!(len(s) == 0))
	return
}

func getQueryUUID(r *http.Request) (UUID uuid.UUID, ok bool) {
	var err error
	var s string
	s, ok = getQueryValue(r, cookieUUIDName)
	if !ok {
		return
	}
	UUID, err = uuid.Parse(s)
	ok = (err == nil)
	return
}

func main() {
	flag.Parse()
	if *flagVerbose {
		log.Println(os.Args)
	}
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc((fmt.Sprintf("/%s", faviconPath)), faviconHandler)
	http.HandleFunc((fmt.Sprintf("/%s", "i")), imageHandler)
	http.HandleFunc((fmt.Sprintf("/%s", "f")), iframeHandler)
	http.HandleFunc((fmt.Sprintf("/%s", "j")), jsHandler)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), nil))
}
