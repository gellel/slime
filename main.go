package main

import (
	"flag"
	"fmt"
)

// flagPortName is the flag name for the port flag.
const flagPortName string = "port"

// flagPortValue is the default value for the port flag.
const flagPortValue string = "5000"

// flagPort is the port flag argument for the application. flagPort controls the network port the application uses.
var flagPort *string = flag.String(flagPortName, flagPortValue, (fmt.Sprintf("-%s %s", flagPortName, flagPortValue)))

// flagVerboseName is the flag name for the verbose flag.
const flagVerboseName string = "verbose"

// flagVerboseValue is the flag value for the verbose flag.
const flagVerboseValue bool = true

// flagVerbose is the flag name for the verbose flag. flagVerbose controls the level of output the application generates.
var flagVerbose *bool = flag.Bool(flagVerboseName, flagVerboseValue, (fmt.Sprintf("-%s %t", flagVerboseName, flagVerboseValue)))

func main() {
	flag.Parse()
}
