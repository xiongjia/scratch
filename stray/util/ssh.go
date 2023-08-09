package util

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func NewClient() {
	config := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.Password("test"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "192.168.71.237:22", config)
	if err != nil {
		fmt.Println("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stderr = &b
	session.Stdout = &b
	if err := session.Run("hostname && ls && df -ah"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
