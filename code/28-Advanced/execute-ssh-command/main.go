package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

func main() {

	argsWithProg := os.Args
	var s string

	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	s = filepath.Base(path)

	if argsWithProg[1] == "help" {
		hlp := s + ` <hostname or ip address> <username> <path to the private key> "<command to execute/ or if multiple commands in brackets>"`
		fmt.Println(hlp)
		return

	}
	if len(argsWithProg) != 5 {
		fmt.Println(s+" <hostname or ip address> <username> <path to the private key>", "<command to execute/ or if multiple commands in brackets>")

		return
	} else {
		SshAndRunCommand(argsWithProg)
	}

}

func SshAndRunCommand(args []string) error {

	key, err := ioutil.ReadFile(args[3])
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: args[2],
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	server := args[1] + ":22"

	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(args[4]); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	return nil
}
