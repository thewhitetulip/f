package main

// This is a very simple http server in Go
// to be used while development whenever we need a server that does just rendering nothing else
// with no configuration just one argument, the port number
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// we require the user to give the port number where the server will run
	if len(os.Args) == 1 {
		fmt.Println("Usage: f 9090 \n Will run server on port 9090")
		os.Exit(1)
	}
	port := os.Args[1]
	if IsValidPort(port) {
		fmt.Println("Running on http://localhost:" + port)
		http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
	} else {
		fmt.Println("Please pass valid port number.(1-65535)")
	}
}

// IsValidPort will take an argument and return True if it is valid port and false othewise
func IsValidPort(v interface{}) bool {
	valid := false // we assume the port is wrong & return true otherwise
	switch t := v.(type) {
	case int:
		if v.(int) > 0 {
			if v.(int) <= 65535 {
				valid = true
			}
		}
		_ = t
	}
	return valid
}
