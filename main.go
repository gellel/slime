package main

import (
	"flag"
	"fmt"
)

// flagPortName is the flag name for the port flag.
const flagPortName string = "port"

// flagPortValue is the default value for the port flag.
const flagPortValue string = "5000"

// flagPort is the port flag argument for the application.
var flagPort = flag.String(flagPortName, flagPortValue, (fmt.Sprintf("-%s %s", flagPortName, flagPortValue)))

// flagVerboseName is the flag name for the verbose flag.
const flagVerboseName string = "v"

// flagVerboseValue is the flag value for the verbose flag.
const flagVerboseValue bool = true

func main() {
	flag.Parse()
}
