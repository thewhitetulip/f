package main

// This is a very simple http server in Go
// to be used while development whenever we need a server that does just rendering nothing else
// with no configuration just one argument, the port number
import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// we require the user to give the port number where the server will run
	if len(os.Args) == 1 {
		fmt.Println("Usage: f 9090 \n Will run server on port 9090")
		os.Exit(1)
	}
	port := os.Args[1]
	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("This is not a number, please pass a valid number")
		os.Exit(1)
	}
	if IsValidPort(portInt) {
		fmt.Println("Running on http://localhost:" + port)
		err := http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
		if err != nil {
			fmt.Println("Please use another port number ")
			os.Exit(1)
		}
	} else {
		fmt.Println("Please pass valid port number.(1-65535)")
	}
}

// IsValidPort will take an argument and return True if it is valid port and false othewise
func IsValidPort(v int) bool {
	valid := false // we assume the port is wrong & return true otherwise
	if v > 0 && v <= 65535 {
		valid = true
	}
	return valid
}
