package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	cnx, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cnx.Close()
	eCh := make(chan error)

	buf := make([]byte, 10)
	for {
		n, err := cnx.Read(buf)
		log.Printf("buf value %v len %v", string(buf), n)
		if err != nil {
			// send an error if it is encountered
			fmt.Println(err)
			eCh <- err
			return
		}
	}
}
