package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

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

var _, filename, _, _ = runtime.Caller(0)

func defaultHandler(w http.ResponseWriter, r *http.Request) {}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(filepath.Dir(filename), "favicon", "favicon.ico"))
}

func main() {
	flag.Parse()
	if *flagVerbose {
		log.Println(os.Args)
	}
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc((fmt.Sprintf("/%s", "favicon.ico")), faviconHandler)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), nil))
}
