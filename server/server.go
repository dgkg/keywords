package main

import (
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	servAddr := "localhost:8080"

	ln, err := net.Listen("tcp", servAddr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go HandlConn(conn)
	}
}

func HandlConn(conn net.Conn) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			ticker = time.NewTicker(time.Duration(rand.Intn(5)) + 1*time.Second)
			stringBuf := String(20)
			_, err := conn.Write([]byte(stringBuf))
			if err != nil {
				println("Write to server failed:", err.Error())
				os.Exit(1)
			}
			println("write to client = ", stringBuf)

		case <-quit:
			conn.Close()
			ticker.Stop()
			return
		}
	}
}
